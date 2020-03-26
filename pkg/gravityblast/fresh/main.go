package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    2,
			"message": "hello,world!",
		})
	})
	r.GET("/name", func(c *gin.Context) {
		fmt.Println("------------")
		c.JSON(http.StatusFound, gin.H{
			"name": "Sady",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
