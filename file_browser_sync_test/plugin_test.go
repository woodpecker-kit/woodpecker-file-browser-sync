package file_browser_sync_test

import (
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/woodpecker-kit/woodpecker-file-browser-sync/file_browser_sync"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
	"testing"
)

func TestCheckArgsPlugin(t *testing.T) {
	t.Log("mock FileBrowserSyncPlugin")

	testDataPathRoot, errTestDataPathRoot := testGoldenKit.GetOrCreateTestDataFullPath("args_plugin_test")
	if errTestDataPathRoot != nil {
		t.Fatal(errTestDataPathRoot)
	}

	// successArgs
	successArgsWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "successArgs")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	successArgsSettings := mockPluginSettings()
	successArgsSettings.SyncWorkSpaceAbsPath = "data"
	successArgsSettings.FileBrowserStandbyUrl = "foo.com"
	successArgsSettings.FileBrowserStandbyUsername = "baz"
	successArgsSettings.FileBrowserStandbyUserPassword = "bar"

	// modeNotSupport
	modeNotSupportWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "modeNotSupport")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	modeNotSupportSettings := mockPluginSettings()
	modeNotSupportSettings.SyncWorkSpaceAbsPath = "data"
	modeNotSupportSettings.SyncMode = "not-support"

	// fileBrowserUrlsNotSetUserName
	fileBrowserUrlsNotSetUserNameWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "fileBrowserUrlsNotSetUserName")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	fileBrowserUrlsNotSetUserNameSettings := mockPluginSettings()
	fileBrowserUrlsNotSetUserNameSettings.SyncWorkSpaceAbsPath = "data"
	fileBrowserUrlsNotSetUserNameSettings.FileBrowserUrls = []string{"foo.com"}
	fileBrowserUrlsNotSetUserNameSettings.FileBrowserUsername = ""

	// fileBrowserUrlsNotSetUserPassword
	fileBrowserUrlsNotSetUserPasswordWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "fileBrowserUrlsNotSetUserPassword")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	fileBrowserUrlsNotSetUserPasswordSettings := mockPluginSettings()
	fileBrowserUrlsNotSetUserPasswordSettings.SyncWorkSpaceAbsPath = "data"
	fileBrowserUrlsNotSetUserPasswordSettings.FileBrowserUrls = []string{"foo.com"}
	fileBrowserUrlsNotSetUserPasswordSettings.FileBrowserUsername = "baz"
	fileBrowserUrlsNotSetUserPasswordSettings.FileBrowserUserPassword = ""

	// fileBrowserStandbyUrlsNotSetUserName
	fileBrowserStandbyUrlsNotSetUserNameWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "fileBrowserStandbyUrlsNotSetUserName")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	fileBrowserStandbyUrlsNotSetUserNameSettings := mockPluginSettings()
	fileBrowserStandbyUrlsNotSetUserNameSettings.SyncWorkSpaceAbsPath = "data"
	fileBrowserStandbyUrlsNotSetUserNameSettings.FileBrowserUrls = []string{"foo.com"}
	fileBrowserStandbyUrlsNotSetUserNameSettings.FileBrowserUsername = "baz"
	fileBrowserStandbyUrlsNotSetUserNameSettings.FileBrowserUserPassword = "bar"
	fileBrowserStandbyUrlsNotSetUserNameSettings.FileBrowserStandbyUrl = "foo.com"
	fileBrowserStandbyUrlsNotSetUserNameSettings.FileBrowserStandbyUsername = ""

	// fileBrowserStandbyUrlsNotSetUserPassword
	fileBrowserStandbyUrlsNotSetUserPasswordWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "fileBrowserStandbyUrlsNotSetUserPassword")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	fileBrowserStandbyUrlsNotSetUserPasswordSettings := mockPluginSettings()
	fileBrowserStandbyUrlsNotSetUserPasswordSettings.SyncWorkSpaceAbsPath = "data"
	fileBrowserStandbyUrlsNotSetUserPasswordSettings.FileBrowserUrls = []string{"foo.com"}
	fileBrowserStandbyUrlsNotSetUserPasswordSettings.FileBrowserUsername = "baz"
	fileBrowserStandbyUrlsNotSetUserPasswordSettings.FileBrowserUserPassword = "bar"
	fileBrowserStandbyUrlsNotSetUserPasswordSettings.FileBrowserStandbyUrl = "foo.com"
	fileBrowserStandbyUrlsNotSetUserPasswordSettings.FileBrowserStandbyUsername = "baz"
	fileBrowserStandbyUrlsNotSetUserPasswordSettings.FileBrowserStandbyUserPassword = ""

	// notSetSyncWorkSpacePath
	notSetSyncWorkSpacePathWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "notSetSyncWorkSpacePath")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	notSetSyncWorkSpacePathSettings := mockPluginSettings()
	notSetSyncWorkSpacePathSettings.SyncWorkSpaceAbsPath = ""
	notSetSyncWorkSpacePathSettings.FileBrowserUrls = []string{"foo.com"}
	notSetSyncWorkSpacePathSettings.FileBrowserUsername = "baz"
	notSetSyncWorkSpacePathSettings.FileBrowserUserPassword = "bar"
	notSetSyncWorkSpacePathSettings.FileBrowserStandbyUrl = "foo.com"
	notSetSyncWorkSpacePathSettings.FileBrowserStandbyUsername = "baz"
	notSetSyncWorkSpacePathSettings.FileBrowserStandbyUserPassword = "some"

	// syncGlobsIncludeAndExcludeBoth
	syncGlobsIncludeAndExcludeBothWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "syncGlobsIncludeAndExcludeBoth")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	syncGlobsIncludeAndExcludeBothSettings := mockPluginSettings()
	syncGlobsIncludeAndExcludeBothSettings.SyncWorkSpaceAbsPath = "data"
	syncGlobsIncludeAndExcludeBothSettings.FileBrowserStandbyUrl = "foo.com"
	syncGlobsIncludeAndExcludeBothSettings.FileBrowserStandbyUsername = "baz"
	syncGlobsIncludeAndExcludeBothSettings.FileBrowserStandbyUserPassword = "bar"

	syncGlobsIncludeAndExcludeBothSettings.SyncIncludeGlobs = []string{"*.go"}
	syncGlobsIncludeAndExcludeBothSettings.SyncExcludeGlobs = []string{"*.md"}

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       file_browser_sync.Settings
		workRoot       string

		isDryRun          bool
		wantArgFlagNotErr bool
	}{
		{
			name:              "successArgs",
			woodpeckerInfo:    successArgsWoodpeckerInfo,
			settings:          successArgsSettings,
			wantArgFlagNotErr: true,
		},
		{
			name:           "modeNotSupport",
			woodpeckerInfo: modeNotSupportWoodpeckerInfo,
			settings:       modeNotSupportSettings,
		},
		{
			name:           "fileBrowserUrlsNotSetUserName",
			woodpeckerInfo: fileBrowserUrlsNotSetUserNameWoodpeckerInfo,
			settings:       fileBrowserUrlsNotSetUserNameSettings,
		},
		{
			name:           "fileBrowserUrlsNotSetUserPassword",
			woodpeckerInfo: fileBrowserUrlsNotSetUserPasswordWoodpeckerInfo,
			settings:       fileBrowserUrlsNotSetUserPasswordSettings,
		},
		{
			name:           "fileBrowserStandbyUrlsNotSetUserName",
			woodpeckerInfo: fileBrowserStandbyUrlsNotSetUserNameWoodpeckerInfo,
			settings:       fileBrowserStandbyUrlsNotSetUserNameSettings,
		},
		{
			name:           "fileBrowserStandbyUrlsNotSetUserPassword",
			woodpeckerInfo: fileBrowserStandbyUrlsNotSetUserPasswordWoodpeckerInfo,
			settings:       fileBrowserStandbyUrlsNotSetUserPasswordSettings,
		},
		{
			name:           "notSetSyncWorkSpacePath",
			woodpeckerInfo: notSetSyncWorkSpacePathWoodpeckerInfo,
			settings:       notSetSyncWorkSpacePathSettings,
		},
		{
			name:           "syncGlobsIncludeAndExcludeBoth",
			woodpeckerInfo: syncGlobsIncludeAndExcludeBothWoodpeckerInfo,
			settings:       syncGlobsIncludeAndExcludeBothSettings,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.OnlyArgsCheck()
			errPluginRun := p.Exec()
			if tc.wantArgFlagNotErr {
				if errPluginRun != nil {
					wdShotInfo := wd_short_info.ParseWoodpeckerInfo2Short(p.GetWoodPeckerInfo())
					wd_log.VerboseJsonf(wdShotInfo, "print WoodpeckerInfoShort")
					wd_log.VerboseJsonf(p.Settings, "print Settings")
					t.Fatalf("wantArgFlagNotErr %v\np.Exec() error:\n%v", tc.wantArgFlagNotErr, errPluginRun)
					return
				}
				infoShot := p.ShortInfo()
				wd_log.VerboseJsonf(infoShot, "print WoodpeckerInfoShort")
			} else {
				if errPluginRun == nil {
					t.Fatalf("test case [ %s ], wantArgFlagNotErr %v, but p.Exec() not error", tc.name, tc.wantArgFlagNotErr)
				}
				t.Logf("check args error: %v", errPluginRun)
			}
		})
	}
}

