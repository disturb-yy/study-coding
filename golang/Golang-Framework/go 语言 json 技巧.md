> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/json-tricks/)

> 李文周的 Blog 总结 go 语言中关于 json 操作的一些技巧，主要包含 go 语言中结构体序列化（json.Unmarshal）成 JSON 格式的数据，以及将 JSON 格式数据反序列化（json.Unmarshal）成 go 语言结构体会遇到的种种问题及相应的解决办法。

本文总结了我平时在项目中遇到的那些关于 go 语言中数据与 JSON 格式之间相互转换的问题及解决办法。

### 基本的序列化

首先我们来看一下 Go 语言中`json.Marshal`（序列化）与`json.Unmarshal`（反序列化）的基本用法。

```
type Person struct {
	Name   string
	Age    int64
	Weight float64
}

func main() {
	p1 := Person{
		Name:   "七米",
		Age:    18,
		Weight: 71.5,
	}
	
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}


```

输出：

```
str:{"Name":"七米","Age":18,"Weight":71.5}
p2:main.Person{Name:"七米", Age:18, Weight:71.5}


```

### 结构体 tag 介绍

`Tag`是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 `Tag`在结构体字段的后方定义，由一对**反引号**包裹起来，具体的格式如下：

```
`key1:"value1" key2:"value2"`


```

结构体 tag 由一个或多个键值对组成。键与值使用**冒号**分隔，值用**双引号**括起来。同一个结构体字段可以设置多个键值对 tag，不同的键值对之间使用**空格**分隔。

### 使用 json tag 指定字段名

序列化与反序列化默认情况下使用结构体的字段名，我们可以通过给结构体字段添加 tag 来指定 json 序列化生成的字段名。

```
type Person struct {
	Name   string `json:"name"` 
	Age    int64
	Weight float64
}


```

### 忽略某个字段

如果你想在 json 序列化 / 反序列化的时候忽略掉结构体中的某个字段，可以按如下方式在 tag 中添加`-`。

```
type Person struct {
	Name   string `json:"name"` 
	Age    int64
	Weight float64 `json:"-"` 
}


```

### 忽略空值字段

当 struct 中的字段没有值时， `json.Marshal()` 序列化的时候不会忽略这些字段，而是默认输出字段的类型零值（例如`int`和`float`类型零值是 0，`string`类型零值是`""`，对象类型零值是 nil）。如果想要在序列序列化时忽略这些没有值的字段时，可以在对应字段添加`omitempty` tag。

举个例子：

```
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Hobby []string `json:"hobby"`
}

func omitemptyDemo() {
	u1 := User{
		Name: "七米",
	}
	
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}


```

输出结果：

```
str:{"name":"七米","email":"","hobby":null}


```

如果想要在最终的序列化结果中去掉空值字段，可以像下面这样定义结构体：

```
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
}


```

此时，再执行上述的`omitemptyDemo`，输出结果如下：

```
str:{"name":"七米"} // 序列化结果中没有email和hobby字段


```

### 忽略嵌套结构体空值字段

首先来看几种结构体嵌套的示例：

```
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
	Profile
}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

func nestedStructDemo() {
	u1 := User{
		Name:  "七米",
		Hobby: []string{"足球", "双色球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}


```

匿名嵌套`Profile`时序列化后的 json 串为单层的：

```
str:{"name":"七米","hobby":["足球","双色球"],"site":"","slogan":""}


```

想要变成嵌套的 json 串，需要改为具名嵌套或定义字段 tag：

```
type User struct {
	Name    string   `json:"name"`
	Email   string   `json:"email,omitempty"`
	Hobby   []string `json:"hobby,omitempty"`
	Profile `json:"profile"`
}



```

想要在嵌套的结构体为空值时，忽略该字段，仅添加`omitempty`是不够的：

```
type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	Profile `json:"profile,omitempty"`
}



```

还需要使用嵌套的结构体指针：

```
type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}



```

### 不修改原结构体忽略空值字段

我们需要 json 序列化`User`，但是不想把密码也序列化，又不想修改`User`结构体，这个时候我们就可以使用创建另外一个结构体`PublicUser`匿名嵌套原`User`，同时指定`Password`字段为匿名结构体指针类型，并添加`omitempty`tag，示例代码如下：

```
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User             
	Password *struct{} `json:"password,omitempty"`
}

func omitPasswordDemo() {
	u1 := User{
		Name:     "七米",
		Password: "123456",
	}
	b, err := json.Marshal(PublicUser{User: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)  
}


```

### 优雅处理字符串格式的数字

有时候，前端在传递来的 json 数据中可能会使用字符串类型的数字，这个时候可以在结构体 tag 中添加`string`来告诉 json 包从字符串中解析相应字段的数据：

```
type Card struct {
	ID    int64   `json:"id,string"`    
	Score float64 `json:"score,string"` 
}

func intAndStringDemo() {
	jsonStr1 := `{"id": "1234567","score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) 
}


