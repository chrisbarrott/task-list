[![GoDoc Reference](https://godoc.org/github.com/mattn/go-sqlite3?status.svg)](http://godoc.org/github.com/mattn/go-sqlite3)

# task-list
Currently this can be ran as 2 different applications, either a local cobra cli application, or a more limited web app powered by REST. <br>

### The cobra cli task list manager can do the follow ran locally in VS code:
- Create a task
- View a task
- Update the status of a task
- Remove a task once its in a completed status
<p>
 
 ### To use the cobra cli ensure these functions are present in main.go:<br>
<i> 	// cobra cli application<br>
	data.OpenDatabase()<br>
	cmd.Execute()<br></i> 
<p>
 
### However a more limited web app can be used by using the following in main.go: <br>
 <i> // to run as a webserver<br>
 srv := api.NewServer()<br>
 http.ListenAndServe(":8080", srv)<br></i>
 <p>

## Install
go get github.com/chrisbarrott/task-list

## Issues
To compile this package on Windows, you must have the gcc compiler installed.

<b>Install a Windows gcc toolchain. </b><br>
- Add the bin folder to the Windows path, if the installer did not do this by default.<br>
- Open a terminal for the TDM-GCC toolchain, which can be found in the Windows Start menu.<br>
- Navigate to your project folder and run the go build ... command for this package.<br>
- For example the TDM-GCC Toolchain can be found here: https://jmeubank.github.io/tdm-gcc/<br>

## How to use
The cli uses Cobra, so once the package is imported you run commands like these examples:
- .\task-list init
- .\task-list task new
- .\task-list task list
- .\task-list task updateStatus

<br>
 Run .\task-list for more instructions
 
 ## Upcoming
 API capability to run the app in a web service and to be able to control the app by REST
