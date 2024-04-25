package file_browser_sync

import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/sinlov/filebrowser-client/tools/folder"
	tools "github.com/sinlov/filebrowser-client/tools/str_tools"
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
					wd_log.Infof("find out exclude file path: %s", excludePath)
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
		wd_log.Infof("dry run mode, skip sync file to remote")
		return nil
	}

	errLogin := p.fileBrowserClient.Login()
	if errLogin != nil {
		return errLogin
	}

	if len(fileSendPathList) == 1 {
		localFileAbsPath := fileSendPathList[0]
		remotePath := fetchRemotePathByLocalRoot(localFileAbsPath, p.Settings.syncWorkSpacePath, p.Settings.syncRemoteRootPath)
		var resourcePostOne = file_browser_client.ResourcePostFile{
			LocalFilePath:  localFileAbsPath,
			RemoteFilePath: remotePath,
		}
		errSendOneFile := p.fileBrowserClient.ResourcesPostFile(resourcePostOne, p.Settings.Debug)
		if errSendOneFile != nil {
			return errSendOneFile
		}
	} else {
		for _, item := range fileSendPathList {
			var resourcePost = file_browser_client.ResourcePostFile{
				LocalFilePath:  item,
				RemoteFilePath: fetchRemotePathByLocalRoot(item, p.Settings.syncWorkSpacePath, p.Settings.syncRemoteRootPath),
			}
			errSendOneFile := p.fileBrowserClient.ResourcesPostFile(resourcePost, p.Settings.Debug)
			if errSendOneFile != nil {
				return errSendOneFile
			}
		}
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
	return nil
}
