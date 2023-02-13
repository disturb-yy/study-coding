> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/makefile/)

> 李文周的 Blog make makefile go build 编译

借助`Makefile`我们在编译过程中不再需要每次手动输入编译的命令和编译的参数，可以极大简化项目编译过程。

make 介绍
-------

`make`是一个构建自动化工具，会在当前目录下寻找`Makefile`或`makefile`文件。如果存在相应的文件，它就会依据其中定义好的规则完成构建任务。

Makefile 介绍
-----------

我们可以把`Makefile`简单理解为它定义了一个项目文件的编译规则。借助`Makefile`我们在编译过程中不再需要每次手动输入编译的命令和编译的参数，可以极大简化项目编译过程。同时使用`Makefile`也可以在项目中确定具体的编译规则和流程，很多开源项目中都会定义`Makefile`文件。

本文不会详细介绍`Makefile`的各种规则，只会给出 Go 项目中常用的`Makefile`示例。关于`Makefile`的详细内容推荐阅读 [Makefile 教程](http://c.biancheng.net/view/7097.html)。

### 规则概述

`Makefile`由多条规则组成，每条规则主要由两个部分组成，分别是依赖的关系和执行的命令。

其结构如下所示：

```
[target] ... : [prerequisites] ...
<tab>[command]
    ...
    ...


```

其中：

*   targets：规则的目标
*   prerequisites：可选的要生成 targets 需要的文件或者是目标。
*   command：make 需要执行的命令（任意的 shell 命令）。可以有多条命令，每一条命令占一行。

举个例子：

```
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o xx


```

示例
--

```
.PHONY: all build run gotool clean help

BINARY="bluebell"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"



```

其中：

*   `BINARY="bluebell"`是定义变量。
*   `.PHONY`用来定义伪目标。不创建目标文件，而是去执行这个目标下面的命令。
