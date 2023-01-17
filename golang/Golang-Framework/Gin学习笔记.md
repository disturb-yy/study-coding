# Gin 学习笔记



## 网络基础知识

新建一个简单的web应用

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http serve failed, err: %v\n", err)
	}
}

```

然后我们想让发送的“Hello”样式改变，则可以使用`html`的语法：

```go
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h2>Golang</h2>")
}
```

当这些语法很多的时候，我们可以新建一个`hello.txt`来保存这些语句，因此可以用下面的代码打开一个文件，并使用：

```go
func sayHello(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("hello.txt")  // 返回[]byte
	if err != nil {
		return
	}
	fmt.Fprintln(w, string(file))
}
```

当我们访问 `127.0.0.1/hello`时，浏览器会自动的返回刚才我们使用`html`语句创建的页面，当时当我们访问其他的目录，如 `127.0.0.1/say`的时候，就回返回一个404。

这个原因就是：当我们访问带有资源的链接时，服务器会返回该资源，而我们的浏览器能识别这种资源，并生成对应的页面。

因此，web的实质就是：**一次请求，一次响应**



但是，这样存在一个弊端，当我们返回的资源文件越来越大的时候，其对带宽的要求也越来越高。

由于页面的`html，css，js`以及图片等资源基本都是不变的，可以预先加载的，因此，我们可以只向用户传递数据，而这些静态资源则预先传递给用户，我们后面只进行数据的更新。





#### go的模板

go的模板就是一堆定义好的文本内容（一般是HTML文件格式），然后通过`{{ . }}`接收信息，生成对应的文件（最终会转换成字符串），生成一个保存了数据和文本的文件，将其写入ResponseWrite，发送给服务器。





## GIn 框架

