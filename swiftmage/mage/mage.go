package mage

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

const arrow = "➜ "

func Run(command ...string) {
	c := strings.Join(command, " ")
	LogColor(color.FgHiGreen, c)
	cmd := exec.Command("sh", "-c", c)
	cmd.Stdout = os.Stdout
	var errorBuffer bytes.Buffer
	cmd.Stderr = &errorBuffer
	err := cmd.Run()
	if err != nil {
		fmt.Println()
		LogColor(color.FgHiRed, "Error running the following command:")
		LogColor(color.FgHiGreen, "\t", c)
		rawError := strings.TrimSpace(errorBuffer.String())
		LogColor(color.FgHiRed, rawError)
		Check(err)
	}
}

func RunAt(path string, command ...string) {
	cmd := append([]string{"cd", path, "&&"}, command...)
	Run(cmd...)
}

func Check(e error) {
	if e != nil {
		color.Set(color.FgHiRed)
		fmt.Print(arrow, " ")
		panic(color.HiRedString(e.Error()))
	}
}

func InitEnvironment() {
	Log("Initializing environment...")
	PlatformSelected = os.Getenv("platform")
	if PlatformSelected == "" {
		PlatformSelected = PlatformAll
	}
}

func Log(a ...interface{}) {
	LogColor(color.FgHiCyan, a...)
}

func LogColor(c color.Attribute, a ...interface{}) {
	color.Set(c)
	msg := append([]interface{}{arrow}, a...)
	fmt.Println(msg...)
	color.Unset()
}
