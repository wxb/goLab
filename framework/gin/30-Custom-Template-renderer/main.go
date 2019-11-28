package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	html := template.Must(template.ParseFiles("default.tmpl", "index.tmpl"))
	router.SetHTMLTemplate(html)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default.tmpl", gin.H{
			"title": "default",
		})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "index",
		})
	})

	router.Run(":8080")
}
