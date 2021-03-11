package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	CreatedAtString string
	CreatedBy       string
	Content         string
	Status          int
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
			CreatedAtString: time.Now().Format("2006-03-11 21:45:05"),
			CreatedBy:       "max",
			Content:         "Frontend",
			Status:          1,
		},
		Todo{
			CreatedAtString: time.Now().Format("2006-03-11 21:45:05"),
			CreatedBy:       "leo",
			Content:         "Backend",
			Status:          0,
		},
		Todo{
			CreatedAtString: time.Now().Format("2006-03-11 21:45:05"),
			CreatedBy:       "gin",
			Content:         "Infra",
			Status:          0,
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
