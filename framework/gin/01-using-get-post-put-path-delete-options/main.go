package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"method": "getting"})
	})
	router.POST("/somePost", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"method": "posting"})
	})
	router.PUT("/somePut", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"method": "putting"})
	})
	router.DELETE("/someDelete", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"method": "deleting"})
	})
	router.PATCH("/somePatch", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"method": "patching"})
	})
	router.HEAD("/someHead", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"method": "head"})
	})
	router.OPTIONS("/someOptions", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"method": "options"})
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
