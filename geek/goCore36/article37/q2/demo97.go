package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/wxb/goLab/geek/goCore36/article37/common"
	"github.com/wxb/goLab/geek/goCore36/article37/common/op"
)

var (
	profileName    = "memprofile.out"
	memProfileRate = 8
)

func startMemProfile() {
	runtime.MemProfileRate = memProfileRate
}

func stopMemProfile(f *os.File) error {

	if f == nil {
		return errors.New("nil file")
	}

	return pprof.WriteHeapProfile(f)
}

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("memory profile creation error: %v\n", err)
		return
	}
	defer f.Close()

	startMemProfile()
	err = common.Execute(op.MemProfile, 10)
	if err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}

	err = stopMemProfile(f)
	if err != nil {
		fmt.Printf("memory profile stop error: %v\n", err)
		return
	}
}
