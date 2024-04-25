package file_browser_sync

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"strings"
)

const (
	CliNameSyncMode = "settings.sync-mode"
	EnvSyncMode     = "PLUGIN_SYNC_MODE"

	CliNameFileBrowserUrls = "settings.file-browser-urls"
	EnvFileBrowserUrls     = "PLUGIN_FILE_BROWSER_URLS"

	CliNameFileBrowserUsername = "settings.file-browser-username"
	EnvFileBrowserUsername     = "PLUGIN_FILE_BROWSER_USERNAME"

	CliNameFileBrowserUserPassword = "settings.file-browser-user-password"
	EnvFileBrowserUserPassword     = "PLUGIN_FILE_BROWSER_USER_PASSWORD"

	CliNameFileBrowserStandbyUrl = "settings.file-browser-standby-url"
	EnvFileBrowserStandbyUrl     = "PLUGIN_FILE_BROWSER_STANDBY_URL"

	CliNameFileBrowserStandbyUsername = "settings.file-browser-standby-username"
	EnvFileBrowserStandbyUsername     = "PLUGIN_FILE_BROWSER_STANDBY_USERNAME"

	CliNameFileBrowserStandbyUserPassword = "settings.file-browser-standby-user-passwords"
	EnvFileBrowserStandbyUserPassword     = "PLUGIN_FILE_BROWSER_STANDBY_USER_PASSWORDS"

	CliNameSyncWorkSpacePath = "settings.sync-work-space-path"
	EnvSyncWorkSpacePath     = "PLUGIN_SYNC_WORK_SPACE_PATH"

	CliNameSyncIncludeGlobs = "settings.sync-include-globs"
	EnvSyncIncludeGlobs     = "PLUGIN_SYNC_INCLUDE_GLOBS"

	CliNameSyncExcludeGlobs = "settings.sync-exclude-globs"
	EnvSyncExcludeGlobs     = "PLUGIN_SYNC_EXCLUDE_GLOBS"

	CliNameSyncDryRun = "settings.sync-dry-run"
	EnvSyncDryRun     = "PLUGIN_SYNC_DRY_RUN"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{

		&cli.StringFlag{
			Name:    CliNameSyncMode,
			Usage:   fmt.Sprintf("set sync mode, support: %s", strings.Join(syncModeSupport, ", ")),
			Value:   SyncModeDown,
			EnvVars: []string{EnvSyncMode},
		},
		&cli.StringSliceFlag{
			Name:    CliNameFileBrowserUrls,
			Usage:   "set file browser support multi urls, will auto switch host fast, if not set or host not work, will use standby url",
			EnvVars: []string{EnvFileBrowserUrls},
		},
		&cli.StringFlag{
			Name:    CliNameFileBrowserUsername,
			Usage:   "set file browser username for multi urls",
			EnvVars: []string{EnvFileBrowserUsername},
		},
		&cli.StringFlag{
			Name:    CliNameFileBrowserUserPassword,
			Usage:   "set file browser user password for multi urls",
			EnvVars: []string{EnvFileBrowserUserPassword},
		},
		&cli.StringFlag{
			Name:    CliNameFileBrowserStandbyUrl,
			Usage:   "set file browser standby url, if  multi urls not work, will use this",
			EnvVars: []string{EnvFileBrowserStandbyUrl},
		},
		&cli.StringFlag{
			Name:    CliNameFileBrowserStandbyUsername,
			Usage:   "set file browser username for standby url",
			EnvVars: []string{EnvFileBrowserStandbyUsername},
		},
		&cli.StringFlag{
			Name:    CliNameFileBrowserStandbyUserPassword,
			Usage:   "set file browser user password for standby url",
			EnvVars: []string{EnvFileBrowserStandbyUserPassword},
		},
		&cli.StringFlag{
			Name:    CliNameSyncWorkSpacePath,
			Usage:   "sync path under workspace path",
			EnvVars: []string{EnvSyncWorkSpacePath},
		},
		&cli.StringSliceFlag{
			Name:    CliNameSyncIncludeGlobs,
			Usage:   "sync include globs",
			EnvVars: []string{EnvSyncIncludeGlobs},
		},
		&cli.StringSliceFlag{
			Name:    CliNameSyncExcludeGlobs,
			Usage:   "sync exclude globs",
			EnvVars: []string{EnvSyncExcludeGlobs},
		},
		&cli.BoolFlag{
			Name:    CliNameSyncDryRun,
			Usage:   "dry run, only print some info, not real sync",
			EnvVars: []string{EnvSyncDryRun},
		},
	}
}

const (
	CliNameFileBrowserSyncTimeoutSecond = "settings.file-browser-sync-timeout-second"
	EnvFileBrowserSyncTimeoutSecond     = "PLUGIN_FILE_BROWSER_SYNC_TIMEOUT_SECOND"
)

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    CliNameFileBrowserSyncTimeoutSecond,
			Usage:   "file browser sync timeout second",
			Hidden:  true,
			EnvVars: []string{EnvFileBrowserSyncTimeoutSecond},
		},
	}
}

func BindCliFlags(c *cli.Context,
	debug bool,
	cliName, cliVersion string,
	wdInfo *wd_info.WoodpeckerInfo,
	rootPath,
	stepsTransferPath string, stepsOutDisable bool,
) (*FileBrowserSyncPlugin, error) {

	config := Settings{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		StepsOutDisable:   stepsOutDisable,
		RootPath:          rootPath,

		DryRun:   c.Bool(CliNameSyncDryRun),
		SyncMode: c.String(CliNameSyncMode),

		FileBrowserUrls:         c.StringSlice(CliNameFileBrowserUrls),
		FileBrowserUsername:     c.String(CliNameFileBrowserUsername),
		FileBrowserUserPassword: c.String(CliNameFileBrowserUserPassword),

		FileBrowserStandbyUrl:          c.String(CliNameFileBrowserStandbyUrl),
		FileBrowserStandbyUsername:     c.String(CliNameFileBrowserStandbyUsername),
		FileBrowserStandbyUserPassword: c.String(CliNameFileBrowserStandbyUserPassword),

		SyncWorkSpaceAbsPath: c.String(CliNameSyncWorkSpacePath),
		SyncIncludeGlobs:     c.StringSlice(CliNameSyncIncludeGlobs),
		SyncExcludeGlobs:     c.StringSlice(CliNameSyncExcludeGlobs),

		SyncTimeoutSecond: c.Uint(CliNameFileBrowserSyncTimeoutSecond),
	}

	wd_log.Debugf("args %s: %v", wd_flag.NameCliPluginTimeoutSecond, config.TimeoutSecond)

	infoShort := wd_short_info.ParseWoodpeckerInfo2Short(*wdInfo)

	p := FileBrowserSyncPlugin{
		Name:           cliName,
		Version:        cliVersion,
		woodpeckerInfo: wdInfo,
		wdShortInfo:    &infoShort,
		Settings:       config,
	}

	return &p, nil
}
