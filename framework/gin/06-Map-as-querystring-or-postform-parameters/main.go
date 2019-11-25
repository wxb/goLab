package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		seqs := c.QueryArray("seqs")
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		ages := c.PostFormArray("ages")

		c.String(http.StatusOK, fmt.Sprintf("sqps:%v; ids: %v; names: %v; ages:%v", seqs, ids, names, ages))
	})
	router.Run(":8080")
}
