> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/install/)

> 李文周的 Blog 提供免费的全套 Go 语言学习教程，本文详细介绍了 Go 语言的安装步骤，还介绍了 GOPATH 和 GOROOT 是什么，以及如何配置 GOPROXY 代理。

Go1.14 版本，一步一步，从零搭建 Go 语言开发环境。

**因为 Go 语言及相关编辑工具的更新迭代，本文已于 2021/05/12 更新，可能会和视频有所出入，请以更新后的本文为准。**

**注意：**Go 语言 1.14 版本之后推荐使用 go modules 管理依赖，也不再需要把代码写在 GOPATH 目录下了，之前旧版本的教程戳这个[链接](https://www.liwenzhou.com/posts/Go/install_go_dev_old/)。

下载
--

### 下载地址

Go 官网下载地址：[https://golang.org/dl/](https://golang.org/dl/)

Go 官方镜像站（推荐）：[https://golang.google.cn/dl/](https://golang.google.cn/dl/)

### 版本的选择

Windows 平台和 Mac 平台推荐下载可执行文件版，Linux 平台下载压缩文件版。

**下图中的版本号可能并不是最新的，但总体来说安装教程是类似的。Go 语言更新迭代比较快，推荐使用较新版本，体验最新特性。**

![](https://www.liwenzhou.com/images/Go/install_go_dev/download1.png)

安装
--

### Windows 安装

此安装实例以 `64位Win10`系统安装 `Go1.14.1可执行文件版本`为例。

将上一步选好的安装包下载到本地。

![](https://www.liwenzhou.com/images/Go/install_go_dev/download2.png)

双击下载好的文件，然后按照下图的步骤安装即可。

![](https://www.liwenzhou.com/images/Go/install_go_dev/install01.png) ![](https://www.liwenzhou.com/images/Go/install_go_dev/install02.png) ![](https://www.liwenzhou.com/images/Go/install_go_dev/install03.png) ![](https://www.liwenzhou.com/images/Go/install_go_dev/install04.png) ![](https://www.liwenzhou.com/images/Go/install_go_dev/install05.png)

### Linux 下安装

如果不是要在 Linux 平台敲 go 代码就不需要在 Linux 平台安装 Go，我们开发机上写好的 go 代码只需要跨平台编译（详见文章末尾的跨平台编译）好之后就可以拷贝到 Linux 服务器上运行了，这也是 go 程序跨平台易部署的优势。

我们在版本选择页面选择并下载好`go1.14.1.linux-amd64.tar.gz`文件：

```
wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz
```

将下载好的文件解压到`/usr/local`目录下：

```
tar -zxvf go1.14.1.linux-amd64.tar.gz -C /usr/local  
```

如果提示没有权限，加上`sudo`以 root 用户的身份再运行。执行完就可以在`/usr/local/`下看到`go`目录了。

配置环境变量： Linux 下有两个文件可以配置环境变量，其中`/etc/profile`是对所有用户生效的；`$HOME/.profile`是对当前用户生效的，根据自己的情况自行选择一个文件打开，添加如下两行代码，保存退出。

```
// 全局 vim /etc/profile
// 只对当前用户 vim /home/用户名/.bashrc

// 加入下面两行
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

修改`/etc/profile`后要重启生效，修改`$HOME/.profile`后使用 source 命令加载`$HOME/.profile`文件即可生效。

```
source ~/.bashrc
```

##### Linux下环境配置详细示例

- 编辑`~/.bash_profile`文件：

```bash
vi ~/.bash_profile
```

- 追加以下内容：

```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/go
```

> `goroot`为go安装目录
> `gopath` go工作区，即编写代码存放的目录

当我们配置完毕后，可以执行 `source ~/.profile` 更新系统环境变量。

- 验证,查看版本

```bash
go version
```

 检查：

```
~ go version
go version go1.14.1 linux/amd64

```

### Mac 下安装

下载可执行文件版，直接点击**下一步**安装即可，默认会将 go 安装到`/usr/local/go`目录下。 ![](https://www.liwenzhou.com/images/Go/install_go_dev/mac_install_go.png)

### 检查

上一步安装过程执行完毕后，可以打开终端窗口，输入`go version`命令，查看安装的 Go 版本。 ![](https://www.liwenzhou.com/images/Go/install_go_dev/install06.png)

GOROOT 和 GOPATH
---------------

`GOROOT`和`GOPATH`都是环境变量，其中`GOROOT`是我们安装 go 开发包的路径，而从 Go 1.8 版本开始，Go 开发包在安装完成后会为`GOPATH`设置一个默认目录，并且在 Go1.14 及之后的版本中启用了 Go Module 模式之后，不一定非要将代码写到 GOPATH 目录下，所以也就**不需要我们再自己配置 GOPATH** 了，使用默认的即可。

想要查看你电脑上的`GOPATH`路径，只需要打开终端输入以下命令并回车：

```
go env


```

在终端输出的内容中找到`GOPATH`对应的具体路径。

### GOPROXY 非常重要

Go1.14 版本之后，都推荐使用`go mod`模式来管理依赖环境了，也不再强制我们把代码必须写在`GOPATH`下面的 src 目录了，你可以在你电脑的任意位置编写 go 代码。（网上有些教程适用于 1.11 版本之前。）

默认 GoPROXY 配置是：`GOPROXY=https://proxy.golang.org,direct`，由于国内访问不到`https://proxy.golang.org`，所以我们需要换一个 PROXY，这里推荐使用`https://goproxy.io`或`https://goproxy.cn`。

可以执行下面的命令修改 GOPROXY：

```
go env -w GOPROXY=https://goproxy.cn,direct


```

Go 开发编辑器
--------

Go 采用的是 UTF-8 编码的文本文件存放源代码，理论上使用任何一款文本编辑器都可以做 Go 语言开发，这里推荐使用`VS Code`和`Goland`。 `VS Code`是微软开源的编辑器，而`Goland`是 jetbrains 出品的付费 IDE。

我们这里使用`VS Code` 加插件做为 go 语言的开发工具。

### VS Code 介绍

`VS Code`全称`Visual Studio Code`，是微软公司开源的一款**免费**现代化轻量级代码编辑器，支持几乎所有主流的开发语言的语法高亮、智能代码补全、自定义热键、括号匹配、代码片段、代码对比 Diff、GIT 等特性，支持插件扩展，支持 Win、Mac 以及 Linux 平台。

虽然不如某些 IDE 功能强大，但是它添加 Go 扩展插件后已经足够胜任我们日常的 Go 开发。

### 下载与安装

`VS Code`官方下载地址：[https://code.visualstudio.com/Download](https://code.visualstudio.com/Download)

三大主流平台都支持，请根据自己的电脑平台选择对应的安装包。 ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_home.png) 双击下载好的安装文件，双击安装即可。

### 配置

#### 安装中文简体插件

点击左侧菜单栏最后一项`管理扩展`，在`搜索框`中输入`chinese` ，选中结果列表第一项，点击`install`安装。

安装完毕后右下角会提示`重启VS Code`，重启之后你的 VS Code 就显示中文啦！ ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode1.gif) `VSCode`主界面介绍： ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_menu.png)

#### 安装 go 扩展

现在我们要为我们的 VS Code 编辑器安装`Go`扩展插件，让它支持 Go 语言开发。 ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_plugin.png)

第一个 Go 程序
---------

### Hello World

现在我们来创建第一个 Go 项目——`hello`。在我们桌面创建一个`hello`目录。

#### go mod init

使用 go module 模式新建项目时，我们**需要**通过`go mod init 项目名`命令对项目进行初始化，该命令会在项目根目录下生成`go.mod`文件。例如，我们使用`hello`作为我们第一个 Go 项目的名称，执行如下命令。

```
go mod init hello


```

#### 编写代码

接下来在该目录中创建一个`main.go`文件：

```
package main  

import "fmt"  

func main(){  
	fmt.Println("Hello World!")  
}


```

**非常重要！！！** 如果此时 VS Code 右下角弹出提示让你安装插件，务必点 **install all** 进行安装。

这一步需要先执行完上面提到的`go env -w GOPROXY=https://goproxy.cn,direct`命令配置好`GOPROXY`。

### 编译

`go build`命令表示将源代码编译成可执行文件。

在 hello 目录下执行：

```
go build


```

编译得到的可执行文件会保存在执行编译命令的当前目录下，如果是`Windows`平台会在当前目录下找到`hello.exe`可执行文件。

可在终端直接执行该`hello.exe`文件：

```
c:\desktop\hello>hello.exe
Hello World!


```

我们还可以使用`-o`参数来指定编译后得到的可执行文件的名字。

```
go build -o heiheihei.exe


```

### Windows 下 VSCode 切换 cmd.exe 作为默认终端

如果你打开 VS Code 的终端界面出现如下图场景（注意观察红框圈中部分），那么你的`VS Code`此时正使用`powershell`作为默认终端： ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell1.png) 十分推荐你按照下面的步骤，选择`cmd.exe`作为默认的终端工具： ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell2.png) 此时，VS Code 正上方中间位置会弹出如下界面，参照下图挪动鼠标使光标选中后缀为`cmd.exe`的那一个，然后点击鼠标左键。

最后**重启 VS Code 中已经打开的终端**或者**直接重启 VS Code** 就可以了。 ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell3.png) 如果没有出现下拉三角，也没有关系，按下`Ctrl+Shift+P`，VS Code 正上方会出现一个框，你按照下图输入`shell`，然后点击指定选项即可出现上面的界面了。 ![](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell4.png)

