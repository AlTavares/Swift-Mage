package mage

import (
	"fmt"
	"strings"
)

type XCodeBuild struct {
	Workspace     string
	Scheme        string
	Destination   string
	Configuration string
	Action        string

	Pretty            bool
	UseNewBuildSystem bool
	EnableTestability bool
	OnlyActiveArch    bool
	AllTargets        bool

	Args []string
}

func NewXCodeBuild() XCodeBuild {
	return XCodeBuild{
		OnlyActiveArch: true,
		Pretty:         true,
	}
}

func NewXCodeBuildWithWorkspace(workspace string, scheme string) XCodeBuild {
	return XCodeBuild{
		OnlyActiveArch: true,
		Pretty:         true,
		Workspace:      workspace,
		Scheme:         scheme,
	}
}

func (xc XCodeBuild) Run() {
	Run(xc.BuildCommand())
}

func (xc XCodeBuild) BuildCommand() string {

	if xc.Workspace != "" {
		xc.AddKVArgument("-workspace", xc.Workspace)
	}
	if xc.Scheme != "" {
		xc.AddKVArgument("-scheme", xc.Scheme)
	}
	if xc.Destination != "" {
		xc.AddKVArgument("-destination", xc.Destination)
	}
	if xc.Configuration != "" {
		xc.AddKVArgument("-configuration", xc.Configuration)
	}

	if xc.UseNewBuildSystem {
		xc.AddArgument("-UseNewBuildSystem=YES")
	}
	if xc.EnableTestability {
		xc.AddArgument("ENABLE_TESTABILITY=YES")
	}
	if !xc.OnlyActiveArch {
		xc.AddArgument("ONLY_ACTIVE_ARCH=NO")
	}
	if xc.AllTargets {
		xc.AddArgument("-alltargets")
	}

	if xc.Action != "" {
		xc.AddArgument(xc.Action)
	}
	cmd := "xcodebuild " + strings.Join(xc.Args, " ")
	if xc.Pretty {
		cmd = fmt.Sprintf("set -o pipefail && %s | xcpretty -c", cmd)
	}
	return cmd
}

func (xc *XCodeBuild) AddArgument(arg string) {
	xc.Args = append(xc.Args, arg)
}

func (xc *XCodeBuild) AddKVArgument(key string, value string) {
	xc.AddArgument(key + " " + value)
}

func XCodeBuilds(command ...string) string {
	c := strings.Join(command, " ")
	return fmt.Sprintf("set -o pipefail && xcodebuild %s | xcpretty -c", c)
}

//Build target for testing
func (xc XCodeBuild) BuildForTesting(platform string) {
	xc.Destination = "generic/platform=" + platform
	xc.OnlyActiveArch = false
	xc.EnableTestability = true
	xc.Action = "build-for-testing"
	xc.Run()
}

func (xc XCodeBuild) Archive(sdk string, path string) {
	xc.AddKVArgument("-sdk", sdk)
	xc.AddKVArgument("-archivePath", path)
	xc.Configuration = ConfigurationRelease
	xc.Action = "archive"
	xc.Run()
}

func (xc XCodeBuild) ExportArchive(archivePath string, exportPath string, exportOptionsPath string) {
	xc.AddKVArgument("-archivePath", PathArchive)
	xc.AddKVArgument("-exportPath", PathExport)
	xc.AddKVArgument("-exportOptionsPlist", PathExportOptions)
	xc.AddArgument("-allowProvisioningUpdates")
	xc.Action = "-exportArchive"
	xc.Run()
}

// "'OS="+osVersion+",name="+simulator+"'"
func (xc XCodeBuild) Test(destination string, configuration string, build bool) {
	xc.Destination = destination
	xc.Configuration = configuration
	if build {
		xc.Action = "test"
	} else {
		xc.Action = "test-without-building"
	}
	xc.Run()
}

func (xc XCodeBuild) Clean() {
	xc.Action = "clean"
	xc.AllTargets = true
	xc.Run()
}
