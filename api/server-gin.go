package api

import (
	"github.com/gin-gonic/gin"
)

var Tasks []Task

func RunServer() {
	r := gin.Default()

	taskRoutes := r.Group("/tasks")

	{
		taskRoutes.GET("/", ListTasks)
		taskRoutes.POST("/", NewTask)
		//taskRoutes.PUT("/:id", UpdateTask)
		//taskRoutes.PUT("/:id",CompleteTask)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve "localhost:8080"
}

func ListTasks(c *gin.Context) {
	// null
	c.JSON(200, Tasks)
}

func NewTask(c *gin.Context) {
	var reqBody Task

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	Tasks = append(Tasks, reqBody)

	c.JSON(200, gin.H{
		"error": false,
	})
}

//func getHandler(c *gin.Context) {
//	c.JSON(200, Tasks) // Return nil
//}

//func UpdateTask() {

//}

//func CompleteTask() {

//}

/*
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
*/
