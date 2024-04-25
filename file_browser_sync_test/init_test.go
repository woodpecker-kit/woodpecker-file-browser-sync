package file_browser_sync_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/woodpecker-kit/woodpecker-file-browser-sync/file_browser_sync"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	keyEnvDebug  = "CI_DEBUG"
	keyEnvCiNum  = "CI_NUMBER"
	keyEnvCiKey  = "CI_KEY"
	keyEnvCiKeys = "CI_KEYS"

	mockVersion = "v1.0.0"
	mockName    = "woodpecker-file-browser-sync"

	keyLinkSpeedTestUrls = "ENV_LINK_SPEED_TEST_URLS"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	// mustSetInCiEnvList
	//  for check set in CI env not empty
	mustSetInCiEnvList = []string{
		wd_flag.EnvKeyCiSystemPlatform,
		wd_flag.EnvKeyCiSystemVersion,
	}

	// mustSetArgsAsEnvList
	mustSetArgsAsEnvList = []string{
		file_browser_sync.EnvFileBrowserUrls,
		file_browser_sync.EnvFileBrowserUsernames,
		file_browser_sync.EnvFileBrowserUserPasswords,
	}

	valEnvTimeoutSecond                   uint
	valEnvPluginDebug                     = false
	valEnvSyncMode                        = file_browser_sync.SyncModeDown
	valEnvFileBrowserUrls                 []string
	valEnvFileBrowserUsernames            string
	valEnvFileBrowserUserPasswords        = ""
	valEnvFileBrowserStandbyUrl           = ""
	varEnvFileBrowserStandbyUsernames     = ""
	valEnvFileBrowserStandbyUserPasswords = ""
	valEnvSyncWorkSpacePath               = ""
	valEnvSyncIncludeGlobs                []string
	valEnvSyncExcludeGlobs                []string
	valEnvSyncDryRun                      = false
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	wd_log.SetLogLineDeep(2)
	// if open wd_template please open this
	//wd_template.RegisterSettings(wd_template.DefaultHelpers)

	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)

	valEnvTimeoutSecond = uint(env_kit.FetchOsEnvInt(wd_flag.EnvKeyPluginTimeoutSecond, 10))
	valEnvPluginDebug = env_kit.FetchOsEnvBool(wd_flag.EnvKeyPluginDebug, false)

	valEnvSyncMode = env_kit.FetchOsEnvStr(file_browser_sync.EnvSyncMode, file_browser_sync.SyncModeDown)
	valEnvFileBrowserUrls = env_kit.FetchOsEnvStringSlice(file_browser_sync.EnvFileBrowserUrls)
	valEnvFileBrowserUsernames = env_kit.FetchOsEnvStr(file_browser_sync.EnvFileBrowserUsernames, "")
	valEnvFileBrowserUserPasswords = env_kit.FetchOsEnvStr(file_browser_sync.EnvFileBrowserUserPasswords, "")
	valEnvFileBrowserStandbyUrl = env_kit.FetchOsEnvStr(file_browser_sync.EnvFileBrowserStandbyUrl, "")
	varEnvFileBrowserStandbyUsernames = env_kit.FetchOsEnvStr(file_browser_sync.EnvFileBrowserStandbyUsername, "")
	valEnvFileBrowserStandbyUserPasswords = env_kit.FetchOsEnvStr(file_browser_sync.EnvFileBrowserStandbyUserPassword, "")
	valEnvSyncWorkSpacePath = env_kit.FetchOsEnvStr(file_browser_sync.EnvSyncWorkSpacePath, "")
	valEnvSyncIncludeGlobs = env_kit.FetchOsEnvStringSlice(file_browser_sync.EnvSyncIncludeGlobs)
	valEnvSyncExcludeGlobs = env_kit.FetchOsEnvStringSlice(file_browser_sync.EnvSyncExcludeGlobs)
	valEnvSyncDryRun = env_kit.FetchOsEnvBool(file_browser_sync.EnvSyncDryRun, false)
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end

func envCheck(t *testing.T) bool {

	if valEnvPluginDebug {
		wd_log.OpenDebug()
	}

	// most CI system will set env CI to true
	envCI := env_kit.FetchOsEnvStr("CI", "")
	if envCI == "" {
		t.Logf("not in CI system, skip envCheck")
		return false
	}
	t.Logf("check env for CI system")
	return env_kit.MustHasEnvSetByArray(t, mustSetInCiEnvList)
}

func envMustArgsCheck(t *testing.T) bool {
	for _, item := range mustSetArgsAsEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %s, than run test\nfull need set env %v", item, mustSetArgsAsEnvList)
			return true
		}
	}
	return false
}

func generateTransferStepsOut(plugin file_browser_sync.FileBrowserSyncPlugin, mark string, data interface{}) error {
	_, err := wd_steps_transfer.Out(plugin.Settings.RootPath, plugin.Settings.StepsTransferPath, plugin.GetWoodPeckerInfo(), mark, data)
	return err
}

func mockPluginSettings() file_browser_sync.Settings {
	// all mock settings can set here
	settings := file_browser_sync.Settings{
		// use env:PLUGIN_DEBUG
		Debug:             valEnvPluginDebug,
		TimeoutSecond:     valEnvTimeoutSecond,
		RootPath:          testGoldenKit.GetTestDataFolderFullPath(),
		StepsTransferPath: wd_steps_transfer.DefaultKitStepsFileName,

		SyncMode:                       valEnvSyncMode,
		FileBrowserUrls:                valEnvFileBrowserUrls,
		FileBrowserUsername:            valEnvFileBrowserUsernames,
		FileBrowserUserPassword:        valEnvFileBrowserUserPasswords,
		FileBrowserStandbyUrl:          valEnvFileBrowserStandbyUrl,
		FileBrowserStandbyUsername:     varEnvFileBrowserStandbyUsernames,
		FileBrowserStandbyUserPassword: valEnvFileBrowserStandbyUserPasswords,
		SyncWorkSpaceAbsPath:           valEnvSyncWorkSpacePath,
		SyncIncludeGlobs:               valEnvSyncIncludeGlobs,
		SyncExcludeGlobs:               valEnvSyncExcludeGlobs,
		DryRun:                         valEnvSyncDryRun,
	}

	return settings

}

func mockPluginWithSettings(t *testing.T, woodpeckerInfo wd_info.WoodpeckerInfo, settings file_browser_sync.Settings) file_browser_sync.FileBrowserSyncPlugin {
	p := file_browser_sync.FileBrowserSyncPlugin{
		Name:    mockName,
		Version: mockVersion,
	}

	// mock woodpecker info
	//t.Log("mockPluginWithStatus")
	p.SetWoodpeckerInfo(woodpeckerInfo)

	if p.ShortInfo().Build.WorkSpace != "" {
		settings.RootPath = p.ShortInfo().Build.WorkSpace
	}

	p.Settings = settings
	return p
}
