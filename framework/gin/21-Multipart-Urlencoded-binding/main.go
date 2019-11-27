package main

import (
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`

	// or for multiple files
	Avatars []*multipart.FileHeader `form:"avatars" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/profile", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		var form ProfileForm
		// in this case proper binding will be automatically selected
		if err := c.ShouldBind(&form); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}

		err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}

		wg := sync.WaitGroup{}
		for _, av := range form.Avatars {
			wg.Add(1)
			go func(c *gin.Context, av *multipart.FileHeader, done func()) {
				defer done()
				c.SaveUploadedFile(av, av.Filename)
			}(c, av, wg.Done)
		}
		wg.Wait()

		// db.Save(&form)

		c.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}
