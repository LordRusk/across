package main

import (
	"flag"
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

var version = flag.String("v", "", "Set version for binaries")
var message = flag.String("m", "", "Set custom binary information")
var parallel = flag.Int("p", 1, "Set the number of parallel compiles")
var useCPUs = flag.Bool("c", false, "Set the number of parallel compiles to the number of cores (overwrites `-p`)")

var programName string
var fsChan = make(chan interface{})

func compile(sys, arch string) {
	execName := programName
	if *version != "" {
		execName += "-" + *version
	}
	if *message != "" {
		execName += "-" + *message
	}
	execName += "-" + sys + "-" + arch
	if sys == "windows" {
		execName += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", execName)
	cmd.Env = append(os.Environ(),
		"GOOS="+sys,
		"GOARCH="+arch,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to compile %s: %s: %s\n", execName, output, err)
		return
	}

	log.Printf("Finished compiling %s\n", execName)
	select {
	case fsChan <- nil:
	default:
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

	stChan := make(chan []string, *parallel) // start chan
	go func() {
		for s := range stChan {
			compile(s[0], s[1])
		}
	}()

	for _, sys := range systems {
		for _, arch := range sys.archs {
			stChan <- []string{sys.name, arch}
		}
	}

	for i := 0; i <= *parallel; i++ {
		select {
		case <-fsChan:
		}
	}
}