### go run

`go run main.go`也可以执行程序，该命令本质上是先在临时目录编译程序然后再执行。

> 如果你不清楚上方关于`go run`执行机制的描述，那么你最好今后都使用`go build`编译再执行。

### go install

`go install`表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到`GOPATH`的 bin 目录下。因为我们**把`GOPATH`下的`bin`目录添加到了环境变量中**，所以我们就可以在任意地方直接执行可执行文件了。

### 跨平台编译

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，Go 语言支持跨平台编译——在当前平台（例如 Windows）下编译其他平台（例如 Linux）的可执行文件。

#### Windows 编译 Linux 可执行文件

如果我想在 Windows 下编译一个 Linux 下可执行文件，那需要怎么做呢？只需要在编译时指定目标操作系统的平台和处理器架构即可。

> 注意：无论你在 Windows 电脑上使用 VsCode 编辑器还是 Goland 编辑器，都要注意你使用的终端类型，因为不同的终端下命令不一样！！！目前的 Windows 通常默认使用的是`PowerShell`终端。

如果你的`Windows`使用的是`cmd`，那么按如下方式指定环境变量。

```
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64


```

如果你的`Windows`使用的是`PowerShell`终端，那么设置环境变量的语法为

```
$ENV:CGO_ENABLED=0
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"


```

在你的`Windows`终端下执行完上述命令后，再执行下面的命令，得到的就是能够在 Linux 平台运行的可执行文件了。

```
go build


```

#### Windows 编译 Mac 可执行文件

Windows 下编译 Mac 平台 64 位可执行程序：

cmd 终端下执行：

```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build


```

PowerShell 终端下执行：

```
$ENV:CGO_ENABLED=0
$ENV:GOOS="darwin"
$ENV:GOARCH="amd64"
go build


```

#### Mac 编译 Linux 可执行文件

Mac 电脑编译得到 Linux 平台 64 位可执行程序：

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


```

#### Mac 编译 Windows 可执行文件

Mac 电脑编译得到 Windows 平台 64 位可执行程序：

```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build


```

#### Linux 编译 Mac 可执行文件

Linux 平台下编译 Mac 平台 64 位可执行程序：

```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build


```

#### Linux 编译 Windows 可执行文件

Linux 平台下编译 Windows 平台 64 位可执行程序：

```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build


```

现在，开启你的 Go 语言学习之旅吧。人生苦短，let’s Go.

![](https://www.liwenzhou.com/images/wxgzh_qrcode.png)