package mage

import (
	"fmt"
	"testing"
)

func init() {
	DryRun = true
}

func TestXCodeBuild_Archive(t *testing.T) {
	xc := XCodeBuild{}
	xc.Archive("sdktest", "pathtest")
	cmd := xc.BuildCommand()
	expected := "xcodebuild -sdk sdktest -archivePath pathtest -configuration Release archive"
	assertEqual(t, cmd, expected)
}

func TestXCodeBuild_ExportArchive(t *testing.T) {
	xc := XCodeBuild{}
	xc.ExportArchive("archivepathtest", "exportpathtest", "exportoptionspathtest")
	cmd := xc.BuildCommand()
	expected := "xcodebuild -archivePath archivepathtest -exportPath exportpathtest -exportOptionsPlist exportoptionspathtest -allowProvisioningUpdates -exportArchive"
	assertEqual(t, cmd, expected)
}

func TestXCodeBuild_Build(t *testing.T) {
	xc := XCodeBuild{}
	xc.Build(DestinationForSimulator("1.0", "destinationtest"), "configurationtest")
	cmd := xc.BuildCommand()
	expected := "xcodebuild -destination 'OS=1.0,name=destinationtest' -configuration configurationtest ONLY_ACTIVE_ARCH=NO build"
	assertEqual(t, cmd, expected)
}

func TestXCodeBuild_BuildForTesting(t *testing.T) {
	xc := XCodeBuild{}
	xc.BuildForTesting(DestinationGeneric(PlatformIOS))
	cmd := xc.BuildCommand()
	expected := "xcodebuild -destination 'generic/platform=iOS' ONLY_ACTIVE_ARCH=NO ENABLE_TESTABILITY=YES build-for-testing"
	assertEqual(t, cmd, expected)
}
func TestXCodeBuild_TestWithoutBuilding(t *testing.T) {
	xc := XCodeBuild{}
	xc.Test(DestinationForSimulator("1.0", "destinationtest"), "configurationtest", false)
	cmd := xc.BuildCommand()
	expected := "xcodebuild -destination 'OS=1.0,name=destinationtest' -configuration configurationtest ONLY_ACTIVE_ARCH=NO ENABLE_TESTABILITY=YES test-without-building"
	assertEqual(t, cmd, expected)
}

func TestXCodeBuild_TestBuilding(t *testing.T) {
	xc := XCodeBuild{}
	xc.Test(DestinationForMac(), "configurationtest", true)
	cmd := xc.BuildCommand()
	expected := "xcodebuild -destination 'arch=x86_64' -configuration configurationtest ONLY_ACTIVE_ARCH=NO ENABLE_TESTABILITY=YES test"
	assertEqual(t, cmd, expected)
}

func TestXCodeBuild_Clean(t *testing.T) {
	xc := XCodeBuild{}
	xc.Clean()
	cmd := xc.BuildCommand()
	expected := "xcodebuild clean -alltargets"
	assertEqual(t, cmd, expected)
}

func TestXCodeBuild_Pretty(t *testing.T) {
	xc := XCodeBuild{}
	xc.Pretty().Action("someaction")
	cmd := xc.BuildCommand()
	expected := "set -o pipefail && xcodebuild someaction | xcpretty -c"
	assertEqual(t, cmd, expected)
}

func TestXCodeBuild_RandomArgs(t *testing.T) {
	xc := XCodeBuild{}
	xc.AddKVArgument("-keyarg", "valuearg").
		AddArgument("anotherargument")
	cmd := xc.BuildCommand()
	expected := "xcodebuild -keyarg valuearg anotherargument"
	assertEqual(t, cmd, expected)
}

func assertEqual(t *testing.T, cmd string, expected string) {
	if cmd != expected {
		fmt.Println()
		fmt.Println("Expected:", expected)
		fmt.Println("Gotten:  ", cmd)
		fmt.Println()
		t.Fail()
	}
}
