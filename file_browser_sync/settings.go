package file_browser_sync

const (
	TimeoutSecondMinimum     = 10
	SyncTimeoutSecondMinimum = 60

	SyncModeSync = "upload"
	SyncModeDown = "download"

	SyncFileBrowseRemoteRootPath = ".sync"
)

type (
	// Settings file_browser_sync private config
	Settings struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string

		DryRun   bool
		SyncMode string

		FileBrowserUrls         []string
		FileBrowserUsername     string
		FileBrowserUserPassword string

		FileBrowserStandbyUrl          string
		FileBrowserStandbyUsername     string
		FileBrowserStandbyUserPassword string

		SyncWorkSpaceAbsPath string

		syncWorkSpacePath string // this path will append to RootPath

		SyncIncludeGlobs []string
		SyncExcludeGlobs []string

		SyncTimeoutSecond uint
	}
)

var (
	syncModeSupport = []string{
		SyncModeSync,
		SyncModeDown,
	}
)
