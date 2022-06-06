package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine:=gin.Default()

	ua:=""
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	engine.LoadHTMLGlob("templates/*")

	engine.GET("/pages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello gin",
		})
	})

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
			"User-Agent": ua,
		})
	})
	engine.Run(":3030")
}