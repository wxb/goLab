package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const (
	col = 10000
	row = 10000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {

	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	// 获取系统信息
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal("could not start CPU profile: ", err)
		return
	}
	defer func() {
		pprof.StopCPUProfile()
		f.Close()
	}()

	// 主逻辑区，进行一些简单的代码运算
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	// runtime.GC()
	err = pprof.WriteHeapProfile(f1)
	if err != nil {
		log.Fatal("could not write memory profile: ", err)
		return
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create groutine profile: ", err)
	}

	gProf := pprof.Lookup("goroutine")
	if gProf == nil {
		log.Fatal("could not write groutine profile: ")
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()

}
