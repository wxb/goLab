package main

import (
    "io"
    "net/http"
)

func main(){
    handler := func(w http.ResponseWriter, req *http.Request){
        io.WriteString(w, "hello, world\n")
    }

    http.HandleFunc("/", handler)

    http.ListenAndServe(":8082", nil)
}
