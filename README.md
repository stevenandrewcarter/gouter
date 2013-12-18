Gouter
======

Router written in GoLang. The project is written as a proof of concept and a learning exercise of the Go Programming Language.
Please refer to the following sections for instructions on how to run the project, configure the project and any additional information,
which might be useful.

Installation
============

Installation of Gouter will require the Go Programming language. After installing Go you will be able to execute the application
as desired.

Install the golang libraries from http://golang.org/doc/install. Check the source code out from the repo and start it by running
	go run gouter.go
In addition you can set the port by using the flag
    go run gouter.go -port=x
where x is the port number you want the router to listen on.

Configuration
=============

You will be able to configure from the http://localhost:9090/config url. Please use the web interface to configure the routes
and for additional statistical information about Gouter.

Data
====

In the current build the data is stored locally in the data folder. In future this may change to use a proper database instead.

References
==========

Here are some useful links for the Project and GoLang in general
Command Line Flags: https://gobyexample.com/command-line-flags
Bootswatch BootStrap Themes: http://bootswatch.com/slate/
Go Templates: http://jan.newmarch.name/go/template/chapter-template.html