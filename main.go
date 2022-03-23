package main

import (
	"github.com/chrisbarrott/task-list/cmd"
	"github.com/chrisbarrott/task-list/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
	//srv := api.NewServer()
	//http.ListenAndServe(":8080", srv)
	//api.NewServer()
	//api.RunServer()
}
