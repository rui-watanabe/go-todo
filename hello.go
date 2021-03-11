package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("hello world")
	todo := []string{
		"JavaScript",
		"TypeScript",
		"Golang",
		"React",
	}
	r := gin.Default()
	r.LoadHTMLFiles("./template/index.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todo": todo,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
