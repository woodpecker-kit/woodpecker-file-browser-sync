package file_browser_sync

import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/sinlov/filebrowser-client/tools/folder"
	tools "github.com/sinlov/filebrowser-client/tools/str_tools"
	"github.com/woodpecker-kit/woodpecker-file-browser-sync/internal/path_glob"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"path"
	"strings"
)

func (p *FileBrowserSyncPlugin) chooseFileBrowserConnect() {

	var connectUrl string
	var connectUsername string
	var connectPassword string

	if len(p.Settings.FileBrowserUrls) > 0 {
		linkSpeed := NewLinkSpeed(TryConnectTimeoutSecond, TryConnectRetries)
		bestUrl, err := linkSpeed.BestLinkIgnoreRetry(p.Settings.FileBrowserUrls)
		if err != nil {
			connectUrl = p.Settings.FileBrowserStandbyUrl
			connectUsername = p.Settings.FileBrowserStandbyUsername
			connectPassword = p.Settings.FileBrowserStandbyUserPassword
		} else {
			connectUrl = bestUrl
			connectUsername = p.Settings.FileBrowserUsername
			connectPassword = p.Settings.FileBrowserUserPassword
		}
	} else {
		connectUrl = p.Settings.FileBrowserStandbyUrl
		connectUsername = p.Settings.FileBrowserStandbyUsername
		connectPassword = p.Settings.FileBrowserStandbyUserPassword
	}

	p.Settings.usedFileBrowserUrl = connectUrl
	p.Settings.usedFileBrowserUsername = connectUsername
	p.Settings.usedFileBrowserUserPassword = connectPassword
}

func (p *FileBrowserSyncPlugin) doSyncCheck() error {
	if p.Settings.usedFileBrowserUrl == "" {
		return fmt.Errorf("do sync before check file browser url is empty")
	}
	if p.Settings.usedFileBrowserUsername == "" {
		return fmt.Errorf("do sync before check file browser username is empty")
	}
	if p.Settings.usedFileBrowserUserPassword == "" {
		return fmt.Errorf("do sync before check file browser user password is empty")
	}

	var remoteRealRootPath = strings.TrimRight(SyncFileBrowseRemoteRootPath, "/")
	remoteRealRootPath = path.Join(remoteRealRootPath, p.ShortInfo().Repo.Hostname,
		p.ShortInfo().Repo.OwnerName, p.ShortInfo().Repo.ShortName, p.ShortInfo().Build.Number)

	wd_log.Debugf("sync remote root path: %s", remoteRealRootPath)
	p.Settings.syncRemoteRootPath = remoteRealRootPath

	return nil
}

func (p *FileBrowserSyncPlugin) doSyncByFileBrowser() error {
	if p.Settings.syncRemoteRootPath == "" {
		return fmt.Errorf("doSyncByFileBrowser check remote root path is empty")
	}

	fileBrowserClient, errNew := file_browser_client.NewClient(
		p.Settings.usedFileBrowserUsername,
		p.Settings.usedFileBrowserUserPassword,
		p.Settings.usedFileBrowserUrl,
		p.Settings.TimeoutSecond,
		p.Settings.SyncTimeoutSecond,
	)
	if errNew != nil {
		return errNew
	}
	fileBrowserClient.Debug(p.Settings.Debug)

	p.fileBrowserClient = &fileBrowserClient

	var syncErr error
	switch p.Settings.SyncMode {
	default:
		return fmt.Errorf("not support sync mode: %s", p.Settings.SyncMode)
	case SyncModeUpload:
		syncErr = p.doSyncModeUpload()
	case SyncModeDown:
		syncErr = p.doSyncModeDown()
	}
	if syncErr != nil {
		return syncErr
	}

	return nil
}