```

### 整数变浮点数

因为在 JSON 协议中是没有整型和浮点型之分的，它们统称为 number。json 字符串中的数字经过 Go 语言中的 json 包反序列化之后都会成为`float64`类型。

通常这并不会有什么问题，但是在某些特殊场景下就会产生意想不到的结果。比如，将 JSON 格式的数据反序列化为`map[string]interface{}`时，数字都变成科学计数法表示的浮点数。

```
func useNumberDemo(){
	type student struct {
		ID int64 `json:"id"`
		Name string `json:"q1mi"`
	}
	s := student{ID: 123456789,Name: "q1mi"}
	b, _ := json.Marshal(s)
	var m map[string]interface{}
	
	json.Unmarshal(b, &m)
	fmt.Printf("id:%#v\n", m["id"])  
	fmt.Printf("id type:%T\n", m["id"])  

	
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	decoder.Decode(&m)
	fmt.Printf("id:%#v\n", m["id"])  
	fmt.Printf("id type:%T\n", m["id"]) 
}


```

这种问题通常出现在将 JSON 格式数据反序列化为`map[string]interface{}`时，再来一个示例。

```
func jsonDemo() {
	
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	
	var m2 map[string]interface{}
	err = json.Unmarshal(b, &m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) 
	fmt.Printf("type:%T\n", m2["count"])  
}


```

这种场景下如果想更合理的处理数字就需要使用`decoder`去反序列化，示例代码如下：

```
func decoderDemo() {
	
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	
	var m2 map[string]interface{}
	
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) 
	fmt.Printf("type:%T\n", m2["count"])  
	
	count, err := m2["count"].(json.Number).Int64()
	if err != nil {
		fmt.Printf("parse to int64 failed, err:%v\n", err)
		return
	}
	fmt.Printf("type:%T\n", int(count)) 
}


```

`json.Number`的源码定义如下：

```
type Number string


func (n Number) String() string { return string(n) }


func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}


func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}


```

我们在处理 number 类型的 json 字段时需要先得到`json.Number`类型，然后根据该字段的实际类型调用`Float64()`或`Int64()`。

### 自定义解析时间字段

Go 语言内置的 json 包使用 `RFC3339` 标准中定义的时间格式，对我们序列化时间字段的时候有很多限制。

```
type Post struct {
	CreateTime time.Time `json:"create_time"`
}

func timeFieldDemo() {
	p1 := Post{CreateTime: time.Now()}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}


```

上面的代码输出结果如下：

```
str:{"create_time":"2020-04-05T12:28:06.799214+08:00"}
json.Unmarshal failed, err:parsing time ""2020-04-05 12:25:42"" as ""2006-01-02T15:04:05Z07:00"": cannot parse " 12:25:42"" as "T"


```

也就是内置的 json 包不识别我们常用的字符串时间格式，如`2020-04-05 12:25:42`。

不过我们通过实现 `json.Marshaler`/`json.Unmarshaler` 接口实现自定义的事件格式解析。

```
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Post struct {
	CreateTime CustomTime `json:"create_time"`
}

func timeFieldDemo() {
	p1 := Post{CreateTime: CustomTime{time.Now()}}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}


```

### 自定义 MarshalJSON 和 UnmarshalJSON 方法

上面那种自定义类型的方法稍显啰嗦了一点，下面来看一种相对便捷的方法。

首先你需要知道的是，如果你能够为某个类型实现了`MarshalJSON()([]byte, error)`和`UnmarshalJSON(b []byte) error`方法，那么这个类型在序列化（MarshalJSON）/ 反序列化（UnmarshalJSON）时就会使用你定制的相应方法。

```
type Order struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedTime time.Time `json:"created_time"`
}

const layout = "2006-01-02 15:04:05"


func (o *Order) MarshalJSON() ([]byte, error) {
	type TempOrder Order 
	return json.Marshal(struct {
		CreatedTime string `json:"created_time"`
		*TempOrder         
	}{
		CreatedTime: o.CreatedTime.Format(layout),
		TempOrder:   (*TempOrder)(o),
	})
}


func (o *Order) UnmarshalJSON(data []byte) error {
	type TempOrder Order 
	ot := struct {
		CreatedTime string `json:"created_time"`
		*TempOrder         
	}{
		TempOrder: (*TempOrder)(o),
	}
	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}
	var err error
	o.CreatedTime, err = time.Parse(layout, ot.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}


