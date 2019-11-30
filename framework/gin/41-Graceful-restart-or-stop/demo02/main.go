package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var setting = struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	HTTPPort     int
}{
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
	HTTPPort:     8080,
}

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	egn := gin.Default()
	egn.GET("/", func(ctx *gin.Context) {
		// time.Sleep(5 * time.Second)
		// ctx.JSON(200, gin.H{"name": "王小二"})
		ctx.JSON(200, gin.H{"name": "王二小"})
	})

	// endless.ListenAndServe(endPoint, egn)
	server := endless.NewServer(endPoint, egn)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
