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

	SDKIPhone  = "iphoneos"
	SDKMacOS   = "macosx"
	SDKAppleTV = "appletvos"
	SDKWatch   = "watchos"

	ConfigurationDebug   = "Debug"
	ConfigurationRelease = "Release"

	IOS11 = "11.3"
	IOS10 = "10.2"
	IOS9  = "9.0"
	IOS8  = "8.1"

	MacOS = "arch=x86_64"

	SimulatorIPhoneX  = "iPhone X"
	SimulatorIPhone7  = "iPhone 7"
	SimulatorIPhone6s = "iPhone 6s"
	SimulatorIPhone5s = "iPhone 5s"

	PathSources       = "./Sources"
	PathExport        = "./build/"
	PathArchive       = PathExport + Name + ".xcarchive"
	PathExportOptions = PathExport + "ExportOptions.plist"
	PathIpa           = PathExport + Name + ".ipa"

	UseNewBuildSystem = true

	ITunesUser     = ""
	ITunesPassword = ""
)
