package main

import (
	"net/http"

	"github.com/chrisbarrott/task-list/api"
	"github.com/chrisbarrott/task-list/cmd"
	"github.com/chrisbarrott/task-list/data"
)

/*
BELOW NEEDS EDITING TO WORK IN 2 WAYS

The  data.OpenDatabase() cmd.Execute() functions run locally as a cobra cli app
The srv := api.NewServer() functions run as a web server but this functionality is limmited
*/

func main() {
	// cobra cli application
	data.OpenDatabase()
	cmd.Execute()

	// to run as a webserver
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
