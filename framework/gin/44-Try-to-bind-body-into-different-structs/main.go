package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Foo string `form:"foo" json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `form:"bar" json:"bar" xml:"bar" binding:"required"`
}

func OnceBodyHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else {
		c.String(http.StatusOK, errA.Error())
	}

	// Always an error is occurred by this because c.Request.Body is EOF now.
	if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, "errB:"+errB.Error())
	}
}

// c.ShouldBindBodyWith stores body into the context before binding.This feature is only needed for some formats -- JSON, XML, MsgPack, ProtoBuf
// This has a slight impact to performance, so you should not use this method if you are enough to call binding at once.
func MultBodyHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This reads c.Request.Body and stores the result into the context.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else {
		c.String(http.StatusOK, errA.Error())
	}

	// At this time, it reuses body stored in the context.
	if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
	} else {
		c.String(http.StatusOK, errB.Error())
	}

	// And it can accepts other formats
	if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, errB2.Error())
	}
}

// For other formats, Query, Form, FormPost, FormMultipart, can be called by c.ShouldBind() multiple times without any damage to performance
func MultFormatHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	if errA := c.ShouldBind(&objA); errA == nil {
		fmt.Println(objA)
		c.String(http.StatusOK, `the body should be formA`)
	} else {
		c.String(http.StatusOK, errA.Error())
	}

	if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
	} else {
		c.String(http.StatusOK, errB.Error())
	}
}

func main() {
	egn := gin.Default()

	egn.POST("/once", OnceBodyHandler)
	egn.POST("/mult", MultBodyHandler)

	egn.GET("/other", MultFormatHandler)
	egn.POST("/other", MultFormatHandler)

	egn.Run(":8080")
}
