package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type system struct {
	name  string
	archs []string
}

type finState struct {
	filename string // filename
	err      error
}

var (
	binName  = flag.String("n", "", "Set a custom binary name")
	version  = flag.String("v", "", "Set version for binaries")
	message  = flag.String("m", "", "Set custom binary information")
	parallel = flag.Int("p", 1, "Set the number of parallel compiles")
	useCPUs  = flag.Bool("c", false, "Set the number of parallel compiles to the number of cores (overwrites `-p`)")
)

var programName string

func compiler(i chan []string, o chan finState) {
	for strs := range i {
		var fs finState
		if *binName != "" {
			fs.filename = *binName
		} else {
			fs.filename = programName
		}
		if *version != "" {
			fs.filename += "-" + *version
		}
		if *message != "" {
			fs.filename += "-" + *message
		}
		fs.filename += "-" + strs[0] + "-" + strs[1]
		if strs[0] == "windows" {
			fs.filename += ".exe"
		}

		cmd := exec.Command("go", "build", "-o", fs.filename)
		cmd.Env = append(os.Environ(),
			"GOOS="+strs[0],
			"GOARCH="+strs[1],
		)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fs.err = fmt.Errorf("Failed to compile %s: %s: %s\n", fs.filename, output, err)
		}

		o <- fs
	}
}

func main() {
	flag.Parse()
	if *useCPUs {
		numCPUs := runtime.NumCPU()
		parallel = &numCPUs
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %s\n", err)
	}
	swd := strings.Split(wd, "/")
	programName = swd[len(swd)-1]

	i := make(chan []string, NUMSYS)
	o := make(chan finState, NUMSYS)
	defer close(i)
	defer close(o)

	for x := 0; x < *parallel; x++ {
		go compiler(i, o)
	}

	for _, sys := range systems { // compile
		for _, arch := range sys.archs {
			i <- []string{sys.name, arch}
		}
	}

	for x := 0; x < NUMSYS; x++ {
		fs := <-o
		if fs.err != nil {
			log.Print(fs.err)
		} else {
			log.Printf("Finished compiling %s\n", fs.filename)
		}
	}
}
