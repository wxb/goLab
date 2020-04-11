package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wxb/goLab/GraphQL/gin/controller"
)

func main() {

	r := gin.Default()

	r.POST("/graphql", controller.GraphqlHandler())
	r.GET("/graphql", controller.GraphqlHandler())

	r.Run(":1234")
}
