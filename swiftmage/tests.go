// +build mage

package main

import . "./mage"

func TestIOSDebug() {
	testIOS("Debug")
}

func TestIOSRelease() {
	testIOS("Release")
}

func testIOS(configuration string) {
	destinations := []Destination{
		DestinationForSimulator("11.3", "iPhone X"),
		DestinationForSimulator("10.3.1", "iPhone 7 Plus"),
		DestinationForSimulator("9.0", "iPhone X"),
	}
	testDestinations(PlatformIOS, destinations, configuration)
}

func TestMacOSDebug() {
	testMacOS("Debug")
}

func TestMacOSRelease() {
	testMacOS("Release")
}

func testMacOS(configuration string) {
	destinations := []Destination{DestinationForMac()}
	testDestinations(PlatformMacOS, destinations, configuration)
}

func TestTVOSDebug() {
	testTVOS("Debug")
}

func TestTVOSRelease() {
	testTVOS("Release")
}

func testTVOS(configuration string) {
	destinations := []Destination{
		DestinationForSimulator("11.3", "Apple TV 4K"),
		DestinationForSimulator("10.2", "Apple TV 1080p"),
		DestinationForSimulator("9.0", "Apple TV 1080p"),
	}
	testDestinations(PlatformTVOS, destinations, configuration)
}

func BuildWatchOSDebug() {
	buildWatchOS("Debug")
}

func BuildWatchOSRelease() {
	buildWatchOS("Release")
}

func buildWatchOS(configuration string) {
	destinations := []Destination{
		DestinationForSimulator("11.3", "Apple TV 4K"),
		DestinationForSimulator("10.2", "Apple TV 1080p"),
		DestinationForSimulator("9.0", "Apple TV 1080p"),
	}
	for _, destination := range destinations {
		build := xCodeBuildWorkspace
		build.Build(destination, configuration)
	}
}

func testDestinations(platform Platform, destinations []Destination, configuration string) {
	Clean()
	build := xCodeBuildWorkspace
	build.BuildForTesting(DestinationGeneric(platform))
	for _, destination := range destinations {
		test := xCodeBuildWorkspace
		test.Test(destination, configuration, false)
	}
}
