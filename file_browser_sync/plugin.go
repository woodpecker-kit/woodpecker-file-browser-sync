package file_browser_sync

import (
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
)

type (
	// FileBrowserSyncPlugin file_browser_sync all config
	FileBrowserSyncPlugin struct {
		Name           string
		Version        string
		woodpeckerInfo *wd_info.WoodpeckerInfo
		wdShortInfo    *wd_short_info.WoodpeckerInfoShort
		onlyArgsCheck  bool

		Settings Settings

		FuncFileBrowserSync FuncFileBrowserSync `json:"-"`

		fileBrowserClient *file_browser_client.FileBrowserClient
	}
)

type FuncFileBrowserSync interface {
	ShortInfo() wd_short_info.WoodpeckerInfoShort

	SetWoodpeckerInfo(info wd_info.WoodpeckerInfo)
	GetWoodPeckerInfo() wd_info.WoodpeckerInfo

	OnlyArgsCheck()

	Exec() error

	loadStepsTransfer() error
	checkArgs() error
	saveStepsTransfer() error
}

//nolint:golint,unused
type syncFileBrowserBiz interface {
	chooseFileBrowserConnect()

	doSyncCheck() error

	doSyncByFileBrowser() error

	doSyncModeUpload() error

	doSyncModeDown() error
}
