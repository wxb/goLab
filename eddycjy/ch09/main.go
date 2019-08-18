package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/wxb/goLab/eddycjy/ch09/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
