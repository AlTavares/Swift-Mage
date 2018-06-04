package mage

import (
	"fmt"
	"strings"
)

type XCodeBuild struct {
	pretty bool
	args   []string
}

func NewXCodeBuild() *XCodeBuild {
	xc := XCodeBuild{}
	return xc.Pretty()
}

func NewXCodeBuildWithWorkspace(workspace string, scheme string) *XCodeBuild {
	return NewXCodeBuild().
		Workspace(workspace).
		Scheme(scheme)
}

func (xc XCodeBuild) Run() {
	Run(xc.BuildCommand())
}

func (xc XCodeBuild) BuildCommand() string {
	cmd := "xcodebuild " + strings.Join(xc.args, " ")
	if xc.pretty {
		cmd = fmt.Sprintf("set -o pipefail && %s | xcpretty -c", cmd)
	}
	return cmd
}

//Build target for testing
func (xc *XCodeBuild) BuildForTesting(platform string) {
	xc.Destination("generic/platform=" + platform).
		DisableOnlyActiveArch().
		EnableTestability().
		Action("build-for-testing").
		Run()
}

func (xc *XCodeBuild) Archive(sdk string, path string) {
	xc.AddKVArgument("-sdk", sdk).
		AddKVArgument("-archivePath", path).
		Configuration("Release").
		Action("archive").
		Run()
}

func (xc *XCodeBuild) ExportArchive(archivePath string, exportPath string, exportOptionsPath string) {
	xc.AddKVArgument("-archivePath", archivePath).
		AddKVArgument("-exportPath", exportPath).
		AddKVArgument("-exportOptionsPlist", exportOptionsPath).
		AddArgument("-allowProvisioningUpdates").
		Action("-exportArchive").
		Run()
}

// "'OS="+osVersion+",name="+simulator+"'"
func (xc *XCodeBuild) Test(destination string, configuration string, build bool) {
	xc.Destination(destination).
		Configuration(configuration)
	if build {
		xc.Action("test")
	} else {
		xc.Action("test-without-building")
	}
	xc.Run()
}

func (xc *XCodeBuild) Clean() {
	xc.Action("clean").
		AllTargets().
		Run()
}

// builder

func (xc *XCodeBuild) Action(value string) *XCodeBuild {
	if value != "" {
		xc.AddArgument(value)
	}
	return xc
}

func (xc *XCodeBuild) Workspace(value string) *XCodeBuild {
	if value != "" {
		xc.AddKVArgument("-workspace", value)
	}
	return xc
}
func (xc *XCodeBuild) Scheme(value string) *XCodeBuild {
	if value != "" {
		xc.AddKVArgument("-scheme", value)
	}
	return xc
}
func (xc *XCodeBuild) Destination(value string) *XCodeBuild {
	if value != "" {
		xc.AddKVArgument("-destination", value)
	}
	return xc
}
func (xc *XCodeBuild) Configuration(value string) *XCodeBuild {
	if value != "" {
		xc.AddKVArgument("-configuration", value)
	}
	return xc
}

func (xc *XCodeBuild) UseNewBuildSystem() *XCodeBuild {
	xc.AddArgument("-UseNewBuildSystem=YES")
	return xc
}
func (xc *XCodeBuild) EnableTestability() *XCodeBuild {
	xc.AddArgument("ENABLE_TESTABILITY=YES")
	return xc
}
func (xc *XCodeBuild) DisableOnlyActiveArch() *XCodeBuild {
	xc.AddArgument("ONLY_ACTIVE_ARCH=NO")
	return xc
}
func (xc *XCodeBuild) AllTargets() *XCodeBuild {
	xc.AddArgument("-alltargets")
	return xc
}

func (xc *XCodeBuild) Pretty() *XCodeBuild {
	xc.pretty = true
	return xc
}

func (xc *XCodeBuild) AddArgument(arg string) *XCodeBuild {
	xc.args = append(xc.args, arg)
	return xc
}

func (xc *XCodeBuild) AddKVArgument(key string, value string) *XCodeBuild {
	xc.AddArgument(key + " " + value)
	return xc
}
