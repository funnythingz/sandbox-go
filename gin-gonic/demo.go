package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hey", "status": 200})
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "hey", "status": 200})
	})

	r.Run(":8080")
}
