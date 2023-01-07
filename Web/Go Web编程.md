# Go Web编程

### 1 Handle请求



### 2 Handlers



### 3 HTTP消息

**结构**：`Request` 和  `Response`具有相同的结构

- 请求（响应）行
- 0个或多个Header
- 空行
- 可选的消息体（Body）

**例子**

```go
GET /Protocols/ra.html(URL) HTTP/1.1(协议版本)
Host: www.w3.org  (Header)
User-Agent: Mozila/5.0 (Header)
(空行)
```

#### 1 HTTP.Request（请求）

##### 1 `Request` 是个struct

​	代表了客户端/服务端发送的HTTP请求消息

##### 2 重要的字段

**①URL**

- 代表请求行（请求信息第一行）里面的部分内容

- URL字段是指向`url.URL`（url是一个package）类型的一个指针，`url.URL`是一个struct

  ```go
  type URL struct {
      Scheme   string
      Opaque   string    // 编码后的不透明数据
      User     *Userinfo // 用户名和密码信息
      Host     string    // host或host:port
      Path     string
      RawQuery string // 编码后的查询字符串，没有'?'
      Fragment string // 引用的片段（文档位置），没有'#'
  }
  ```

- URL的通用格式：`scheme://[userinfo@]host/path[?query][#fragment]`

  - 不以斜杆开头的URL（不透明的URL）被解释为: `scheme:opaque[?query][#fragment]`

- URL Query：RawQuery会提供实际查询的字符串

  - 如`http://www.example.com/post?id=123&thread_id=456`
  - 它的RawQuery的值就是`id=123&thread_id=456`
  - `r.URL.RawQuery`会提供实际查询的原始字符串
  - `r.URL.Query()`方法会提供查询字符串对应的`map[string][]string`，为什么使用[]string，因为url中可以有多个id

- URL Fragment

  - 如果从浏览器发出的请求，那么你无法提取出Fragment字段的值，**因为浏览器在发送请求时会把fragment部分去掉，即服务器端只收到#之前的信息**

**②Header**

- 请求和响应（Request、Response）的headers是通过Header类型来描述的，**它是一个map**，用来描述`HTTP Header`里的`Key-Value`
- `Header map`的`key`是`string`类型，`value`是`[]string`
- 设置`key`的时候会创建一个空的`[]string`作为`value`，`value`里面第一个元素就是新`header`值
- 为指定的`key`添加一个新的`header`值，执行`append`操作即可
- 获取`Header`
  - `r.Header`返回`map`
  - `r.Header["Accept-Encoding"]`返回`[gzip, deflate]([]string类型)`
  - `r.Header.Get("Accept-Encoding")`返回`gzip, deflate(string类型)`

**③Body**

- 请求和响应的`bodies`都是使用`Body`字段来表示的
- `Body`是一个`io.ReadCloser`接口
  - 一个`Reader`接口：定义了一个`Read`方法
  - 一个`Closer`接口定义了一个`Close`方法

**④Form、PostFrom、MultipartForm**

​	**读取Form的值**

- Form
- PostForm
- FormValue()
- PostFormValue()
- FormFile()
- MultipartReader()

- 表单

  ```htm
  <from action="/process" method="post" enctype="application/x-www-form-urlencoded">
  	<input type="text" name="first_name"/>
  	<input type="text" name="last_name"/>
  	<input type="submit"/>
  </from>
  ```

  - 通过表单的`method`属性，可以设置`POST`还是`GET`
  - 通过POST发送的`name-value`数据对的格式可以通过表单的`Content Type`来指定，也就是`enctype`属性
    - enctype的默认值是`application/x-www-form-urlencoded`，其会将表单数据编码到**查询字符串**里面（简单文本是使用这种模式）
    - enctype是`multipart/form-data`，那么每一个`name-value`对都会被转化为MIME消息部分（主要用于上传文件）

- Form字段

  - Form里面的数据是`key-value`对存在的
  - 先调用`ParseForm`或`ParseMultipartForm`来解析`Request`，然后相应的访问`Form、PostForm或MultipartForm`字段

- PostForm字段

  - 只会提供表单里的数据，而不会提供URL中的数据

- MultipartForm字段

  - 需要先调用`ParseMultipartForm`方法来解析`Request`，其只包含**表单的key-value对**
  - 返回类型是一个`struct`而不是`map`，但是这个`struct`里有两个`map`
    - `map1`——`key-value`:`string-[]string`
    - `map2`—— `key-value`:`string-文件` —— 用于上传文件

