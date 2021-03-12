package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id              int    `form:"id"`
	CreatedBy       string `form:"name"`
	Content         string `form:"content"`
	CreatedAtString string
	Status          int `form:"status"`
}

var todo []Todo
var idMax = 0

func addId() int {
	idMax += 1
	return idMax
}

func GetDataTodo(c *gin.Context) {
	var b Todo
	c.Bind(&b)
	b.Id = addId()
	b.Status = 0
	b.CreatedAtString = time.Now().Format("2006-03-11 21:45:05")
	todo = append(todo, b)
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func GetDoneTodo(c *gin.Context) {
	var b Todo
	c.Bind(&b)

	if b.Status == 0 {
		b.Status = 1
	} else {
		b.Status = 0
	}

	for idx, t := range todo {
		if t.Id == b.Id {
			todo[idx].Status = b.Status
		}
	}

	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
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
	todo = []Todo{
		Todo{
			Id:              addId(),
			CreatedAtString: time.Now().Format("2006-03-11 21:45:05"),
			CreatedBy:       "max",
			Content:         "Frontend",
			Status:          1,
		},
		Todo{
			Id:              addId(),
			CreatedAtString: time.Now().Format("2006-03-11 21:45:05"),
			CreatedBy:       "leo",
			Content:         "Backend",
			Status:          0,
		},
		Todo{
			Id:              addId(),
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
	r.GET("/todo", GetDataTodo)
	r.GET("/done", GetDoneTodo)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
