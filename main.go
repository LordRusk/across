package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type system struct {
	name  string
	archs []string
}

var version = flag.String("v", "", "Set version for binaries")
var message = flag.String("m", "", "Set custom binary information")

func main() {
	flag.Parse()
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
				execName = programName
				if *version != "" {
					execName += "-" + *version
				}
				if *message != "" {
					execName += "-" + *message
				}
				execName += "-" + sys.name + "-" + arch + ".exe"
			} else {
				execName = programName
				if *version != "" {
					execName += "-" + *version
				}
				if *message != "" {
					execName += "-" + *message
				}
				execName += "-" + sys.name + "-" + arch
			}

			fmt.Printf("Compiling %s...\n", execName)
			if err := os.Setenv("GOOS", sys.name); err != nil {
				fmt.Printf("Failed to compile: %s\n", err)
				continue
			}

			if err := os.Setenv("GOARCH", arch); err != nil {
				fmt.Printf("Failed to compile: %s\n", err)
				continue
			}

			output, err := exec.Command("go", "build").CombinedOutput()
			if err != nil {
				fmt.Printf("Failed to compile!: %s: %s\n", output, err)
				continue
			}

			if sys.name == "windows" {
				err = os.Rename(programName+".exe", execName)
			} else {
				err = os.Rename(programName, execName)
			}
			if err != nil {
				fmt.Printf("Failed to rename '%s' to '%s': %s\n", programName, execName, err)
			}
		}
	}
}
