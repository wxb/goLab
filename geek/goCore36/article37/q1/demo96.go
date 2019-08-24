package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/wxb/goLab/geek/goCore36/article37/common"
	"github.com/wxb/goLab/geek/goCore36/article37/common/op"
)

var (
	profileName = "cpuprofile.out"
)

func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}

	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}

func main() {

	f, err := common.CreateFile("", profileName)
	if err != nil {

	}
	defer f.Close()

	err = startCPUProfile(f)
	if err != nil {
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}

	err = common.Execute(op.CPUProfile, 10)
	if err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}

	stopCPUProfile()
}
