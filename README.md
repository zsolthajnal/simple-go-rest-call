# Introduction 
This implements a REST API call example.

# Getting Started
Checkout and compile the source.

Modify parameters.json and set target.

Execute the binary in cmd.

This project is written in [Go](https://golang.org/). 

# Build
Execute the following to compile the project:
`go build`

######Compiling the project for different architecture:
Execute the following to list available platforms:
`go tool dist list`

Execute the following to verify current go platform settings:
`go env GOOS GOARCH`

Execute the following to compile for Windows platform (if compiled on different platform):
`GOOS=windows GOARCH=amd64 go build`