func (p *FileBrowserSyncPlugin) doSyncModeUpload() error {
	if p.fileBrowserClient == nil {
		return fmt.Errorf("doSyncModeUpload file browser client is nil")
	}
	var fileSendPathList []string
	if len(p.Settings.SyncIncludeGlobs) > 0 {
		wd_log.Debugf("target file want include find by File Glob: %v", p.Settings.SyncIncludeGlobs)
		for _, glob := range p.Settings.SyncIncludeGlobs {
			walkByGlob, errWalkAllByGlob := folder.WalkAllByGlob(p.Settings.syncWorkSpacePath, glob, true)
			if errWalkAllByGlob != nil {
				return fmt.Errorf("file browser want send file local path with glob %s be err: %v", p.Settings.syncWorkSpacePath, errWalkAllByGlob)
			}
			wd_log.Debugf("target path include find by File Glob [ %s ] files:\n%s", glob, strings.Join(walkByGlob, "\n"))
			fileSendPathList = append(fileSendPathList, walkByGlob...)
		}
	} else {
		// sync full file path to remote
		wd_log.Debugf("sync full file path to remote")
		walkAllByMatchPath, errWalkAllByMatchPath := folder.WalkAllByMatchPath(p.Settings.syncWorkSpacePath, ".*", true)
		if errWalkAllByMatchPath != nil {
			return fmt.Errorf("file browser want send file local path be err: %v", errWalkAllByMatchPath)
		}
		fileSendPathList = walkAllByMatchPath
	}

	if len(fileSendPathList) == 0 {
		wd_log.Warnf("no file need sync to remote")
		return nil
	}

	var fileExcludePathList []string
	if len(p.Settings.SyncExcludeGlobs) > 0 {
		wd_log.Debugf("target file want exclude by File Glob: %v", p.Settings.SyncExcludeGlobs)
		for _, glob := range p.Settings.SyncExcludeGlobs {
			walkByGlob, errWalkAllByGlob := folder.WalkAllByGlob(p.Settings.syncWorkSpacePath, glob, true)
			if errWalkAllByGlob != nil {
				return fmt.Errorf("file browser want send file local path with glob %s be err: %v", p.Settings.syncWorkSpacePath, errWalkAllByGlob)
			}
			wd_log.Debugf("target path exclued find by File Glob [ %s ] files:\n%s", glob, strings.Join(walkByGlob, "\n"))
			fileExcludePathList = append(fileExcludePathList, walkByGlob...)
		}
	}

	if len(fileExcludePathList) > 0 {
		// remove exclude file path from fileSendPathList
		wd_log.Debugf("remove exclude file path from fileSendPathList")
		var finalFileSendPathList []string
		for _, sendPath := range fileSendPathList {
			isExistSend := ""
			for _, excludePath := range fileExcludePathList {
				if sendPath == excludePath {
					wd_log.Debugf("find out exclude file path: %s", excludePath)
					isExistSend = excludePath
					continue
				}
			}
			if isExistSend == "" {
				finalFileSendPathList = append(finalFileSendPathList, sendPath)
			}
		}
		fileSendPathList = finalFileSendPathList
	}

	if len(fileSendPathList) == 0 {
		wd_log.Warnf("no file need sync to remote")
		return nil
	}
	fileSendPathList = tools.StrArrRemoveDuplicates(fileSendPathList)

	if p.Settings.Debug {
		wd_log.Debugf("file send path list:\n%s", strings.Join(fileSendPathList, "\n"))
	}

	if p.Settings.DryRun {
		wd_log.Infof("dry run mode, skip sync file upload to remote")
		wd_log.Infof("local sync upload root path  : %s", p.Settings.SyncWorkSpaceAbsPath)
		wd_log.Infof("remote upload file root path : %s", p.Settings.syncRemoteRootPath)
		if len(fileSendPathList) == 0 {
			wd_log.Infof("no file need upload to remote")
			return nil
		}
		var absPathList []string
		for _, fullPath := range fileSendPathList {
			shortPath := strings.TrimLeft(fullPath, p.Settings.syncWorkSpacePath)
			absPathList = append(absPathList, shortPath)
		}
		wd_log.Infof("want send file path:\n%s", strings.Join(absPathList, "\n"))
		return nil
	}

	if len(fileSendPathList) == 0 {
		wd_log.Infof("no file need upload to remote")
		return nil
	}

	errLogin := p.fileBrowserClient.Login()
	if errLogin != nil {
		return errLogin
	}

	for _, item := range fileSendPathList {
		var resourcePost = file_browser_client.ResourcePostFile{
			LocalFilePath:  item,
			RemoteFilePath: fetchRemotePathByLocalRoot(item, p.Settings.syncWorkSpacePath, p.Settings.syncRemoteRootPath),
		}
		errSendOneFile := p.fileBrowserClient.ResourcesPostFile(resourcePost, true)
		if errSendOneFile != nil {
			return errSendOneFile
		}
		shortPath := strings.TrimLeft(item, p.Settings.syncWorkSpacePath)
		wd_log.Debugf("-> send file: %s\nto remote: %s", shortPath, resourcePost.RemoteFilePath)
		wd_log.Infof("-> send file: %s\nto remote: %s", shortPath, resourcePost.RemoteFilePath)
	}

	return nil
}

