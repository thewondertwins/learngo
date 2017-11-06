# Getting Started

## Installing Go 

Mac or Linux

Download and install Go - always use the packages from golang.org - never use homebrew or apt-get, yum, etc. They're broken, or worse -- modified by someone upstream.

Set a GOPATH in .bashrc, .bash_profile, .zshrc etc:

	export GOPATH=$HOME/go

Add go binaries (compilers and tools) to your path:

	export PATH=$PATH:/usr/local/go/bin

Log out and back in to get the changes or

	$ source .bashrc

to hot-reload.


## Installing Go 

Windows

Download and install Go - Use MSI installer

Set a GOPATH in user Environment Variables

	GOPATH=%userdir%/go

Add go binaries (compilers and tools) to your path:

	%userdir%/go/bin	

## Verify Installation

From a command prompt:
	
	go version

You should see something like:

	go version 1.8 linux/amd64


## Editing Go Code

Popular Go Editors:

vim and neovim with vim-go plugin 

emacs with go-mode.el

Visual Studio Code with vscode-go (works with debugging!) 

Atom with go-plus

IntelliJ IDEA with Go plugin


## The Go Playground

Even if you don't have an editor configured locally you can still play with Go from your browser.

[The Go Playground](https://play.golang.org)

The Go Playground is a web service that runs on golang.org's servers. The service receives a Go program, compiles, links, and runs the program inside a sandbox, then returns the output.

## Playground Limitations

There are limitations to the programs that can be run in the playground:

The playground can use most of the standard library, with some exceptions. The only communication a playground program has to the outside world is by writing to standard output and standard error.

In the playground the time begins at 2009-11-10 23:00:00 UTC (determining the significance of this date is an exercise for the reader). This makes it easier to cache programs by giving them deterministic output.

There are also limits on execution time and on CPU and memory usage.

Therefore: No file IO, nothing useful with time or dates, can't use any external packages.

##  The Go Playground

Even with all those limitations Go developers love the Go Playground - it's a great place to share code, even if it can't run or compile. You can enter code then click the "SHARE" button which will give you a permanent URL to that code.

Try it now with this link: 

[Hello World!](https://play.golang.org/p/992fMmkkxr) 

## The Go Command

All of your command line interaction with Go will be through the `go` command.

Several common commands:
	
	go build some/package
	go run main.go
	go test some/package
	go get github.com/someone/package
	go install some/package



## Exercise

From your command prompt type `go` and hit return to see the various tools the `go` command implements.  Try some like:

	go env
	go list
	go version

## Download Material

From your command prompt use the `go` command to download the materials and exercises for this book:

	go get github.com/thewondertwins/learngo