func customMethodDemo() {
	o1 := Order{
		ID:          123456,
		Title:       "《七米的Go学习笔记》",
		CreatedTime: time.Now(),
	}
	
	b, err := json.Marshal(&o1)
	if err != nil {
		fmt.Printf("json.Marshal o1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	
	jsonStr := `{"created_time":"2020-04-05 10:18:20","id":123456,"title":"《七米的Go学习笔记》"}`
	var o2 Order
	if err := json.Unmarshal([]byte(jsonStr), &o2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("o2:%#v\n", o2)
}


```

输出结果：

```
str:{"created_time":"2020-04-05 10:32:20","id":123456,"title":"《七米的Go学习笔记》"}
o2:main.Order{ID:123456, Title:"《七米的Go学习笔记》", CreatedTime:time.Time{wall:0x0, ext:63721678700, loc:(*time.Location)(nil)}}


```

### 使用匿名结构体添加字段

使用内嵌结构体能够扩展结构体的字段，但有时候我们没有必要单独定义新的结构体，可以使用匿名结构体简化操作：

```
type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func anonymousStructDemo() {
	u1 := UserInfo{
		ID:   123456,
		Name: "七米",
	}
	
	b, err := json.Marshal(struct {
		*UserInfo
		Token string `json:"token"`
	}{
		&u1,
		"91je3a4s72d1da96h",
	})
	if err != nil {
		fmt.Printf("json.Marsha failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	
}


```

### 使用匿名结构体组合多个结构体

同理，也可以使用匿名结构体来组合多个结构体来序列化与反序列化数据：

```
type Comment struct {
	Content string
}

type Image struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func anonymousStructDemo2() {
	c1 := Comment{
		Content: "永远不要高估自己",
	}
	i1 := Image{
		Title: "赞赏码",
		URL:   "https://www.liwenzhou.com/images/zanshang_qr.jpg",
	}
	
	b, err := json.Marshal(struct {
		*Comment
		*Image
	}{&c1, &i1})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	
	jsonStr := `{"Content":"永远不要高估自己","title":"赞赏码","url":"https://www.liwenzhou.com/images/zanshang_qr.jpg"}`
	var (
		c2 Comment
		i2 Image
	)
	if err := json.Unmarshal([]byte(jsonStr), &struct {
		*Comment
		*Image
	}{&c2, &i2}); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c2:%#v i2:%#v\n", c2, i2)
}


```

输出：

```
str:{"Content":"永远不要高估自己","title":"赞赏码","url":"https://www.liwenzhou.com/images/zanshang_qr.jpg"}
c2:main.Comment{Content:"永远不要高估自己"} i2:main.Image{Title:"赞赏码", URL:"https://www.liwenzhou.com/images/zanshang_qr.jpg"}


```

### 处理不确定层级的 json

如果 json 串没有固定的格式导致不好定义与其相对应的结构体时，我们可以使用`json.RawMessage`原始字节数据保存下来。

```
type sendMsg struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}

func rawMessageDemo() {
	jsonStr := `{"sendMsg":{"user":"q1mi","msg":"永远不要高估自己"},"say":"Hello"}`
	
	var data map[string]json.RawMessage
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Printf("json.Unmarshal jsonStr failed, err:%v\n", err)
		return
	}
	var msg sendMsg
	if err := json.Unmarshal(data["sendMsg"], &msg); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("msg:%#v\n", msg)
	
}


```

### 序列化时不转义

json 包中的`encoder`可以通过`SetEscapeHTML`指定是否应该在 JSON 字符串中转义有问题的 HTML 字符。其默认行为是将`&`、`<`和`>`转义为`\u0026`、`\u003c`和`\u003e`，以避免在 HTML 中嵌入 JSON 时可能出现的某些安全问题。

如果是非 HTML 场景下不想被转义，可以通过`SetEscapeHTML(false)`禁用此行为。

例如有些业务场景下可能需要序列化带查询参数的 URL，这种场景下我们并不希望转义`&`符号。

```
type URLInfo struct {
	URL string
	
}





func JSONEncodeDontEscapeHTML(data URLInfo) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json.Marshal(data) failed, err:%v\n", err)
	}
	fmt.Printf("json.Marshal(data) result:%s\n", b)

	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false) 
	if err := encoder.Encode(data); err != nil {
		fmt.Printf("encoder.Encode(data) failed, err:%v\n", err)
	}
	fmt.Printf("encoder.Encode(data) result:%s\n", buf.String())
}


```

输出：

```
json.Marshal(data) result:{"URL":"https://liwenzhou.com?name=q1mi\u0026age=18"}
encoder.Encode(data) result:{"URL":"https://liwenzhou.com?}


```

参考链接：

[https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format](https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format)

[https://colobu.com/2017/06/21/json-tricks-in-Go/](https://colobu.com/2017/06/21/json-tricks-in-Go/)

[https://stackoverflow.com/questions/11066946/partly-json-unmarshal-into-a-map-in-go](https://stackoverflow.com/questions/11066946/partly-json-unmarshal-into-a-map-in-go)

[http://choly.ca/post/go-json-marshalling/](http://choly.ca/post/go-json-marshalling/)

![](https://www.liwenzhou.com/images/wxgzh_qrcode.png)