

# HTML



## 基础入门

**学习网站**：[HTML 教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/html/html-tutorial.html)

**参考手册**：

- [HTML 标签速查列表](https://www.runoob.com/html/html-quicklist.html)：**重要**
- [HTML 标签简写及全称](https://www.runoob.com/html/html-tag-name.html)

- [HTML 标签参考手册](https://www.runoob.com/tags/html-reference.html)

- [HTML 标准属性参考手册](https://www.runoob.com/tags/ref-standardattributes.html)

  





### 1 HTML 简介

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>菜鸟教程(runoob.com)</title>
</head>
<body>
 
<h1>我的第一个标题</h1>
 
<p>我的第一个段落。</p>
 
</body>
</html>
```

*对于中文网页需要使用* **<meta charset="utf-8">** *声明编码，否则会出现乱码。有些浏览器(如 360 浏览器)会设置 GBK 为默认编码，则你需要设置为* **<meta charset="gbk">。**

<img src="https://www.runoob.com/wp-content/uploads/2013/06/02A7DD95-22B4-4FB9-B994-DDB5393F7F03.jpg" alt="img" style="zoom:50%;" />

**什么是HTML?**

HTML 是用来描述网页的一种语言。

- HTML 指的是超文本标记语言: **H**yper**T**ext **M**arkup **L**anguage
- HTML 不是一种编程语言，而是一种**标记**语言
- 标记语言是一套**标记标签** (markup tag)
- HTML 使用标记标签来**描述**网页
- HTML 文档包含了HTML **标签**及**文本**内容
- HTML文档也叫做 **web 页面**



### 2 HTML 基础



#### 标题

HTML 标题（Heading）是通过<h1> - <h6> 标签来定义的。



#### 段落

HTML 段落是通过标签 <p> 来定义的。



#### 链接

HTML 链接是通过标签 <a> 来定义的。

```html
// 在 href 属性中指定链接的地址
<a href="https://www.runoob.com">这是一个链接</a>
```



#### 图像

HTML 图像是通过标签 <img> 来定义的

```html
<img decoding="async" src="/images/logo.png" width="258" height="39" />
```



### 3 HTML 元素

| 开始标签 *             | 元素内容     | 结束标签 * |
| :--------------------- | :----------- | :--------- |
| <p>                    | 这是一个段落 | </p>       |
| <a href="default.htm"> | 这是一个链接 | </a>       |
| <br>                   | 换行         |            |

开始标签常被称为**起始标签（opening tag）**，结束标签常称为**闭合标签（closing tag）**



#### 语法

- HTML 元素以**开始标签**起始
- HTML 元素以**结束标签**终止
- **元素的内容**是开始标签与结束标签之间的内容
- 某些 HTML 元素具有**空内容（empty content）**
- 空元素**在开始标签中进行关闭**（以开始标签的结束而结束）
- 大多数 HTML 元素可拥有**属性**



#### HTML 空元素

**没有内容的 HTML 元素被称为空元素**。空元素是在开始标签中关闭的。

\<br> 就是没有关闭标签的空元素（\<br> 标签定义换行）。

在 XHTML、XML 以及未来版本的 HTML 中，所有元素都必须被关闭。

在开始标签中添加斜杠，比如 \<br />，是关闭空元素的正确方法，HTML、XHTML 和 XML 都接受这种方式。

**使用\<br />其实是更长远的保障。**





### 4 HTML 属性



#### HTML 属性

- HTML 元素可以设置**属性**
- 属性可以在元素中添加**附加信息**
- 属性一般描述于**开始标签**
- 属性**总是以名称/值对的形式出现**，**比如：name="value"**。



#### HTML 属性常用引用属性值

属性值应该始终被包括在引号内。

双引号是最常用的，不过使用单引号也没有问题。

**提示:** 在某些个别的情况下，比如属性值本身就含有双引号，那么您必须使用单引号，例如

```html
name='John "ShotGun" Nelson'
```



### 5 HTML 标题



#### HTML 水平线

`<hr />`标签在 HTML 页面中创建水平线。

hr 元素可用于分隔内容。



#### HTML 注释

```html
<!-- 这是一个注释 -->
```

​	开始括号之后（左边的括号）需要紧跟一个叹号 **!** (英文标点符号)，结束括号之前（右边的括号）不需要，合理地使用注释可以对未来的代码编辑工作产生帮助。



#### HTML 段落

​	可以将文档分割为若干段落



#### HTML 输出

对于 HTML，您无法通过在 HTML 代码中添加额外的空格或换行来改变输出的效果。

当显示页面时，**浏览器会移除源代码中多余的空格和空行**。所有连续的空格或空行都会被算作一个空格。**需要注意的是，HTML 代码中的所有连续的空行（换行）也被显示为一个空格。**





### 6 HTML 文本格式化

- 粗体：bold（`<b></b>`）
- 斜体：italic（`<i></i>`）

```html
<b>加粗文本</b>
<i>斜体文本</i>
<code>电脑自动输出</code>
这是 <sub> 下标</sub> 和 <sup> 上标</sup>

```

详细的文本格式化标签：[HTML 文本格式化](https://www.runoob.com/html/html-formatting.html)





### 7 HTML 链接

​	href 属性描述了链接的目标

```html
<a href="url">链接文本</a>
```



#### target 属性

```html
<a href="https://www.runoob.com/" target="_blank" rel="noopener noreferrer">访问菜鸟教程!</a>
```



#### id 属性

```html
<!-- 为文本添加标记 -->
<a id="tips">有用的提示部分</a>
	在HTML文档中创建一个链接到"有用的提示部分(id="tips"）"：
<!-- 跳转到对应的标记 -->
<a href="#tips">访问有用的提示部分</a>
	或者，从另一个页面创建一个链接到"有用的提示部分(id="tips"）"：
```



**Notice**：请始终将正斜杠添加到子文件夹。假如这样书写链接：href="https://www.runoob.com/html"，就会向服务器产生两次 HTTP 请求。这是因为服务器会添加正斜杠到这个地址，然后创建一个新的请求，就像这样：href="https://www.runoob.com/html/"。



### 8 HTML 头部



#### HTML `<head>` 元素

`<head>` 元素包含了所有的头部标签元素。在 `<head>`元素中你可以插入脚本（scripts）, 样式文件（CSS），及各种meta信息。

可以添加在头部区域的元素标签为: `<title>`, `<style>`, `<meta>`, `<link>`, `<script>`, `<noscript>` 和 <base



#### HTML `<title>` 元素

`<title> `标签定义了不同文档的标题。

`<title>` 在 HTML/XHTML 文档中是必需的。

`<title>` 元素:

- 定义了浏览器工具栏的标题
- 当网页添加到收藏夹时，显示在收藏夹中的标题
- 显示在搜索引擎结果页面的标题



#### HTML` <base> `元素

<base> 标签描述了基本的链接地址/链接目标，该标签作为HTML文档中所有的链接标签的默认链接:

```html
<head>
<base href="http://www.runoob.com/images/" target="_blank">
</head>
```



#### HTML `<link> `元素

`<link>` 标签定义了文档与外部资源之间的关系。

`<link>` 标签通常用于链接到样式表:

```html
<head>
<link rel="stylesheet" type="text/css" href="mystyle.css">
</head>
```



#### HTML `<style>` 元素

`<style>` 标签定义了`HTML`文档的样式文件引用地址.

在`<style>` 元素中你也可以直接添加样式来渲染 HTML 文档:

```html
<head>
<style type="text/css">
body {
    background-color:yellow;
}
p {
    color:blue
}
</style>
</head>
```



#### HTML `<meta> `元素

meta标签描述了一些基本的元数据。

`<meta>` 标签提供了元数据，元数据也不显示在页面上，但会被浏览器解析。

META 元素通常用于指定网页的描述，关键词，文件的最后修改时间，作者，和其他元数据。

元数据可以使用于浏览器（如何显示内容或重新加载页面），搜索引擎（关键词），或其他Web服务。

`<meta>` 一般放置于 `<head>` 区域

```html
为搜索引擎定义关键词:
<meta name="keywords" content="HTML, CSS, XML, XHTML, JavaScript">
为网页定义描述内容:
<meta name="description" content="免费 Web & 编程 教程">
定义网页作者:
<meta name="author" content="Runoob">
每30秒钟刷新当前页面:
<meta http-equiv="refresh" content="30">
```



#### HTML `<script> `元素

`<script>`标签用于加载脚本文件，如： JavaScript。

`<script>` 元素在以后的章节中会详细描述。





### 9 HTML 图像



#### 图像标签（` <img>`）和源属性（Src）

在 HTML 中，图像由`<img>` 标签定义。

`<img> `是空标签，意思是说，它只包含属性，并且没有闭合标签。

要在页面上显示图像，你需要使用源属性（src）。src 指 "source"。源属性的值是图像的 URL 地址。

**定义图像的语法是：**

```html
<img src="url" alt="some_text">
```

URL 指存储图像的位置。如果名为 "pulpit.jpg" 的图像位于 www.runoob.com 的 images 目录中，那么其 URL 为 [http://www.runoob.com/images/pulpit.jpg](https://www.runoob.com/images/pulpit.jpg)。

浏览器将图像显示在文档中图像标签出现的地方。如果你将图像标签置于两个段落之间，那么浏览器会首先显示第一个段落，然后显示图片，最后显示第二段。





#### HTML 图像- Alt属性

alt 属性用来为图像定义一串预备的可替换的文本。

替换文本属性的值是用户定义的。

```html
<img src="boat.gif" alt="Big Boat">
```

在浏览器无法载入图像时，替换文本属性告诉读者她们失去的信息。此时，浏览器将显示这个替代性的文本而不是图像。为页面上的图像都加上替换文本属性是个好习惯，这样有助于更好的显示信息，并且对于那些使用纯文本浏览器的人来说是非常有用的。



#### HTML 图像- 设置图像的高度与宽度

height（高度） 与 width（宽度）属性用于设置图像的高度与宽度。

属性值默认单位为像素:

```html
<img src="pulpit.jpg" alt="Pulpit rock" width="304" height="228">
```

**提示:** 指定图像的高度和宽度是一个很好的习惯。如果图像指定了高度宽度，页面加载时就会保留指定的尺寸。如果没有指定图片的大小，加载页面时有可能会破坏HTML页面的整体布局。



### 10 HTML 表格



表格由` <table>`标签来定义。每个表格均有若干行（由 `<tr>` 标签定义），每行被分割为若干单元格（由 `<td>` 标签定义）。字母 td 指表格数据（table data），即数据单元格的内容。

数据单元格可以包含文本、图片、列表、段落、表单、水平线、表格等等。



#### 边框属性

如果不定义边框属性，表格将不显示边框。有时这很有用，但是大多数时候，我们希望显示边框。

使用边框属性来显示一个带有边框的表格：

```html
<table border="1">
    <tr>
        <td>Row 1, cell 1</td>
        <td>Row 1, cell 2</td>
    </tr>
</table>
```



#### HTML 表格表头

表格的表头使用` <th>`标签进行定义。

大多数浏览器会把表头显示为粗体居中的文本：



**更多的表格标签**：[HTML 表格](https://www.runoob.com/html/html-tables.html)



### 11 HTML 列表



#### HTML无序列表

无序列表是一个项目的列表，此列项目使用粗体圆点（典型的小黑圆圈）进行标记。

无序列表使用` <ul>` 标签，表项则用`<li>`



#### HTML 有序列表

同样，有序列表也是一列项目，列表项目使用数字进行标记。 有序列表始于`<ol>` 标签。每个列表项始于` <li>` 标签。

列表项使用数字来标记。



#### HTML 自定义列表

自定义列表不仅仅是一列项目，而是项目及其注释的组合。

自定义列表以 <dl> 标签开始。每个自定义列表项以 <dt> 开始。每个自定义列表项的定义以 <dd> 开始。

```html
<dl>
<dt>Coffee</dt>
<dd>- black hot drink</dd>
<dt>Milk</dt>
<dd>- white cold drink</dd>
</dl>
```

浏览器显示如下：

`-` Coffee
  `- `black hot drink
`-` Milk
  `-` white cold drink



### 12 HTML 区块

HTML 可以通过` <div> `和 `<span>`将元素组合起来。

 

#### HTML 区块元素

大多数 HTML 元素被定义为**块级元素**或**内联元素**。

块级元素在浏览器显示时，通常会以新行来开始（和结束）。

实例: `<h1>`, `<p>`, `<ul>`, `<table>`



#### HTML 内联元素

内联元素在显示时通常不会以新行开始。

实例:` <b>`, `<td>`, `<a>`,` <img>`



#### HTML` <div>` 元素

HTML `<div>` 元素是块级元素，它可用于组合其他 HTML 元素的容器。

<div> 元素没有特定的含义。除此之外，由于它属于块级元素，浏览器会在其前后显示折行。

如果与 CSS 一同使用，`<div>` 元素可用于对大的内容块设置样式属性。

`<div>` 元素的另一个常见的用途是文档布局。它取代了使用表格定义布局的老式方法。使用 `<table>` 元素进行文档布局不是表格的正确用法。`<table>` 元素的作用是显示表格化的数据。



#### HTML `<span>` 元素

HTML `<span> `元素是内联元素，可用作文本的容器

`<span>` 元素也没有特定的含义。

当与 CSS 一同使用时，`<span>` 元素可用于为部分文本设置样式属性。





### HTML 布局



#### 使用`<div>` 元素

div 元素是用于分组 HTML 元素的块级元素。**可以理解成`div`生成一处区域，可以对区域的属性进行设置**。

下面的例子使用五个 div 元素来创建多列布局：

```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
</head>
<body>
 
<div id="container" style="width:500px">
 
<div id="header" style="background-color:#FFA500;">
<h1 style="margin-bottom:0;">主要的网页标题</h1></div>
 
<div id="menu" style="background-color:#FFD700;height:200px;width:100px;float:left;">
<b>菜单</b><br>
HTML<br>
CSS<br>
JavaScript</div>
 
<div id="content" style="background-color:#EEEEEE;height:200px;width:400px;float:left;">
内容在这里</div>
 
<div id="footer" style="background-color:#FFA500;clear:both;text-align:center;">
版权 © runoob.com</div>
 
</div>
 
</body>
</html>
```

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301102338879.jpeg" alt="img" style="zoom:50%;" />





### HTML 表单和输入

- HTML 表单用于收集用户的输入信息。

- HTML 表单表示文档中的一个区域，此区域包含交互控件，将用户收集到的信息发送到 Web 服务器。

**更多的HTML表单标签**：[HTML 表单](https://www.runoob.com/html/html-forms.html)



#### HTML 表单

**表单是一个包含表单元素的区域。**

表单元素是允许用户在表单中输入内容，比如：文本域（textarea）、下拉列表（select）、单选框（radio-buttons）、复选框（checkbox） 等等。

我们可以使用 `<form>` 标签来创建表单:



#### 输入元素

多数情况下被用到的表单标签是输入标签` **<input>`**。

输入类型是由 **type** 属性定义。



#### 输入类型



##### 文本域（Text Fields）

​	文本域通过 `<input type="text">` 标签来设定，当用户要在表单中键入字母、数字等内容时，就会用到文本域。

**注意:**表单本身并不可见。同时，在大多数浏览器中，文本域的默认宽度是 20 个字符。



##### 密码字段（password）

​	密码字段通过标签 `<input type="password">` 来定义:

**注意：**密码字段字符不会明文显示，而是以星号 ***** 或圆点 **.** 替代。



##### 单选按钮（Radio Buttons）

​	`<input type="radio">` 标签定义了表单的单选框选项:



##### 复选框（Checkboxes）

​	`<input type="checkbox">` 定义了复选框。

​	**复选框可以选取一个或多个选项：**



##### 文本框（textarea）

​	本例演示如何创建文本域（多行文本输入控件）。用户可在文本域中写入文本。可写入字符的字数不受限制。

```html
<textarea rows="10" cols="30">
我是一个文本框。
</textarea>
```



##### 提交按钮(Submit)

`<input type="submit">` 定义了提交按钮。

当用户单击确认按钮时，表单的内容会被传送到服务器。表单的动作属性 **action** 定义了服务端的文件名。

**action** 属性会对接收到的用户输入数据进行相关的处理:

```html
<form name="input" action="html_form_action.php" method="get">
Username: <input type="text" name="user">
<input type="submit" value="Submit">
</form>
```

​	假如您在上面的文本框内键入几个字母，然后点击确认按钮，那么输入数据会传送到 **html_form_action.php** 文件，该页面将显示出输入的结果。

以上实例中有一个 method 属性，它用于定义表单数据的提交方式，可以是以下值：

- **post**：指的是 HTTP POST 方法，表单数据会包含在表单体内然后发送给服务器，用于提交敏感数据，如用户名与密码等。
- **get**：默认值，指的是 HTTP GET 方法，表单数据会附加在 **action** 属性的 URL 中，并以 **?**作为分隔符，一般用于不敏感信息，如分页等。例如：https://www.runoob.com/?page=1，这里的 page=1 就是 get 方法提交的数据。



​	以下实例创建了一个表单，包含一个普通输入框和一个密码输入框：

```html
<form action="">
Username: <input type="text" name="user"><br>
Password: <input type="password" name="password">
</form>
```





### HTML 框架



**iframe语法:**

```html
<iframe src="URL"></iframe>
```

该URL指向不同的网页。



#### iframe - 设置高度与宽度

height 和 width 属性用来定义iframe标签的高度与宽度。

属性默认以像素为单位, 但是你可以指定其按比例显示 (如："80%")。



#### iframe - 移除边框

frameborder 属性用于定义iframe表示是否显示边框。

设置属性值为 "0" 移除iframe的边框:



#### 使用 iframe 来显示目标链接页面

iframe 可以显示一个目标链接的页面

目标链接的属性必须使用 iframe 的属性，如下实例:

```html
<iframe src="demo_iframe.htm" name="iframe_a"></iframe>
<p><a href="https://www.runoob.com" target="iframe_a" rel="noopener">RUNOOB.COM</a></p>
```

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301110007901.png" alt="image-20230111000732955" style="zoom: 67%;" />





### HTML 颜色

------

HTML 颜色由红色、绿色、蓝色混合而成。



#### 颜色值

HTML 颜色由一个十六进制符号来定义，这个符号由红色、绿色和蓝色的值组成（RGB）。

每种颜色的最小值是0（十六进制：#00）。最大值是255（十六进制：#FF）

十六进制值的写法为 # 号后跟三个或六个十六进制字符。

**三位数表示法为：#RGB，转换为6位数表示为：#RRGGBB。**

**RGB表示**：**rgb(0,0,0)**



#### 颜色名

**目前所有浏览器都支持以下颜色名。**

141个颜色名称是在HTML和CSS颜色规范定义的（17标准颜色，再加124）。下表列出了所有颜色的值，包括十六进制值。

![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301110011566.gif) **提示:** 17标准颜色：黑色，蓝色，水，紫红色，灰色，绿色，石灰，栗色，海军，橄榄，橙，紫，红，白，银，蓝绿色，黄色。点击其中一个颜色名称（或一个十六进制值）就可以查看与不同文字颜色搭配的背景颜色。

[HTML 颜色名 | 菜鸟教程 (runoob.com)](https://www.runoob.com/html/html-colornames.html)



### HTML 脚本

----------------

JavaScript 使 HTML 页面具有更强的动态和交互性。



#### HTML `<script>` 标签

`<script>` 标签用于定义客户端脚本，比如 JavaScript。

`<script>` 元素既可包含脚本语句，也可通过 src 属性指向外部脚本文件。

JavaScript 最常用于图片操作、表单验证以及内容动态更新。

​	下面的脚本会向浏览器输出"Hello World!"：

```html
<script>
document.write("Hello World!");
</script>
```



#### HTML`<noscript>` 标签

`<noscript>` 标签提供无法使用脚本时的替代内容，比方在浏览器禁用脚本时，或浏览器不支持客户端脚本时。

`<noscript>`元素可包含普通 HTML 页面的 body 元素中能够找到的所有元素。

**只有在浏览器不支持脚本或者禁用脚本时，才会显示` <noscript>` 元素中的内容：**





### HTML 字符实体

------

HTML 中的预留字符必须被替换为字符实体。

一些在键盘上找不到的字符也可以使用字符实体来替换。

**在 HTML 中，某些字符是预留的。**

在 HTML 中不能使用小于号（<）和大于号（>），这是因为浏览器会误认为它们是标签。

如果希望正确地显示预留字符，我们必须在 HTML 源代码中使用字符实体（character entities）。 字符实体类似这样：

```html
&entity_name;
或
&#entity_number;
```

如需显示小于号，我们必须这样写：`&lt;` 或 `&#60;`或 `&#060;`

![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301110024422.gif)**提示：** 使用实体名而不是数字的好处是，名称易于记忆。不过坏处是，浏览器也许并不支持所有实体名称（对实体数字的支持却很好)。



#### 不间断空格(Non-breaking Space)

​	HTML 中的常用字符实体是不间断空格(`&nbsp;`)。

​	浏览器总是会截短 HTML 页面中的空格。如果您在文本中写 10 个空格，在显示该页面之前，浏览器会删除它们中的 9 个。如需在页面中增加空格的数量，您需要使用 `&nbsp; `字符实体。



#### 结合音标符

发音符号是加到字母上的一个"glyph(字形)"。

一些变音符号, 如 尖音符 ( ̀) 和 抑音符 ( ́) 。

变音符号可以出现字母的上面和下面，或者字母里面，或者两个字母间。

变音符号可以与字母、数字字符的组合来使用。



#### HTML字符实体

| ![Note](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301110030366.jpeg)  实体名称对大小写敏感！ |      |
| ------------------------------------------------------------ | ---- |

| 显示结果 | 描述   | 实体名称 | 实体编号 |
| :------- | :----- | :------- | :------- |
|          | 空格   | `&nbsp;` | `&#160;` |
| <        | 小于号 | `&lt`;   | `&#60`;  |
| >        | 大于号 | `&gt`;   | `&#62`;  |
| &        | 和号   | `&amp`;  | `&#38;`  |
| "        | 引号   | `&quot`; | `&#34;`  |

详细看：[HTML ISO-8859-1 参考手册 | 菜鸟教程 (runoob.com)](https://www.runoob.com/tags/ref-entities.html)





### HTML 统一资源定位器(Uniform Resource Locators)

URL 是一个网页地址。URL可以由字母组成，如"runoob.com"，或互联网协议（IP）地址： 192.68.20.50。大多数人进入网站使用网站域名来访问，因为 名字比数字更容易记住。



#### URL - 统一资源定位器

Web浏览器通过URL从Web服务器请求页面。

当您点击 HTML 页面中的某个链接时，对应的` <a> `标签指向万维网上的一个地址。

一个统一资源定位器(URL) 用于定位万维网上的文档。

一个网页地址实例: http://www.runoob.com/html/html-tutorial.html 语法规则:

**scheme`://`host.domain`:`port`/`path`/`filename**

说明:

- - scheme - 定义因特网服务的类型。最常见的类型是 http
  - host - 定义域主机（http 的默认主机是 www）
  - domain - 定义因特网域名，比如 runoob.com
  - port - 定义主机上的端口号（http 的默认端口号是 80）
  - path - 定义服务器上的路径（如果省略，则文档必须位于网站的根目录中）。
  - filename - 定义文档/资源的名称



#### 常见的 URL Scheme

以下是一些URL scheme：

| Scheme | 访问               | 用于...                             |
| :----- | :----------------- | :---------------------------------- |
| http   | 超文本传输协议     | 以 http:// 开头的普通网页。不加密。 |
| https  | 安全超文本传输协议 | 安全网页，加密所有信息交换。        |
| ftp    | 文件传输协议       | 用于将文件下载或上传至网站。        |
| file   |                    | 您计算机上的文件。                  |



#### URL 字符编码

URL 只能使用 [ASCII 字符集](https://www.runoob.com/tags/html-ascii.html).

来通过因特网进行发送。由于 URL 常常会包含 ASCII 集合之外的字符，URL 必须转换为有效的 ASCII 格式。

URL 编码使用 "%" 其后跟随两位的十六进制数来替换非 ASCII 字符。

URL 不能包含空格。URL 编码通常使用 + 来替换空格。



