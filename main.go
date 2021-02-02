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
			execName := programName
			if *version != "" {
				execName += "-" + *version
			}
			if *message != "" {
				execName += "-" + *message
			}
			execName += "-" + sys.name + "-" + arch
			if sys.name == "windows" {
				execName += ".exe"
			}

			fmt.Printf("Compiling %s...\n", execName)

			cmd := exec.Command("go", "build", "-o", execName)
			cmd.Env = append(os.Environ(),
				"GOOS="+sys.name,
				"GOARCH="+arch,
			)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Failed to compile!: %s: %s\n", output, err)
			}
		}
	}
}
