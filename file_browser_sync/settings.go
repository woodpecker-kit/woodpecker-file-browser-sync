package file_browser_sync

const (
	TimeoutSecondMinimum     = 10
	SyncTimeoutSecondMinimum = 60

	TryConnectTimeoutSecond = 5
	TryConnectRetries       = 3

	SyncModeUpload = "upload"
	SyncModeDown   = "download"

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

		usedFileBrowserUrl          string
		usedFileBrowserUsername     string
		usedFileBrowserUserPassword string

		SyncWorkSpaceAbsPath string

		syncWorkSpacePath  string // this path will append to RootPath
		syncRemoteRootPath string // this path will append to remote root path

		SyncIncludeGlobs []string
		SyncExcludeGlobs []string

		SyncTimeoutSecond uint
	}
)

var (
	syncModeSupport = []string{
		SyncModeUpload,
		SyncModeDown,
	}
)
