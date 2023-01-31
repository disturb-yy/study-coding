> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/gin/)

> 李文周的 Blog 提供免费的全套的图文和视频 Go 语言教程，本文是 Go 语言 web 框架 gin 框架（gin framework）的图文教程，详细介绍了 gin 框架的安装、基本使用、参数绑定、中间件、上传文件、路由组、重定向以及开发 RESTful API 等相关内容。

`Gin`是一个用 Go 语言编写的 web 框架。它是一个类似于`martini`但拥有更好性能的 API 框架, 由于使用了`httprouter`，速度提高了近 40 倍。 如果你是性能和高效的追求者, 你会爱上`Gin`。

Gin 框架介绍
--------

Go 世界里最流行的 Web 框架，[Github](https://github.com/gin-gonic/gin) 上有`32K+`star。 基于 [httprouter](https://github.com/julienschmidt/httprouter) 开发的 Web 框架。 [中文文档](https://gin-gonic.com/zh-cn/docs/)齐全，简单易用的轻量级框架。

Gin 框架安装与使用
-----------

### 安装

下载并安装`Gin`:

```
go get -u github.com/gin-gonic/gin


```

### 第一个 Gin 示例：

```
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}


```

将上面的代码保存并编译执行，然后使用浏览器打开`127.0.0.1:8080/hello`就能看到一串 JSON 字符串。

RESTful API
-----------

REST 与技术无关，代表的是一种软件架构风格，REST 是 Representational State Transfer 的简称，中文翻译为 “表征状态转移” 或“表现层状态转化”。

推荐阅读[阮一峰 理解 RESTful 架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)

简单来说，REST 的含义就是客户端与 Web 服务器之间进行交互的时候，使用 HTTP 协议中的 4 个请求方法代表不同的动作。

*   `GET`用来获取资源
*   `POST`用来新建资源
*   `PUT`用来更新资源
*   `DELETE`用来删除资源。

只要 API 程序遵循了 REST 风格，那就可以称其为 RESTful API。目前在前后端分离的架构中，前后端基本都是通过 RESTful API 来进行交互。

例如，我们现在要编写一个管理书籍的系统，我们可以查询对一本书进行查询、创建、更新和删除等操作，我们在编写程序的时候就要设计客户端浏览器与我们 Web 服务端交互的方式和路径。按照经验我们通常会设计成如下模式：

<table><thead><tr><th>请求方法</th><th>URL</th><th>含义</th></tr></thead><tbody><tr><td>GET</td><td>/book</td><td>查询书籍信息</td></tr><tr><td>POST</td><td>/create_book</td><td>创建书籍记录</td></tr><tr><td>POST</td><td>/update_book</td><td>更新书籍信息</td></tr><tr><td>POST</td><td>/delete_book</td><td>删除书籍信息</td></tr></tbody></table>

同样的需求我们按照 RESTful API 设计如下：

<table><thead><tr><th>请求方法</th><th>URL</th><th>含义</th></tr></thead><tbody><tr><td>GET</td><td>/book</td><td>查询书籍信息</td></tr><tr><td>POST</td><td>/book</td><td>创建书籍记录</td></tr><tr><td>PUT</td><td>/book</td><td>更新书籍信息</td></tr><tr><td>DELETE</td><td>/book</td><td>删除书籍信息</td></tr></tbody></table>

Gin 框架支持开发 RESTful API 的开发。

```
func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
}


```

开发 RESTful API 的时候我们通常使用 [Postman](https://www.getpostman.com/) 来作为客户端的测试工具。

Gin 渲染
------

### HTML 渲染

我们首先定义一个存放模板文件的`templates`文件夹，然后在其内部按照业务分别定义一个`posts`文件夹和一个`users`文件夹。 `posts/index.html`文件的内容如下：

```
{{define "posts/index.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta >
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>posts/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}


```

`users/index.html`文件的内容如下：

```
{{define "users/index.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta >
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>users/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}


```

Gin 框架中使用`LoadHTMLGlob()`或者`LoadHTMLFiles()`方法进行 HTML 模板渲染。

```
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
}


```

### 自定义模板函数

定义一个不转义相应内容的`safe`模板函数如下：

```
func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})

	router.Run(":8080")
}


```

在`index.tmpl`中使用定义好的`safe`模板函数：

```
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>修改模板引擎的标识符</title>
</head>
<body>
<div>{{ . | safe }}</div>
</body>
</html>


```

### 静态文件处理

当我们渲染的 HTML 文件中引用了静态文件时，我们只需要按照以下方式在渲染页面前调用`gin.Static`方法即可。

```
func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
   
	r.Run(":8080")
}


```

### 使用模板继承

Gin 框架默认都是使用单模板，如果需要使用`block template`功能，可以通过`"github.com/gin-contrib/multitemplate"`库实现，具体示例如下：

首先，假设我们项目目录下的 templates 文件夹下有以下模板文件，其中`home.tmpl`和`index.tmpl`继承了`base.tmpl`：

```
templates
├── includes
│   ├── home.tmpl
│   └── index.tmpl
├── layouts
│   └── base.tmpl
└── scripts.tmpl


```

然后我们定义一个`loadTemplates`函数如下：

```
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}


```

我们在`main`函数中

```
func indexFunc(c *gin.Context){
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func homeFunc(c *gin.Context){
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func main(){
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)
	r.Run()
}


```

### 补充文件路径处理

关于模板文件和静态文件的路径，我们需要根据公司 / 项目的要求进行设置。可以使用下面的函数获取当前执行程序的路径。

```
func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}


```

### JSON 渲染

```
func main() {
	r := gin.Default()

	
	r.GET("/someJSON", func(c *gin.Context) {
		
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	r.Run(":8080")
}


```

### XML 渲染

注意需要使用具名的结构体类型。

```
func main() {
	r := gin.Default()
	
	r.GET("/someXML", func(c *gin.Context) {
		
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreXML", func(c *gin.Context) {
		
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})
	r.Run(":8080")
}


```

### YMAL 渲染

```
r.GET("/someYAML", func(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
})


```

### protobuf 渲染

```
r.GET("/someProtoBuf", func(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "test"
	
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	
	
	c.ProtoBuf(http.StatusOK, data)
})


```

获取参数
----

### 获取 querystring 参数

`querystring`指的是 URL 中`?`后面携带的参数，例如：`/user/search?username=小王子&address=沙河`。 获取请求的 querystring 参数的方法如下：

```
func main() {
	
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		
		address := c.Query("address")
		
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}


```

### 获取 form 参数

当前端请求的数据通过 form 表单提交时，例如向`/user/search`发送一个 POST 请求，获取请求数据的方式如下：

```
func main() {
	
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		
		
		username := c.PostForm("username")
		address := c.PostForm("address")
		
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}


```

### 获取 json 参数

当前端请求的数据通过 JSON 提交时，例如向`/json`发送一个 POST 请求，则获取请求参数的方式如下：

```
r.POST("/json", func(c *gin.Context) {
	
	b, _ := c.GetRawData()  
	
	var m map[string]interface{}
	
	_ = json.Unmarshal(b, &m)

	c.JSON(http.StatusOK, m)
})


```

更便利的获取请求参数的方式，参见下面的 **参数绑定** 小节。

### 获取 path 参数

请求的参数通过 URL 路径传递，例如：`/user/search/小王子/沙河`。 获取请求 URL 路径中的参数的方式如下。

```
func main() {
	
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8080")
}


```

### 参数绑定

为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的`Content-Type`识别请求数据类型并利用反射机制自动提取请求中`QueryString`、`form表单`、`JSON`、`XML`等参数到结构体中。 下面的示例代码演示了`.ShouldBind()`强大的功能，它能够基于请求自动提取`JSON`、`form表单`和`QueryString`类型的数据，并把值绑定到指定的结构体对象。

```
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	
	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	
	router.GET("/loginForm", func(c *gin.Context) {
		var login Login
		
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	
	router.Run(":8080")
}


```

`ShouldBind`会按照下面的顺序解析请求中的数据完成绑定：

1.  如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
2.  如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。

文件上传
----

### 单个文件上传

文件上传前端页面代码：

```
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>上传文件示例</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" >
    <input type="submit" value="上传">
</form>
</body>
</html>


```

后端 gin 框架部分代码：

```
func main() {
	router := gin.Default()
	
	
	
	router.POST("/upload", func(c *gin.Context) {
		
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	router.Run()
}


```

### 多个文件上传

```
func main() {
	router := gin.Default()
	
	
	
	router.POST("/upload", func(c *gin.Context) {
		
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	router.Run()
}


```

重定向
---

### HTTP 重定向

HTTP 重定向很容易。 内部、外部重定向均支持。

```
r.GET("/test", func(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
})


```

### 路由重定向

路由重定向，使用`HandleContext`：

```
r.GET("/test", func(c *gin.Context) {
    
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "world"})
})


```

Gin 路由
------

### 普通路由

```
r.GET("/index", func(c *gin.Context) {...})
r.GET("/login", func(c *gin.Context) {...})
r.POST("/login", func(c *gin.Context) {...})


```

此外，还有一个可以匹配所有请求方法的`Any`方法如下：

```
r.Any("/test", func(c *gin.Context) {...})


```

为没有配置处理函数的路由添加处理程序，默认情况下它返回 404 代码，下面的代码为没有匹配到路由的请求都返回`views/404.html`页面。

```
r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})


```

### 路由组

我们可以将拥有共同 URL 前缀的路由划分为一个路由组。习惯性一对`{}`包裹同组的路由，这只是为了看着清晰，你用不用`{}`包裹功能上没什么区别。

```
func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {...})
		userGroup.GET("/login", func(c *gin.Context) {...})
		userGroup.POST("/login", func(c *gin.Context) {...})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
	}
	r.Run()
}


```

路由组也是支持嵌套的，例如：

```
shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
		
		xx := shopGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {...})
	}


```

通常我们将路由分组用在划分业务逻辑或划分 API 版本时。

### 路由原理

Gin 框架中的路由使用的是 [httprouter](https://github.com/julienschmidt/httprouter) 这个库。

其基本原理就是构造一个路由地址的前缀树。

Gin 中间件
-------

Gin 框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。

### 定义中间件

Gin 中的中间件必须是一个`gin.HandlerFunc`类型。

#### 记录接口耗时的中间件

例如我们像下面的代码一样定义一个统计请求耗时的中间件。

```
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") 
		
		c.Next()
		
		
		
		cost := time.Since(start)
		log.Println(cost)
	}
}


```

#### 记录响应体的中间件

我们有时候可能会想要记录下某些情况下返回给客户端的响应数据，这个时候就可以编写一个中间件来搞定。

```
type bodyLogWriter struct {
	gin.ResponseWriter               
	body               *bytes.Buffer 
}


func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  
	return w.ResponseWriter.Write(b) 
}



func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
	c.Writer = blw 

	c.Next() 

	fmt.Println("Response body: " + blw.body.String()) 
}


```

#### 跨域中间件 cors

推荐使用社区的 [https://github.com/gin-contrib/cors](https://github.com/gin-contrib/cors) 库，一行代码解决前后端分离架构下的跨域问题。

**注意：** 该中间件需要注册在业务处理函数前面。

这个库支持各种常用的配置项，具体使用方法如下。

```
package main

import (
  "time"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  
  
  
  
  
  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://foo.com"},  
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE",  "OPTIONS"},  
    AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {  
      return origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
  }))
  router.Run()
}


```

当然你可以简单的像下面的示例代码那样使用默认配置，允许所有的跨域请求。

```
func main() {
  router := gin.Default()
  
  
  
  
  router.Use(cors.Default())
  router.Run()
}


```

### 注册中间件

在 gin 框架中，我们可以为每个路由添加任意数量的中间件。

#### 为全局路由注册

```
func main() {
	
	r := gin.New()
	
	r.Use(StatCost())
	
	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) 
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}


```

#### 为某个路由单独注册

```
	r.GET("/test2", StatCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) 
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})


```

#### 为路由组注册中间件

为路由组注册中间件有以下两种写法。

写法 1：

```
shopGroup := r.Group("/shop", StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}


```

写法 2：

```
shopGroup := r.Group("/shop")
shopGroup.Use(StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}


```

### 中间件注意事项

#### gin 默认中间件

`gin.Default()`默认使用了`Logger`和`Recovery`中间件，其中：

*   `Logger`中间件将日志写入`gin.DefaultWriter`，即使配置了`GIN_MODE=release`。
*   `Recovery`中间件会 recover 任何`panic`。如果有 panic 的话，会写入 500 响应码。

如果不想使用上面两个默认的中间件，可以使用`gin.New()`新建一个没有任何默认中间件的路由。

#### gin 中间件中使用 goroutine

当在中间件或`handler`中启动新的`goroutine`时，**不能使用**原始的上下文（c *gin.Context），必须使用其只读副本（`c.Copy()`）。

运行多个服务
------

我们可以在多个端口启动服务，例如：

```
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
   
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}


```

![](https://www.liwenzhou.com/images/wxgzh_qrcode.png)