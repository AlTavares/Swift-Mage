package mage

var (
	Name      = "MyProject"
	Workspace = Name + ".xcworkspace"
	Scheme    = Name

	FrameworkIOSScheme     = Name + "-iOS"
	FrameworkMacOSScheme   = Name + "-macOS"
	FrameworkTVOSScheme    = Name + "-tvOS"
	FrameworkWatchOSScheme = Name + "-watchOS"

	PlatformSelected string
	PlatformAll      = "all"
	PlatformIOS      = "iOS"
	PlatformMacOS    = "macOS"
	PlatformTVOS     = "tvOS"
	PlatformWatchOS  = "watchOS"

	PathSources       = "./Sources"
	PathExport        = "./build/"
	PathArchive       = PathExport + Name + ".xcarchive"
	PathExportOptions = PathExport + "ExportOptions.plist"
	PathIpa           = PathExport + Name + ".ipa"

	ITunesUser     = ""
	ITunesPassword = ""
)
