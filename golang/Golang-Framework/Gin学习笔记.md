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





## Gin 框架



官方文档：[Gin Web Framework ](https://gin-gonic.com/zh-cn/docs/)



### 快速入门

要安装 Gin 软件包，需要先安装 Go 并设置 Go 工作区。

1.下载并安装 gin：

```sh
$ go get -u github.com/gin-gonic/gin
```

2.将 gin 引入到代码中：

```go
import "github.com/gin-gonic/gin"
```

3.实例

```go
// main.go
package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "message",
		})
	})
	r.Run(":9000")  // 监听并在 0.0.0.0:9000 上启动服务
}
```

-------------------------



### API



#### AsciiJSON

使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。

```go
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		data := map[string]any{
			"lang": "GO语言",
			"tag":  "<br>",
		}
        // 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":9000")
}
```

![image-20230123000812196](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301230008589.png)



#### HTML 渲染

使用 LoadHTMLGlob() 或者 LoadHTMLFiles()。

















#### Gin 框架模板渲染





#### Gin 返回`json`数据

##### 1 使用`map`

可以使用`map[string]any`类型来返回一个map，从而实现`json`数据的返回

```go
// gin.H 是gin框架提前定义好的 map[string]any 类型的数据
data := gin.H{
	"name":    "小王子",
	"message": "Hello world!",
	"age":     18,
}
r.GET("/json", data)
```

##### 2 使用结构体







#### GIn 获取 query string 参数

在URL的`?`后面的是query string 参数，其使用`key-value`的格式，并使用`&`连接多个`key-value`

```go
name := c.Query("query")
name := c.DefaultQuery("query", "哈哈")
name, ok := c.GetQuery("query")
```

