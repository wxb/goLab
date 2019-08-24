package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func fibonacci(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1]+ret[i-2])
	}
	return ret
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000000; i++ {
		fbs = fibonacci(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/fb", createFBS)

	prefix := "/d/pprof"
	http.HandleFunc(prefix+"/cmdline", pprof.Cmdline)
	http.HandleFunc(prefix+"/profile", pprof.Profile)
	http.HandleFunc(prefix+"/symbol", pprof.Symbol)
	http.HandleFunc(prefix+"/trace", pprof.Trace)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
