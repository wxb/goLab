package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {

	*reply = "hello:" + request
	return nil
}

type A struct {
	Name string
}

type F func(i int) string

type B struct {
	A
	io.Writer
	F

	S string
}

func main() {
	b := B{
		Writer: nil,
		A:      A{"anme"},
		F:      func(i int) string { return fmt.Sprintf("%d", i) },
		// S:      "99",
	}
	fmt.Println(b.A, b.Name)
	return
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
			// ReadCloser: r.Body,
			// Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