- `FromValue`方法和`PostFormValue`方法

  - `FromValue`方法会返回**Form**字段中指定`key`对应的第一个`value`
  - `PostFormValue`方法会返回**PostForm**字段中指定`key`对应的第一个`value`
  - 无需进行解析，即方法会自己调用`ParseMultipartForm`进行解析

**上传文件**

​	`multipart/from-data` 最常用的应用场景就是上传文件

- 首先调用`ParseMultipartForm`方法
- 从`File`字段获得`FileHeader`，调用其`Open`方法来获得文件
- 可以使用`ioutil.ReadAll`函数把文件内容读取到`byte`切片里
- 上传文件还可以使用`FromFile`方法
  - 只上传一个文件，使用这种方式会快一些

3 通过Request的方法可以访问Request中的Cookie、URL、User Agent等信息

4 Request即可**代表发送到服务器的请求，又可代表客户端发出的请求**



#### 2 HTTP.ResponseWriter（响应）

- 从服务器向客户端返回响应需要使用`ResponseWriter`
- `ResponseWriter`是一个接口，`handler`用它来返回响应
- 真正支撑`ResponseWriter`的幕后`struct`是非导出的`http.response`

**问题**

1 为什么`Handler`的`ServeHTTP(w ResponseWriter, r *Request)`，只有一个是指针类型，而w是按值传递的吗？

​	因为`ResponseWriter`是一个接口类型，因此其实际是**引用传递**

**方法**

1 `Write`方法

- Write方法接收一个byte切片作为参数，然后把它写入到HTTP响应的Body里面
- 如果Write方法被调用时，header里面没有设定content type，那么数据的前512字节就会被用来检测content type

2 `WriteHeader`方法

- `WriteHeader`方法接收一个整数类型（HTTP状态码）作为参数，并把它作为HTTP响应的状态码返回
- 如果该方法没有显式调用，那么在第一次调用`Write`方法前，会隐式的调用`WriteHeader(http.StatusOK)
- `WriteHeader`主要用来发送错误类的HTTP状态码
- 调用完`WriteHeader`方法后，仍然可以写入到`ResponseWriter`，但无法再修改`header`

3 `Header`方法

- `Header`方法返回`headers`的`map`，可以进行修改
- 修改后的`headers`将会体现再返回给客户端的`HTTP`响应里

#### 3 内置的Response

1 `NotFound`函数：包装一个404状态码和一个额外的信息

2 `ServeFile`函数：从文件系统提供文件，返回给请求者

3 `ServeContent`函数：把实现了`io.ReadSeeker`接口的任何东西里面的内容返回给请求者。其还可以处理`Range`请求（范围请求），如果只请求了资源的一部分内容，那么`ServeContent`就可以如此响应，而`ServeFile`或`io.Copy`则不行

4 `Redirect`函数：告诉客户端重定向到另外一个URL



### 4 模板

**定义**：

#### 1 模板

- Web模板就是预先设计好的`HTML`页面，它可以被模板引擎反复的使用，来产生`HTML`页面。
- Go的标准库提供了`text/template`（通用模板引擎），`html/template`（HTML模板引擎）两个模板库
- 模板必须是**可读的文本格式**，扩展名任意，对于Web应用通常就是HTML（里面会内嵌一些命令 —— action）
- `action`位于**双层花括号之间的点**：`{{ . }}`，其可以命令模板引擎将其替换成一个值

```html
<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title>Go Web Programming</title>
    </head>
</html>
<body>
    {{ . }}  # action
