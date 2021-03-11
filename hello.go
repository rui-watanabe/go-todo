package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	CreatedAt time.Time
	CreatedBy string
	Content   string
	Status    int
}

func main() {
	// fmt.Println("hello world")
	// todo := make([]Todo, 0, 5)
	// t := Todo {
	// 	CreatedAt: time.Now(),
	// 	CreatedBy: "max",
	// 	Content:   "hello",
	// 	Status:    0,
	// }
	// todo = append(todo, t)
	todo := []Todo{
		Todo{
			CreatedAt: time.Now(),
			CreatedBy: "max",
			Content:   "hello",
			Status:    0,
		},
		Todo{
			CreatedAt: time.Now(),
			CreatedBy: "leo",
			Content:   "hi",
			Status:    0,
		},
		Todo{
			CreatedAt: time.Now(),
			CreatedBy: "gin",
			Content:   "good morning",
			Status:    0,
		},
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
