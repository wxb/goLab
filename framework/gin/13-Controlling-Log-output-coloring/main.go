package main

import "github.com/gin-gonic/gin"

func main() {
	// Disable log's color
	{
		gin.DisableConsoleColor()
	}
	// Force log's color
	{
		// gin.ForceConsoleColor()
	}

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