func TestPlugin(t *testing.T) {
	t.Log("do FileBrowserSyncPlugin")
	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}
	t.Log("mock FileBrowserSyncPlugin args")

	testDataPathRoot, errTestDataPathRoot := testGoldenKit.GetOrCreateTestDataFullPath("file_browser_sync_test")
	if errTestDataPathRoot != nil {
		t.Fatal(errTestDataPathRoot)
	}

	// uploadSom
	uploadSomeWorkRoot := filepath.Join(testDataPathRoot, "uploadSome")
	uploadSomeWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(uploadSomeWorkRoot),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	uploadSomeSettings := mockPluginSettings()
	uploadSomeSettings.SyncMode = file_browser_sync.SyncModeUpload
	uploadSomeSettings.SyncWorkSpaceAbsPath = "dist"
	uploadSomeSettings.SyncExcludeGlobs = []string{
		"**/*.md",
		"**/**/*.md",
		"*.apk",
		"**/*.json",
		"**/**/*.json",
		"**/**/**/*.json",
	}

	// downloadSome
	downloadSomeWorkRoot := filepath.Join(testDataPathRoot, "downloadSome")
	downloadSomeWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(downloadSomeWorkRoot),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	downloadSomeSettings := mockPluginSettings()
	downloadSomeSettings.SyncMode = file_browser_sync.SyncModeDown
	downloadSomeSettings.SyncWorkSpaceAbsPath = "dist"
	downloadSomeSettings.SyncIncludeGlobs = []string{
		"*.json",
	}

	tests := []struct {
		name             string
		woodpeckerInfo   wd_info.WoodpeckerInfo
		settings         file_browser_sync.Settings
		uploadWorkRoot   string
		downloadWorkRoot string

		ossTransferKey  string
		ossTransferData interface{}

		isDryRun bool
		wantErr  bool
	}{
		{
			name:           "uploadSome",
			woodpeckerInfo: uploadSomeWoodpeckerInfo,
			settings:       uploadSomeSettings,
			uploadWorkRoot: uploadSomeWorkRoot,
			isDryRun:       true,
		},
		{
			name:             "downloadSome",
			woodpeckerInfo:   downloadSomeWoodpeckerInfo,
			settings:         downloadSomeSettings,
			downloadWorkRoot: downloadSomeWorkRoot,
			isDryRun:         true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.uploadWorkRoot != "" {
				mockRootPath := filepath.Join(tc.uploadWorkRoot, tc.settings.SyncWorkSpaceAbsPath)
				errMockFileData := initTestDataForWorkspace(mockRootPath)
				if errMockFileData != nil {
					t.Fatal(errMockFileData)
				}
			}
			if tc.downloadWorkRoot != "" {
				mockRootPath := filepath.Join(tc.downloadWorkRoot, tc.settings.SyncWorkSpaceAbsPath)
				errNewDownloadPath := unittest_file_kit.Mkdir(mockRootPath)
				if errNewDownloadPath != nil {
					t.Fatal(errNewDownloadPath)
				}
			}
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.Settings.DryRun = tc.isDryRun
			if tc.ossTransferKey != "" {
				errGenTransferData := generateTransferStepsOut(
					p,
					tc.ossTransferKey,
					tc.ossTransferData,
				)
				if errGenTransferData != nil {
					t.Fatal(errGenTransferData)
				}
			}
			err := p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
