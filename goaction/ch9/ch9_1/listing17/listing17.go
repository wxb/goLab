package main

import (
	"log"
	"net/http"

	"github.com/wxb/goLab/goaction/ch9/ch9_1/listing17/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener: Started:Listening on:4000")
	http.ListenAndServe(":4000", nil)
}