</body>
```

#### 2 模板引擎

- 模板引擎可以合并**模板与上下文数据**，产生最终的HTML

  ![image-20230102230926540](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301032233513.png)

##### 2.1 无逻辑模板引擎

- 通过占位符，动态数据被替换到模板中
- 不做任何逻辑处理，只做字符串替换
- 处理完全由handler来完成
- 目标是展示层和逻辑的完全分离

##### 2.2 逻辑嵌入模板引擎

- 编程语言被嵌入到模板中
- 在运行时由模板引擎来执行，也包含替换功能
- 逻辑代码遍布handler和模板，难以维护

##### 2.3 Go的模板引擎

- 主要使用的是`text/template`，`html/template`两个模板库，是个混合体。

- 模板可以完全无逻辑，但又具有足够的嵌入特性
- 和大多数模板引擎一样，Go Web的模板位于无逻辑和嵌入逻辑之间

##### 2.4 Go模板引擎的工作原理

- 在Web应用中，通常由handler来出发模板引擎
- handler调用模板引擎，并将使用的模板传递给引擎（通常是一组模板文件和动态数据）
- 模板引擎生成HTML，并将其写入到ResponseWriter
- ResponseWriter再将它加入到HTTP响应中，返回给客户端

![image-20230102231434844](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301032233128.png)

#### 3 使用模板引擎

##### 3.1 解析模板源

​	解析模板源（可以是字符串或模板文件），从而创建一个解析好的模板struct

##### 3.2 执行解析好的模板

​	执行解析好的模板，并传入`ResponseWriter`和数据，这会触发模板引擎组合解析好的模板和数据，来产生最终的`HTML`，并将它传递给`ResponseWriter`。



#### 4 解析模板

##### 4.1 ParseFile

- 解析模板文件，并创建一个解析好的模板struct，后续可被执行
- `ParseFIles`函数是`Template struct` 上`ParseFiles`方法的简便调用
- 调用`ParseFiles`后，会创建一个新的模板，模板的名字是文件名
- `ParseFiles`的参数数量可变，但只返回一个模板。当解析多个文件时，第一个文件作为返回的模板（名、内容），其余的作为map，供后续使用

##### 4.2 ParseGlob

- 使用模式匹配来解析特定的文件
- 返回找到的第一个模板的文件名作为返回的模板名

##### 4.3 Parse

- 可以解析字符串模板，其它最终方式都会调用Parse方法

##### 4.4 Lookup方法

- 通过模板名来寻找模板，如果没找到就返回nil

##### 4.5 Must函数

- 可以包裹一个函数，返回一个模板的指针和一个错误
  - 如果错误不为nil，那么就会发生panic

#### 5 执行模板

##### 5.1 Execute

- 参数是`ResponseWriter`，写入的数据
- 适用于**单模板**，只能使用模板集中的第一个模板

##### 5.2 ExecuteTemplate

- 参数是`ResponseWriter`，**模板名**，写入的数据
- 适用于**模板集**来选定要使用的模板

#### 6 Action

**定义**：

- Action就是GO模板中嵌入的命令，位于两组花括号之间{{ xxx }}
- `.` 就是一个Action，而且是最重要的一个，它代表传入模板的数据

**Action主要可分为五类**

- 条件类
- 迭代类
- 设置类
- 包含类
- 定义类



##### 5.1 条件类

```html
{{if arg }}
	some content
{{ else }}
	other content
{{ end }}
```

##### 5.2 迭代类

```html
{{ range array }}
	Dot is set to the element {{ . }}
{{ end }}
```

- 这类Action可以用来遍历数组、slice、map或channel等数据结构
- `"."`用来代表每次迭代循环中的元素

##### 5.3 设置类

```html
{{ with arg }}
	Dot is set to {{ . }}  # 将点设置为arg，而不是传入的数据 
{{ else }}
	The dot is still {{ . }} # arg为""时，执行else
{{ end }}
```

- 它允许在指定范围内，让`“.”`来表示其它指定的值

##### 5.4 包含类

```html
{{ template "name" (arg) }}
```

- 它允许在模板中包含其他的模板
- arg为向模板传入的数据

##### 5.5 定义类

```html
{{ define name }}
# 修改模板名为name
{{ end }}
```



#### 7 参数

##### 7.1 参数

- 可以在action中设置变量，变量以`$`开头：`$variable := value`

```html
{{ range $key, $value := . }}
	The key is {{ $key }} and the value is {{ $value }}
{{ end }}
```

##### 7.2 管道（pipeline）

- 管道是按顺序连接到一起的参数，函数和方法
- 如`{{p1 | p2 | p3 }}`，p1, p2, p3 要么是参数，要么是函数
- 管道允许我们把参数的输出发给下一个参数，下一个参数由管道（|）分隔开

```html
{{ 12.3456 | printf "%.2f" }}
```

##### 7.3 函数

- 参数可以是一个函数】
- GO模板引擎提供了一些基本的内置函数，功能有限
- 可自定义函数
  - 可以接收任意数量的输入参数
  - 返回一个值和一个可选错误
- 内置函数
  - define、template、block
  - html、js、urlquery
  - index
  - print、printf、println
  - len
  - with
- 如何自定义函数

```go
// 1. 创建一个FUncMap(map类型)，key是函数名，value是函数
funcMap := template.FuncMap{"fdate": formatDate}
func formtDate(t time.Tine) string {
    layout := "2006-01-02"
    return t.Format(layout)
}
// 2. 把FUncMap附加到模板
t := template.New("t1.html").FUncs(funcMap)
// 3. 在模板中使用函数
{{ . | fdate }}
```

#### 8 Layout模板

- Layout模板就是网页中固定的部分，它可以被多个网页重复使用

##### 8.1 如何制作layout模板

- 在模板文件里面使用define action再定义一个模板
- 也可在多个模板文件中，定义同名的模板

##### 8.2 使用block action定义默认模板

```htm
{{ block arg }}
	Dot is set to arg
{{ end }}
```

- block action可以定义模板，并同时使用它
- template模板必须可用
- bolck当模板不存在时会创建它



### 5 路由

#### 1 `Controller`角色

- `main()`：设置类工作
- `controller`：
  - 静态资源
  - 把不同的请求发送到不同的`controller`进行处理
  - 路由结构

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301032234292.png" alt="image-20230103223427645" style="zoom: 67%;" />



#### 2 路由参数

##### **静态路由：**一个路径对应一个页面

**带参数的路由：**根据路由参数，创建出一族不同的页面（模板相同，数据不同）



#### 3 第三方路由器

- `gorilla/mux`: 灵活性高、功能强大、性能**相对较差一些**
- `httprouter`: 注重性能、功能简单



### 6 JSON

#### 1 JSON格式与Go Struct的对比

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301032236075.png" alt="image-20230103223632200" style="zoom: 67%;" />

#### 2 `JSON`如何对应`Go struct`

##### 2.1 对象名的映射

​	`JSON`字段名一般是**小写格式**，`Go struct`为了导出，往往使用的是**大写格式**，因此往往使用`Tags`来让`JSON`对应`Go struct`中的对象。

```go
type Company struct {
    ID		int	   `json:"id"`
    Name	string `json:"name"`
    Country string `json:"country"`
}
```

##### 2.2 类型映射

- `Go bool`: `JSON boolean`
- `Go float64`: `JSON 数值`
- `Go string`: `JSON strings`
- `Go nil`: `JSON null`

##### 2.3 对于未知结构的`JSON`

- `map[string]interface{}`可以存储任意对象`JSON`对象
- `[]interface{}`可以存储任意的`JSON`数组

##### 2.4 读取`JSON`

**①使用解码器**

```go
// 创建一个解码器，传进去的参数要实现Reader接口
dec := json.NewDecoder(r.Body)
// 在解码器上进行解码，把解码好的数据放入query
dec.Decoder(&query)
```

**②使用`Unmarshal`**

- 把`json`转化为`go struct`

##### 2.5 写入`JSON`

**①使用编码器**

```go
// 创建一个编码器，传进去的参数要实现Writer接口
enc := json.NewEncoder(w)
// 在编码器上进行编码，把要编码的数据results传入到w
enc.Encoder(results)
```

**②使用`Marshal`**

- 把`go struct`转化为`json`格式
- 可以使用`MarshalIndent`，让生成的`json`文件带**缩进符**





### 7 中间件

#### 1 什么是中间件（Middleware）

- 根据请求类型，选择响应的Handler
- 处理Handler的响应
- 因此，中间件可以对传入的请求和传出的响应进行操作，增加一些功能

#### 2 中间件的用途

- `Logging`: 日志文件
- `安全`: 用户认证
- `请求超时`
- `响应压缩`: 减小响应的文件体积

#### 3 创建中间件

```go
type MyMiddleware struct {
    Next http.Handler  // 实现了Handler接口
}
// 要让MyMiddleware也实现Handler接口，就要实现ServeHTTP方法
func(m MyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // 在next handler之前做一些事情
    m.Next.ServeHTTP(w, r)
    // 在next handler之后做一些事情
}
```





### 8 请求上下文（Request Conntext）

#### 1 方法

- `func(*Request) Context() context.Context`
  - 返回当前请求的上下文
- `func(*Request) WithContext(ctx context.Context) context.Context`
  - 基于`Context`进行“修改”，（实际上）是创建一个新的`Context`

#### 2 接口定义

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)  // 什么时候失效
	Done() <-chan struct{}  // 截至时关闭channel
	Err() error  // 取消原因
	Value(key any) any  
}
// 这些方法都是只读的，不用进行设置
```

- `WihtCancel()`：它有一个`CancelFunc`
- `WithDeadline()`: 带有一个时间戳（time.Time)
- `WithTimeout()`: 带有一个具体的时间段(time.Duration)
- `WithValue()`: 在里面可以添加一些值