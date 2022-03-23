/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/task-list/api"
	"github.com/task-list/cmd"
	"github.com/task-list/data"
)

var Tasks []api.Task

func main() {
	data.OpenDatabase()
	cmd.Execute()
	//srv := api.NewServer()
	//http.ListenAndServe(":8080", srv)
	//api.StartServer()
	//runServer()
}

func runServer() {
	r := gin.Default()

	taskRoutes := r.Group("/tasks")

	{
		taskRoutes.GET("/", ListTasks)
		taskRoutes.POST("/", NewTask)
		//		taskRoutes.PUT("/:id", UpdateTask)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve "localhost:8080"
}

//func getHandler(c *gin.Context) {
//	c.JSON(200, Tasks) // Return nil
//}

func ListTasks(c *gin.Context) {
	// null
	c.JSON(200, Tasks)
}

//func UpdateTask() {

//}

func NewTask(c *gin.Context) {
	var reqBody api.Task

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
