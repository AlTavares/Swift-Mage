package mage

type Platform string

var (
	Name      = "MyProject"
	Workspace = Name + ".xcworkspace"
	Scheme    = Name

	FrameworkIOSScheme     = Name + "-iOS"
	FrameworkMacOSScheme   = Name + "-macOS"
	FrameworkTVOSScheme    = Name + "-tvOS"
	FrameworkWatchOSScheme = Name + "-watchOS"

	PlatformSelected = PlatformAll
	PlatformAll      = Platform("all")
	PlatformIOS      = Platform("iOS")
	PlatformMacOS    = Platform("macOS")
	PlatformTVOS     = Platform("tvOS")
	PlatformWatchOS  = Platform("watchOS")

	PathSources       = "./Sources"
	PathExport        = "./build/"
	PathArchive       = PathExport + Name + ".xcarchive"
	PathExportOptions = PathExport + "ExportOptions.plist"
	PathIpa           = PathExport + Name + ".ipa"

	ITunesUser     = ""
	ITunesPassword = ""
)
