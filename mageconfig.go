// +build mage

package main

var (
	Name             = "MyProject"
	Workspace        = Name + ".xcworkspace"
	Scheme           = Name
	SchemeTestflight = Name

	SchemeIOS     = Name + "-iOS"
	SchemeMacOS   = Name + "-macOS"
	SchemeTVOS    = Name + "-tvOS"
	SchemeWatchOS = Name + "-watchOS"

	PlatformSelected = "all"

	PathSources       = "./Sources"
	PathExport        = "./build/"
	PathArchive       = PathExport + Name + ".xcarchive"
	PathExportOptions = PathExport + "ExportOptions.plist"
	PathIpa           = PathExport + Name + ".ipa"

	ITunesUser     = ""
	ITunesPassword = ""
)
