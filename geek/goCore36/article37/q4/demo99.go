package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/wxb/goLab/geek/goCore36/article37/common"
	"github.com/wxb/goLab/geek/goCore36/article37/common/op"
)

var (
	profileNames = []string{
		"goroutine",
		"heap",
		"allocs",
		"threadcreate",
		"block",
		"mutex",
	}

	profileOps = map[string]common.OpFunc{
		"goroutine":    op.BlockProfile,
		"heap":         op.MemProfile,
		"allocs":       op.MemProfile,
		"threadcreate": op.BlockProfile,
		"block":        op.BlockProfile,
		"mutex":        op.BlockProfile,
	}

	debugOpts = []int{
		0,
		1,
		2,
	}
)

func init() {
	runtime.MemProfileRate = 8
	runtime.SetBlockProfileRate(2)
}

func createDir() (string, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(currDir, "profiles")
	err = os.Mkdir(path, 0766)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	return path, nil
}

func genProfile(dir string, name string, debug int) error {
	fmt.Printf("Generate %s profile (debug: %d) ...\n", name, debug)
	fileName := fmt.Sprintf("%s_%d.out", name, debug)
	f, err := common.CreateFile(dir, fileName)
	if err != nil {
		fmt.Printf("create error: %v (%s)\n", err, fileName)
		return err
	}
	defer f.Close()
	if err = common.Execute(profileOps[name], 10); err != nil {
		fmt.Printf("execute error: %v (%s)\n", err, fileName)
		return err
	}
	profile := pprof.Lookup("block")
	err = profile.WriteTo(f, debug)
	if err != nil {
		fmt.Printf("write error: %v (%s)\n", err, fileName)
		return err
	}
	return nil
}

func main() {

	dir, err := createDir()
	if err != nil {
		fmt.Printf("dir creation error: %v\n", err)
		return
	}

	for _, name := range profileNames {
		for _, debug := range debugOpts {
			err = genProfile(dir, name, debug)
			if err != nil {
				return
			}
			time.Sleep(time.Millisecond)
		}
	}
}
