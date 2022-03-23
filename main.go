package main

import (
	"github.com/task-list/cmd"
	"github.com/task-list/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
	//srv := api.NewServer()
	//http.ListenAndServe(":8080", srv)
	//api.StartServer()
	//api.RunServer()
}