func fetchRemotePathByLocalRoot(localAbsPath, localRootPath, remoteRootPath string) string {
	remotePath := strings.Replace(localAbsPath, localRootPath, "", -1)
	remotePath = folder.Path2WebPath(remotePath)
	return fmt.Sprintf("%s/%s", remoteRootPath, remotePath)
}

func (p *FileBrowserSyncPlugin) doSyncModeDown() error {
	if p.fileBrowserClient == nil {
		return fmt.Errorf("doSyncModeDown file browser client is nil")
	}

	errLogin := p.fileBrowserClient.Login()
	if errLogin != nil {
		return errLogin
	}

	remoteResourceRoot, errRemoteResourceRoot := p.fileBrowserClient.ResourcesGet(p.Settings.syncRemoteRootPath)
	if errRemoteResourceRoot != nil {
		wd_log.Infof("now find any file at remote root path: %s", p.Settings.syncRemoteRootPath)
		return nil
	}

	wd_log.Debugf("remoteResourceRoot NumFiles: %d", remoteResourceRoot.NumFiles)
	wd_log.Debugf("remoteResourceRoot NumDirs : %d", remoteResourceRoot.NumDirs)

	// findOut RemotePathList
	var remoteFullFilePathList []string
	errFetchFilePaths := fetchFileBrowserRemoteFilePathsByRemote(&remoteFullFilePathList, p.fileBrowserClient, p.Settings.syncRemoteRootPath)
	if errFetchFilePaths != nil {
		return errFetchFilePaths
	}
	wd_log.Debugf("remoteFullFilePathList:\n%s", strings.Join(remoteFullFilePathList, "\n"))

	if len(remoteFullFilePathList) == 0 {
		wd_log.Infof("no file need sync to local at remote root path: %s", p.Settings.syncRemoteRootPath)
		return nil
	}

	var totalDownloadRemoteShortPath []string

	if len(p.Settings.SyncIncludeGlobs) == 0 && len(p.Settings.SyncExcludeGlobs) == 0 {
		var sortFilePaths []string
		for _, s := range remoteFullFilePathList {
			left := strings.TrimLeft(s, "/")
			sortPath := strings.Replace(left, p.Settings.syncRemoteRootPath, "", 1)
			sortFilePaths = append(sortFilePaths, sortPath)
		}
		totalDownloadRemoteShortPath = sortFilePaths
	} else {
		if len(p.Settings.SyncIncludeGlobs) > 0 {
			includeRemotePath, errFindInclude := findIncludeRemotePathByGlob(remoteFullFilePathList, p.Settings.SyncIncludeGlobs, p.Settings.syncRemoteRootPath)
			if errFindInclude != nil {
				return errFindInclude
			}
			totalDownloadRemoteShortPath = includeRemotePath
		}

		if len(p.Settings.SyncExcludeGlobs) > 0 {
			excludeRemotePath, errFindExclude := findExcludeRemotePathByGlob(remoteFullFilePathList, p.Settings.SyncExcludeGlobs, p.Settings.syncRemoteRootPath)
			if errFindExclude != nil {
				return errFindExclude
			}
			if len(totalDownloadRemoteShortPath) > 0 {
				for _, excludePath := range excludeRemotePath {
					for index, downloadPath := range totalDownloadRemoteShortPath {
						if downloadPath == excludePath {
							totalDownloadRemoteShortPath = append(totalDownloadRemoteShortPath[:index], totalDownloadRemoteShortPath[index+1:]...)
						}
					}
				}
			}
		}
	}

	if p.Settings.DryRun {
		wd_log.Infof("dry run mode, skip sync file download to local")
		wd_log.Infof("local sync download root path  : %s", p.Settings.SyncWorkSpaceAbsPath)
		wd_log.Infof("remote download file root path : %s", p.Settings.syncRemoteRootPath)
		if len(totalDownloadRemoteShortPath) == 0 {
			wd_log.Infof("no file need download to local")
			return nil
		}
		var absPathList []string
		for _, fullPath := range totalDownloadRemoteShortPath {
			shortPath := strings.TrimLeft(fullPath, p.Settings.syncWorkSpacePath)
			absPathList = append(absPathList, shortPath)
		}
		wd_log.Infof("want dowload file path:\n%s", strings.Join(absPathList, "\n"))
		return nil
	}

	wd_log.Infof("-> download remote file to workspace abs path: %s", p.Settings.SyncWorkSpaceAbsPath)
	errDownload := downloadRemoteByShortPath(p.Settings.syncWorkSpacePath, p.fileBrowserClient, p.Settings.syncRemoteRootPath, totalDownloadRemoteShortPath)
	if errDownload != nil {
		return errDownload
	}

	return nil
}

