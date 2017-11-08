# 入门

## 安装 Go (Mac 或 Linux)

下载并安装 Go - 务必使用来自 golang.org 的包 - 不要使用 homebrew, apt-get, yum 等等。它们有问题，甚至更糟 -- 被上游的谁修改过。

在 .bashrc， .bash_profile， .zshrc 等中设定 GOPATH：

	export GOPATH=$HOME/go

添加 go 命令（编译器和相关工具）到 PATH 路径：

	export PATH=$PATH:/usr/local/go/bin

退出并重新登录，或者

	$ source .bashrc

做热加载。

## 安装 Go (Windows)

从 golang.org 下载并安装 Go - 使用 MSI 安装

在用户环境变量中设置 GOPATH

	GOPATH=%userdir%/go

添加 go 命令（编译器和相关工具）到 path 路径：

	%userdir%/go/bin	

## 验证安装

在命令行中输入：
	
	go version

应当显示类似如下信息：

	go version 1.8 linux/amd64


## 编辑 Go 代码

流行的 Go 编辑器包括：

安装了 vim-go 插件的 vim 或 neovim

带有 go-mode.el 的 emacs

Visual Studio Code 加 vscode-go (可以调试！) 

Atom 加 go-plus

IntelliJ IDEA 加 Go 插件


## Go 游乐场

即使在本地环境中没有安装合理设定的编辑器，你仍然可以在浏览器中和 Go 一起玩耍。

[Go 游乐场](https://play.golang.org)

Go 游乐场是一个运行在 golang.org 的服务器上的 Web 服务。服务器收到一个 Go 程序后，会编译，链接并且在沙盒中运行，最后输出结果。

## 游乐场的限制

游乐场对其可运行的程序有一些限制：

除了一些例外，游乐场可以使用大多数标准库。游乐场中，程序对外通讯的唯一方式是向标准输出和标准错误输出写数据。

游乐场的时间开始于 2009-11-10 23:00:00 UTC (找出为什么是这个时间，是留给读者的一个练习)。通过提供确定的输出结果，这使得缓存程序更加容易。

对于执行时间，CPU 时间和内存使用同样有限制。

因此：没有文件 IO，时间日期毫无意义，不能使用任何外部包。

##  Go 游乐场

尽管有如此之多的限制，Go 的开发者仍然热衷使用 Go 游乐场 - 它使得分享代码变得容易，即使代码无法被执行或编译。你可以输入代码，点击“分享”按钮得到这段代码的永久 URL 地址。

使用以下链接赶紧试试：

[Hello World!](https://play.golang.org/p/992fMmkkxr) 

## Go 命令

所有命令行下和 Go 的交互都是通过 `go` 命令完成的。

一些常见的命令：
	
	go build some/package
	go run main.go
	go test some/package
	go get github.com/someone/package
	go install some/package



## 练习

在命令行下输入 `go` 并回车查看 `go` 命令具有的各种工具。尝试以下这些：

	go env
	go list
	go version

## 下载课件

在命令行下使用 `go` 命令下载本书的课件和练习：

	go get github.com/thewondertwins/learngo

