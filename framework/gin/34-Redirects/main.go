// Issuing a HTTP redirect is easy. Both internal and external locations are supported.
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// external
	route.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	// internal
	route.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		route.HandleContext(c)
	})
	route.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	route.Run(":8080")
}
