package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hey", "status": 200})
	})

	router.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "hey", "status": 200})
	})

	router.Run(":8080")
}