func downloadRemoteByShortPath(localRootPath string, client *file_browser_client.FileBrowserClient, remoteRootPath string, downloadPathList []string) error {
	if len(downloadPathList) == 0 {
		wd_log.Infof("no file need download to local")
		return nil
	}
	for _, shotPath := range downloadPathList {
		remoteFullPath := path.Join(remoteRootPath, shotPath)
		localFullPath := path.Join(localRootPath, shotPath)
		errDownload := client.ResourceDownload(remoteFullPath, localFullPath, true)
		if errDownload != nil {
			return errDownload
		}
		showInWsPath := strings.TrimLeft(shotPath, "/")
		wd_log.Debugf("-> download remote file: %s\nto local: %s", remoteFullPath, showInWsPath)
		wd_log.Infof("download to abs path: %s", showInWsPath)
	}
	return nil
}

func findIncludeRemotePathByGlob(list []string, globs []string, rootPath string) ([]string, error) {

	var sortFilePaths []string
	for _, s := range list {
		left := strings.TrimLeft(s, "/")
		sortPath := strings.Replace(left, rootPath, "", 1)
		sortFilePaths = append(sortFilePaths, sortPath)
	}

	var matchFilePaths []string
	for _, glob := range globs {
		for _, sortPath := range sortFilePaths {
			// check path is match glob
			matchGlob, errGlob := path_glob.IsPathMatchGlob(sortPath, glob)

			if matchGlob && errGlob == nil {
				// remove not match path
				matchFilePaths = append(matchFilePaths, sortPath)
			}

		}
	}

	if len(matchFilePaths) == 0 {
		return nil, nil
	}

	return matchFilePaths, nil
}

func findExcludeRemotePathByGlob(list []string, globs []string, rootPath string) ([]string, error) {

	var sortFilePaths []string
	for _, s := range list {
		left := strings.TrimLeft(s, "/")
		sortPath := strings.Replace(left, rootPath, "", 1)
		sortFilePaths = append(sortFilePaths, sortPath)
	}

	var matchFilePaths []string
	for _, glob := range globs {
		for _, sortPath := range sortFilePaths {
			// check path is match glob
			matchGlob, errGlob := path_glob.IsPathMatchGlob(sortPath, glob)

			if matchGlob || errGlob != nil {
				// remove not match path
				matchFilePaths = append(matchFilePaths, sortPath)
			}

		}
	}

	if len(matchFilePaths) == 0 {
		return nil, nil
	}

	return matchFilePaths, nil
}

func fetchFileBrowserRemoteFilePathsByRemote(fileList *[]string, client *file_browser_client.FileBrowserClient, path string) error {

	remoteResourceRoot, errRemoteResourceRoot := client.ResourcesGet(path)
	if errRemoteResourceRoot != nil {
		return errRemoteResourceRoot
	}

	if remoteResourceRoot.IsDir {
		if len(remoteResourceRoot.Items) > 0 {
			for _, item := range remoteResourceRoot.Items {
				err := fetchFileBrowserRemoteFilePathsByRemote(fileList, client, item.Path)
				if err != nil {
					return err
				}
			}
		}
	} else {
		nowList := *fileList
		nowList = append(nowList, remoteResourceRoot.Path)
		*fileList = nowList
	}
	return nil
}
