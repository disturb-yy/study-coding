> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [blog.csdn.net](https://blog.csdn.net/zqw1597471882/article/details/125673231)

### 这里就不赘述 GO 安装教程了，百度靠谱的教程一大堆，这里只说 AIR 解决问题，网上教程废话连篇，还是靠我自己解决的问题，如下，不墨迹：

1. 安装 GO 语言
===========

自行百度，配置好 go 的环境变量，cmd 可以直接运行 go 之后，进行下一步

2. 安装 AIR
=========

```
go get -u github.com/cosmtrek/air

```

安装完毕后，我的 air 目录为

D:\software\go\bin\[pkg](https://so.csdn.net/so/search?q=pkg&spm=1001.2101.3001.7020)\mod\github.com\cosmtrek\air@v1.40.3

每个人的目录都不一样，自行进入你 go 的安装目录来找

3.windows 配置 AIR 环境变量
=====================

由于配置环境变量需要 exe 执行文件，先进入 ari 目录，cmd 运行

```
go build .

```

然后目录下会多出一个 air.exe 文件，再配置 windows 环境变量

![](https://img-blog.csdnimg.cn/d0262a714023492a979e99a391efb1f2.png)

 ![](https://img-blog.csdnimg.cn/d1d0cd375e6042dfa90382cc52dd1bf3.png)

 ![](https://img-blog.csdnimg.cn/3fe65b64a7ba486f8743d4cdeac54abd.png)

![](https://img-blog.csdnimg.cn/dc78a289231046c4a872e5c7f5b12817.png)

 ![](https://img-blog.csdnimg.cn/c97d8398d46941369b984858b50ea20d.png)

 之后一路确定，直到关闭该页面，重启 cmd，执行 air -v 命令

```
air -v

```

![](https://img-blog.csdnimg.cn/2a9271b68b1b407a868fc996f11b84ad.png)

 到现在就算成功配置完环境变量了

4.Go 项目根目录中配置 air.conf 文件
=========================

没有这个配置文件自己创建一个，名字为 .air.conf

![](https://img-blog.csdnimg.cn/c167f740043140abb2395131400e47f4.png)

 复制粘贴配置进去

```
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件
 
# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"
 
[build]
# 只需要写你平常编译使用的shell命令。你也可以使用 `make`
# Windows平台示例: cmd = "go build -o ./tmp/main.exe ."
cmd = "go build -o ./tmp/main.exe ."
# 由`cmd`命令得到的二进制文件名
# Windows平台示例：bin = "tmp/main.exe"
bin = "tmp/main.exe"
# 自定义执行程序的命令，可以添加额外的编译标识例如添加 GIN_MODE=release
# Windows平台示例：full_bin = "./tmp/main.exe"
# Linux平台示例：full_bin = "APP_ENV=dev APP_USER=air ./tmp/main.exe"
full_bin = "./tmp/main.exe"
# 监听以下文件扩展名的文件.
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略这些文件扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听以下指定目录的文件
include_dir = []
# 排除以下文件
exclude_file = []
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件。
stop_on_error = true
# air的日志文件名，该日志文件放置在你的`tmp_dir`中
log = "air_errors.log"
 
[log]
# 显示日志时间
time = true
 
[color]
# 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"
 
[misc]
# 退出时删除tmp目录
clean_on_exit = true
```

5. 运行 AIR
=========

[打开 cmd](https://so.csdn.net/so/search?q=%E6%89%93%E5%BC%80cmd&spm=1001.2101.3001.7020)，进入你 GO 项目根目录，执行 air 命令，执行成功，完毕

```
air

```

结束，终结了这个 windows air 运行问题。