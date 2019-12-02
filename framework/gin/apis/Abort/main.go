package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	egn := gin.Default()
	egn.Use(gin.BasicAuth(gin.Accounts{
		"wangxb": "123",
	}))

	v1 := egn.Group("/v1")
	{
		v1.Use(func(ctx *gin.Context) {
			// ctx.AbortWithStatusJSON(404, "ss")
			fmt.Println("----")
			ctx.Next()
		}).Use(func(ctx *gin.Context) {
			ctx.Next()
			fmt.Println("===")
		})

		v1.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "123")
		})
	}

	v2 := egn.Group("/v2")
	{
		v2.Use(func(ctx *gin.Context) {
			fmt.Printf("v2v2v2v2")
		})

		v2.GET("/x", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "xxx")
		})
	}

	// http.ListenAndServe(":8080", egn)
	egn.Run(":8080")
}
