[![GoDoc Reference](https://godoc.org/github.com/mattn/go-sqlite3?status.svg)](http://godoc.org/github.com/mattn/go-sqlite3)

# task-list
This is a task list manager. With this tool you will be able to:
- Create a task
- View a task
- Update the status of a task
- Remove a task once its in a completed status

# Issues
To compile this package on Windows, you must have the gcc compiler installed.

<b>Install a Windows gcc toolchain. </b><br>
Add the bin folder to the Windows path, if the installer did not do this by default.<br>
Open a terminal for the TDM-GCC toolchain, which can be found in the Windows Start menu.<br>
Navigate to your project folder and run the go build ... command for this package.<br>
For example the TDM-GCC Toolchain can be found here: https://jmeubank.github.io/tdm-gcc/<br>

# How to use
The cli uses Cobra, so once the package is imported you run commands like these examples:
- .\task-list init
- .\task-list task new
- .\task-list task list
- .\task-list task updateStatus

<br>
 Run .\task-list for more instructions
 
 # Upcoming
 API capability to run the app in a web service and to be able to control the app by REST
