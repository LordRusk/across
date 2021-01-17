package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type system struct {
	name  string
	archs []string
}

// supported operating systems
// with their support architectures.
var systems = []system{
	{name: "android", archs: []string{
		"arm",
	}},
	{name: "darwin", archs: []string{
		"amd64",
		"arm64",
	}},
	{name: "dragonfly", archs: []string{
		"amd64",
	}},
	{name: "freebsd", archs: []string{
		"386",
		"amd64",
		"arm",
	}},
	{name: "linux", archs: []string{
		"386",
		"amd64",
		"arm",
		"arm64",
		"ppc64",
		"ppc64le",
		"mips",
		"mips64",
		"mips64le",
		"mipsle",
	}},
	{name: "netbsd", archs: []string{
		"386",
		"amd64",
		"arm",
	}},
	{name: "openbsd", archs: []string{
		"386",
		"amd64",
		"arm",
	}},
	{name: "plan9", archs: []string{
		"386",
		"amd64",
	}},
	{name: "solaris", archs: []string{
		"amd64",
	}},
	{name: "windows", archs: []string{
		"386",
		"amd64",
	}},
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get working directory: %s\n", err)
		os.Exit(1)
	}
	swd := strings.Split(wd, "/")
	programName := swd[len(swd)-1]

	// compile
	for _, sys := range systems {
		for _, arch := range sys.archs {
			var execName string
			if sys.name == "windows" {
				execName = programName + "-" + sys.name + "-" + arch + ".exe"
			} else {
				execName = programName + "-" + sys.name + "-" + arch
			}

			fmt.Printf("Compiling %s...\n", execName)
			output, err := exec.Command("env", "GOOS="+sys.name, "GOARCH="+arch, "go", "build", "./").CombinedOutput()
			if err != nil {
				fmt.Printf("Failed to compile!: %s: %s\n", output, err)
			} else {
				if sys.name == "windows" {
					output, err = exec.Command("mv", programName+".exe", execName).CombinedOutput()
				} else {
					output, err = exec.Command("mv", programName, execName).CombinedOutput()
				}
				if err != nil {
					fmt.Printf("Failed to rename '%s' to '%s': %s: %s\n", programName, execName, output, err)
				}
			}
		}
	}
}
