# CSS



## 参考手册

- [CSS 实例 ](https://www.runoob.com/css/css-examples.html) ⏱️

- [CSS 属性](https://www.runoob.com/cssref/css-reference.html)

- [CSS 选择器参考手册](https://www.runoob.com/cssref/css-selectors.html)

- [CSS 声音参考手册](https://www.runoob.com/cssref/css-ref-aural.html)

- [CSS 单位](https://www.runoob.com/cssref/css-units.html)

- [CSS 颜色参考手册](https://www.runoob.com/cssref/css-colors.html)

  



## CSS 基础

**CSS** (Cascading Style Sheets，层叠样式表），是一种用来为结构化文档（如 HTML 文档或 XML 应用）添加样式（字体、间距和颜色等）的计算机语言，**CSS** 文件扩展名为 **.css**。通过使用 **CSS** 我们可以大大提升网页开发的工作效率！



### 简介



#### 什么是 CSS?

- CSS 指层叠样式表 (**C**ascading **S**tyle **S**heets)
- 样式定义**如何显示** HTML 元素
- 样式通常存储在**样式表**中
- 把样式添加到 HTML 4.0 中，是为了**解决内容与表现分离的问题**
- **外部样式表**可以极大提高工作效率
- 外部样式表通常存储在 **CSS 文件**中
- 多个样式定义可**层叠**为一个



### 语法



CSS 规则由两个主要的部分构成：选择器，以及一条或多条声明:

- 可以把`CSS`理解成`json`的格式，选择器是一个`key`，其`value`由多个声明组成
- 在声明中，属性是`key`，值是`value`，声明之间用**分号**隔开

![img](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111212275.jpeg)

选择器通常是您需要改变样式的 HTML 元素。

每条声明由一个属性和一个值组成。

属性（property）是您希望设置的样式属性（style attribute）。每个属性有一个值。属性和值被冒号分开。



#### 实例

CSS声明总是以分号` ; `结束，声明总以大括号 **{}** 括起来:

```
p {color:red;text-align:center;}
```



#### 注释

注释是用来解释你的代码，并且可以随意编辑它，浏览器会忽略它。

CSS注释以 **/\*** 开始, 以 ***/** 结束, 实例如下:

```css
/*这是个注释*/
p
{
    text-align:center;
    /*这是另一个注释*/
    color:black;
    font-family:arial;
}
```



-----------------------------------



### id 和 class 选择器

如果你要在HTML元素中设置CSS样式，你需要在元素中设置"id" 和 "class"选择器。



#### id 选择器

**id 选择器可以为标有特定 id 的 HTML 元素指定特定的样式。HTML元素以id属性来设置id选择器，CSS 中 id 选择器以 "#" 来定义。**

以下的样式规则应用于元素属性 id="para1":

```css
<style>
#para1
{
	text-align:center;
	color:red;
} 
</style>

<p id="para1">Hello World!</p>
<p>这个段落不受该样式的影响。</p>
```

![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111222777.gif) ID属性不要以数字开头，数字开头的ID在 Mozilla/Firefox 浏览器中不起作用。



#### class 选择器

class 选择器用于**描述一组元素的样式**，class 选择器有别于id选择器，class可以在多个元素中使用。**class 选择器在 HTML 中以 class 属性表示, 在 CSS 中，类选择器以一个点 . 号显示：**

在以下的例子中，所有拥有 center 类的 HTML 元素均为居中。

```csss
<style>
.center
{
	text-align:center;
}
</style>

<h1 class="center">标题居中</h1>
<p class="center">段落居中。</p> 
```

你也可以指定特定的 HTML 元素使用 class。

在以下实例中, 所有的 p 元素使用 class="center" 让该元素的文本居中:

```CSS
p.center
{
	text-align:center;
}
</style>

<h1 class="center">这个标题不受影响</h1>
<p class="center">这个段落居中对齐。</p> 
```

-----------------------------



### CSS 创建



#### 如何插入样式表

插入样式表的方法有三种:

- 外部样式表(External style sheet)
- 内部样式表(Internal style sheet)
- 内联样式(Inline style)



#### 外部样式表

​	当样式需要应用于很多页面时，外部样式表将是理想的选择。在使用外部样式表的情况下，你可以通过改变一个文件来改变整个站点的外观。每个页面使用 `<link>` 标签链接到样式表。 `<link> `标签在（文档的）头部：

```html
<head>
<link rel="stylesheet" type="text/css" href="mystyle.css">
</head>
```

浏览器会从文件 mystyle.css 中读到样式声明，并根据它来格式文档。**外部样式表可以在任何文本编辑器中进行编辑。文件不能包含任何的 html 标签。样式表应该以 .css 扩展名进行保存**。下面是一个样式表文件的例子：

```css
hr {color:sienna;}
p {margin-left:20px;}
body {background-image:url("/images/back40.gif");}
```

> ![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111238422.gif) 不要在属性值与单位之间留有空格（如："margin-left: 20 px" ），正确的写法是 "margin-left: 20px" 。



#### 内部样式表

当单个文档需要特殊的样式时，就应该使用内部样式表。你可以使用` <style>` 标签在文档头部定义内部样式表，就像这样:
```html
<head> 
    <style>
        hr {color:sienna;} 
        p {margin-left:20px;} 
        body {background-image:url("images/back40.gif");} 
    </style> 
</head>
```



#### 内联样式

由于要将表现和内容混杂在一起，内联样式会损失掉样式表的许多优势。请慎用这种方法，例如当样式仅需要在一个元素上应用一次时。

要使用内联样式，你需要在相关的标签内使用样式（style）属性。Style 属性可以包含任何 CSS 属性。本例展示如何改变段落的颜色和左外边距：

```html
<p style="color:sienna;margin-left:20px">这是一个段落。</p>
```



#### 多重样式

如果某些属性在不同的样式表中被同样的选择器定义，那么属性值将从更具体的样式表中被继承过来。 

某个属性在不同的样式中被同样的选择器定义时，**采取兼容覆盖（对只有低级样式才有的属性，会保存下来）原则，即高优先级的样式会覆盖低级样式，在head中引用顺序也会影响覆盖**

例如，外部样式表拥有针对 h3 选择器的三个属性：

```CSS
h3 {    color:red;    text-align:left;    font-size:8pt; }
```

而内部样式表拥有针对 h3 选择器的两个属性：

```html
h3 {    text-align:right;    font-size:20pt; }
```

假如拥有内部样式表的这个页面同时与外部样式表链接，那么 h3 得到的样式是：

```html
color:red; text-align:right; font-size:20pt;
```

即颜色属性将被继承于外部样式表，而文字排列（text-alignment）和字体尺寸（font-size）会被内部样式表中的规则取代。



#### 多重样式优先级

样式表允许以多种方式规定样式信息。样式可以规定在单个的 HTML 元素中，在 HTML 页的头元素中，或在一个外部的 CSS 文件中。甚至可以在同一个 HTML 文档内部引用多个外部样式表。

一般情况下，优先级如下：

**（内联样式）Inline style > （内部样式）Internal style sheet >（外部样式）External style sheet > 浏览器默认样式**

----------------------------------



### CSS 背景

**CSS 背景属性用于定义HTML元素的背景。**

CSS 属性定义背景效果:

- background-color
- background-image
- background-repeat
- background-attachment
- background-position



#### 背景颜色

**background-color** 属性定义了元素的**背景颜色.**

```css
body {background-color:#b0c4de;}
```

CSS中，颜色值通常以以下方式定义:

- 十六进制 - 如："#ff0000"
- RGB - 如："rgb(255,0,0)"
- 颜色名称 - 如："red"





#### 背景图像

**background-image 属性描述了元素的背景图像**。默认情况下，背景图像进行平铺重复显示，以覆盖整个元素实体.

页面背景图片设置实例:

```CSS
body {background-image:url('paper.gif');}
```



##### 水平或垂直平铺

默认情况下 background-image 属性会在页面的水平或者垂直方向平铺。一些图像如果在水平方向与垂直方向平铺，这样看起来很不协调，如下所示: 

```css
body
{
background-image:url('gradient2.png');
background-repeat:repeat-x;  /* 设置只在水平方向平铺  */
}
```



##### 设置定位与不平铺

让背景图像不影响文本的排版。**如果你不想让图像平铺，你可以使用 background-repeat 属性:**

```css
body
{
background-image:url('img_tree.png');
background-repeat:no-repeat;
}
```

以上实例中，背景图像与文本显示在同一个位置，为了让页面排版更加合理，不影响文本的阅读，我们可以改变图像的位置。

可以利用 **background-position 属性改变图像在背景中的位置:**

```css
body
{
background-image:url('img_tree.png');
background-repeat:no-repeat;
background-position:right top;
background-attachment:fixed;  /* 固定图像，让其不会随页面下滑而消失 */
}
```



#### 简写属性

在以上实例中我们可以看到页面的背景颜色通过了很多的属性来控制。为了简化这些属性的代码，我们可以将这些属性合并在同一个属性中。背景颜色的简写属性为 "background":

```css
body {background:#ffffff url('img_tree.png') no-repeat right top;}
```

当使用简写属性时，属性值的顺序为：:

- background-color：背景颜色
- background-image：背景图片
- background-repeat：重复铺设
- background-attachment：是否固定
- background-position：放置位置

**以上属性无需全部使用，你可以按照页面的实际需要使用.**



#### CSS 背景属性

| Property                                                     | 描述                                         |
| :----------------------------------------------------------- | :------------------------------------------- |
| [background](https://www.runoob.com/cssref/css3-pr-background.html) | 简写属性，作用是将背景属性设置在一个声明中。 |
| [background-attachment](https://www.runoob.com/cssref/pr-background-attachment.html) | 背景图像是否固定或者随着页面的其余部分滚动。 |
| [background-color](https://www.runoob.com/cssref/pr-background-color.html) | 设置元素的背景颜色。                         |
| [background-image](https://www.runoob.com/cssref/pr-background-image.html) | 把图像设置为背景。                           |
| [background-position](https://www.runoob.com/cssref/pr-background-position.html) | 设置背景图像的起始位置。                     |
| [background-repeat](https://www.runoob.com/cssref/pr-background-repeat.html) | 设置背景图像是否及如何重复。                 |

---------------------------------



### CSS 文本格式

**具体的文本属性：[CSS Text(文本) | 菜鸟教程 (runoob.com)](https://www.runoob.com/css/css-text.html)**



#### 文本颜色

颜色属性被用来设置文字的颜色。

颜色是通过CSS最经常的指定：

- 十六进制值 - 如: **＃FF0000**
- 一个RGB值 - 如: **RGB(255,0,0)**
- 颜色的名称 - 如: **red**

```css
body {color:red;}
h1 {color:#00ff00;}
h2 {color:rgb(255,0,0);}
```



#### 文本的对齐方式

文本排列属性是用来设置文本的水平对齐方式。

文本可居中或对齐到左或右,两端对齐.

当text-align设置为"justify"，每一行被展开为宽度相等，左，右外边距是对齐（如杂志和报纸）。

```css
h1 {text-align:center;}
p.date {text-align:right;}
p.main {text-align:justify;}
```



#### 文本修饰

text-decoration 属性用来设置或删除文本的装饰。

从设计的角度看 text-decoration属性主要是用来删除链接的下划线：

```css
a {text-decoration:none;}
```



#### 文本转换

文本转换属性是用来指定在一个文本中的大写和小写字母。

可用于**所有字句变成大写或小写字母，或每个单词的首字母大写。**

```css
p.uppercase {text-transform:uppercase;}
p.lowercase {text-transform:lowercase;}
p.capitalize {text-transform:capitalize;}  
```



#### 文本缩进

文本缩进属性是用来指定文本的第一行的缩进。

```css
p {text-indent:50px;}
```

--------------------------------------------------------------------------------------



### CSS 字体

**CSS字体属性定义字体，加粗，大小，文字样式。** 	



#### CSS字型

在CSS中，有两种类型的字体系列名称：

- **通用字体系列** - 拥有相似外观的字体系统组合（如 "Serif" 或 "Monospace"）
- **特定字体系列** - 一个特定的字体系列（如 "Times" 或 "Courier"）

| Generic family | 字体系列                   | 说明                                        |
| :------------- | :------------------------- | :------------------------------------------ |
| Serif          | Times New Roman Georgia    | Serif字体中字符在行的末端拥有额外的装饰     |
| Sans-serif     | Arial Verdana              | "Sans"是指无 - 这些字体在末端没有额外的装饰 |
| Monospace      | Courier New Lucida Console | 所有的等宽字符具有相同的宽                  |



#### 字体系列

font-family 属性设置文本的字体系列。

font-family 属性应该设置几个字体名称作为一种"后备"机制，如果浏览器不支持第一种字体，他将尝试下一种字体。

**注意**: 如果字体系列的名称超过一个字，它必须用引号，如Font Family："宋体"。多个字体系列是用一个逗号分隔指明：

```css
p{font-family:"Times New Roman", Times, serif;}
```



#### 字体样式

**front-style**主要是用于指定斜体文字的字体样式属性。

这个属性有三个值：

- 正常 - 正常显示文本
- 斜体 - 以斜体字显示的文字
- 倾斜的文字 - 文字向一边倾斜（和斜体非常类似，但不太支持）

```css
p.normal {font-style:normal;}
p.italic {font-style:italic;}
p.oblique {font-style:oblique;}
```



#### 字体大小

**font-size 属性设置文本的大小。**

能否管理文字的大小，在网页设计中是非常重要的。但是，你不能通过调整字体大小使段落看上去像标题，或者使标题看上去像段落。

请务必使用正确的HTML标签，就`<h1> - <h6>`表示标题和`<p>`表示段落：**字体大小的值可以是绝对或相对的大小。**

绝对大小：

- 设置一个指定大小的文本
- 不允许用户在所有浏览器中改变文本大小
- 确定了输出的物理尺寸时绝对大小很有用

相对大小：

- 相对于周围的元素来设置大小
- 允许用户在浏览器中改变文字大小

![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111343439.gif) 如果你不指定一个字体的大小，默认大小和普通文本段落一样，是16像素（16px=1em）。

```css
h1 {font-size:40px;}
```

上面的例子可以在 Internet Explorer 9, Firefox, Chrome, Opera, 和 Safari 中通过缩放浏览器调整文本大小。虽然可以通过浏览器的缩放工具调整文本大小，但是，这种调整是整个页面，而不仅仅是文本



#### 用em来设置字体大小

为了避免Internet Explorer 中无法调整文本的问题，许多开发者使用 em 单位代替像素。

em的尺寸单位由W3C建议。1em和当前字体大小相等。在浏览器中默认的文字大小是16px。因此，1em的默认大小是16px。可以通过下面这个公式将像素转换为em：**px/16=em**



#### 使用百分比和EM组合

在所有浏览器的解决方案中，设置 `<body>`元素的默认字体大小的是百分比：

```css
body {font-size:100%;}
```



#### 设置字体粗细

```css
p.normal {font-weight:normal;}
p.light {font-weight:lighter;}
p.thick {font-weight:bold;}
p.thicker {font-weight:900;}
```

---------------------------------------------



### CSS 链接



#### 链接样式

链接的样式，可以用任何CSS属性（如颜色，字体，背景等）。特别的链接，可以有不同的样式，这取决于他们是什么状态。

**这四个链接状态是：**

- a:link - 正常，未访问过的链接
- a:visited - 用户已访问过的链接
- a:hover - 当用户鼠标放在链接上时
- a:active - 链接被点击的那一刻

```css
/* 本来就有的，用：访问
a:link {color:#000000;}      /* 未访问链接*/
a:visited {color:#00FF00;}  /* 已访问链接 */
a:hover {color:#FF00FF;}  /* 鼠标移动到链接上 */
a:active {color:#0000FF;}  /* 鼠标点击时 */
```

当设置为若干链路状态的样式，也有一些顺序规则：

- a:hover 必须跟在 a:link 和 a:visited后面
- a:active 必须跟在 a:hover后面



#### 文本修饰

text-decoration 属性主要用于删除链接中的下划线：



----------------------




### CSS 列表


CSS 列表属性作用如下：

- 设置不同的列表项标记为有序列表
- 设置不同的列表项标记为无序列表
- 设置列表项标记为图像



#### 列表

在 HTML中，有两种类型的列表：

- 无序列表 **ul** - 列表项标记用特殊图形（如小黑点、小方框等）
- 有序列表 **ol** - 列表项的标记有数字或字母

使用 CSS，可以列出进一步的样式，并可用图像作列表项标记。



#### 不同的列表项标记

**list-style-type属性**指定列表项标记的类型是

```css
ul.a {list-style-type: circle;}
ul.b {list-style-type: square;}
 
ol.c {list-style-type: upper-roman;}
ol.d {list-style-type: lower-alpha;}
```



#### 作为列表项标记的图像

要指定列表项标记的图像，使用列表样式图像属性：

```css
ul
{
    list-style-image: url('sqpurple.gif');
}
```



#### 移除默认设置

list-style-type:none 属性可以用于移除小标记。默认情况下列表 <ul> 或 <ol> 还设置了内边距和外边距，可使用 `margin:0` 和 `padding:0` 来移除:

```css
ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
}
```

------------------------



### CSS 表格

------

使用 CSS 可以使 HTML 表格更美观。



#### 表格边框

**指定CSS表格边框，使用border属性**。下面的例子指定了一个表格的Th和TD元素的黑色边框：

```css
table, th, td
{
    border: 1px solid black;
}
```



#### 折叠边框

**border-collapse 属性设置表格的边框是否被折叠成一个单一的边框或隔开**：

```css
table
{
    border-collapse:collapse;
}
table,th, td
{
    border: 1px solid black;
}
```



#### 表格宽度和高度

**Width和height属性定义表格的宽度和高度。**下面的例子是设置100％的宽度，50像素的th元素的高度的表格：

```css
table 
{
    width:100%;
}
th
{
    height:50px;
}
```



#### 表格文字对齐

表格中的文本对齐和垂直对齐属性。**text-align属性设置水平对齐方式，向左，右，或中心：**

```css
td
{
    text-align:right;
}
```



#### 表格填充

如需控制边框和表格内容之间的间距，应使用td和th元素的填充属性：

```css
td
{
    padding:15px;
}
```



#### 表格颜色

下面的例子指定边框的颜色，和th元素的文本和背景颜色：

```css
table, td, th
{
    border:1px solid green;
}
th
{
    background-color:green;
    color:white;
}
```

------------------------------------



### CSS 盒子模型(Box Model)

所有HTML元素可以看作盒子，在CSS中，"box model"这一术语是用来设计和布局时使用。**CSS盒模型本质上是一个盒子，封装周围的HTML元素，它包括：边距，边框，填充，和实际内容。**盒模型允许我们在其它元素和周围元素边框之间的空间放置元素。

下面的图片说明了盒子模型(Box Model)：

![CSS box-model](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111438683.gif)

不同部分的说明：

- **Margin(外边距)** - 清除边框外的区域，外边距是透明的。
- **Border(边框)** - 围绕在内边距和内容外的边框。
- **Padding(内边距)** - 清除内容周围的区域，内边距是透明的。
- **Content(内容)** - 盒子的内容，显示文本和图像。

为了正确设置元素在所有浏览器中的宽度和高度，你需要知道的盒模型是如何工作的。



#### 元素的宽度和高度

![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111439415.gif)**重要:** 当您指定一个 CSS 元素的宽度和高度属性时，你只是设置内容区域的宽度和高度。要知道，完整大小的元素，你还必须添加内边距，边框和外边距。

**最终元素的总宽度计算公式是这样的：**

总元素的宽度=宽度+左填充+右填充+左边框+右边框+左边距+右边距

元素的总高度最终计算公式是这样的：

总元素的高度=高度+顶部填充+底部填充+上边框+下边框+上边距+下边距

------------------------



### CSS 边框



#### CSS 边框属性

CSS边框属性允许你指定一个元素边框的样式和颜色。



#### 边框样式

边框样式属性指定要显示什么样的边界。

![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111444319.gif) **border-style**属性用来定义边框的样式



#### 边框宽度

您可以通过 **border-width 属性为边框指定宽度。**

为边框指定宽度有两种方法：可以指定长度值，比如 2px 或 0.1em(单位为 px, pt, cm, em 等)，或者使用 3 个关键字之一，它们分别是 thick 、medium（默认值） 和 thin。

**注意：**CSS 没有定义 3 个关键字的具体宽度，所以一个用户可能把 thick 、medium 和 thin 分别设置为等于 5px、3px 和 2px，而另一个用户则分别设置为 3px、2px 和 1px。



#### 边框颜色

**border-color属性用于设置边框的颜色。**可以设置的颜色：

- name - 指定颜色的名称，如 "red"
- RGB - 指定 RGB 值, 如 "rgb(255,0,0)"
- Hex - 指定16进制值, 如 "#ff0000"

您还可以设置边框的颜色为"transparent"。



#### 单独设置各边

在CSS中，可以指定不同的侧面不同的边框：

```css
p
{
    border-top-style:dotted;
    border-right-style:solid;
    border-bottom-style:dotted;
    border-left-style:solid;
}
```

----------------





### CSS 轮廓（outline）

轮廓（outline）是绘制于元素周围的一条线，位于边框边缘的外围，可起到突出元素的作用。轮廓（outline）属性指定元素轮廓的样式、颜色和宽度。

![Outline](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111535240.gif)

-----------------------





### CSS margin(外边距)

margin 清除周围的（外边框）元素区域。**margin 没有背景颜色，是完全透明的。**

margin 可以单独改变元素的上，下，左，右边距，也可以一次改变所有的属性。

![img](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111535693.png)



#### 可能的值

| 值       | 说明                                        |
| :------- | :------------------------------------------ |
| auto     | 设置浏览器边距。 这样做的结果会依赖于浏览器 |
| *length* | 定义一个固定的margin（使用像素，pt，em等）  |
| *%*      | 定义一个使用百分比的边距                    |

![Remark](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111535244.gif) Margin可以使用负值，重叠的内容。



#### Margin - 单边外边距属性

在CSS中，它可以指定不同的侧面不同的边距：

```css
margin-top:100px;
margin-bottom:100px;
margin-right:50px;
margin-left:50px;
```

-------------------------



### CSS padding（填充）

CSS padding（填充）是一个简写属性，定义元素边框与元素内容之间的空间，即上下左右的内边距。



#### padding（填充）

当元素的 padding（填充）内边距被清除时，所释放的区域将会受到元素背景颜色的填充。

单独使用 padding 属性可以改变上下左右的填充。

![img](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301111625090.png)



#### 填充- 单边内边距属性

在CSS中，它可以指定不同的侧面不同的填充：

```css
padding-top:25px;
padding-bottom:25px;
padding-right:50px;
padding-left:50px;
```



#### 填充 - 简写属性

为了缩短代码，它可以在一个属性中指定的所有填充属性。

这就是所谓的简写属性。所有的填充属性的简写属性是 **padding** :

```css
padding:25px 50px;
```



------------------------------------



### CSS 分组 和 嵌套 选择器



#### 分组选择器

在样式表中有很多具有相同样式的元素。**为了尽量减少代码，你可以使用分组选择器。**

**每个选择器用逗号分隔。**在下面的例子中，我们对以上代码使用分组选择器

```css
h1,h2,p
{
    color:green;
}
```



#### 嵌套选择器

它可能适用于选择器内部的选择器的样式。

在下面的例子设置了四个样式：

- **p{ }**: 为所有 **p** 元素指定一个样式。
- **.marked{ }**: 为所有 **class="marked"** 的元素指定一个样式。
- **.marked p{ }**: 为所有 **class="marked"** 元素内的 **p** 元素指定一个样式。
- **p.marked{ }**: 为所有 **class="marked"** 的 **p** 元素指定一个样式。

```css
p
{
    color:blue;
    text-align:center;
}
.marked
{
    background-color:red;
}
.marked p
{
    color:white;
}
p.marked{
    text-decoration:underline;
}

<p>这个段落是蓝色文本，居中对齐。</p>   /* p元素 */
<div class="marked">   /* .marked元素 */
<p>这个段落不是蓝色文本。</p>   /* .marked p 元素 */
</div>
<p>所有 class="marked"元素内的 p 元素指定一个样式，但有不同的文本颜色。</p>  
	
<p class="marked">带下划线的 p 段落。</p>  /* p.marked 元素 */
```

----------------------



### CSS 尺寸 (Dimension)

CSS 尺寸 (Dimension) 属性允许你控制元素的高度和宽度。同样，它允许你增加行间距。

| 属性                                                         | 描述                 |
| :----------------------------------------------------------- | :------------------- |
| [height](https://www.runoob.com/cssref/pr-dim-height.html)   | 设置元素的高度。     |
| [line-height](https://www.runoob.com/cssref/pr-dim-line-height.html) | 设置行高。           |
| [max-height](https://www.runoob.com/cssref/pr-dim-max-height.html) | 设置元素的最大高度。 |
| [max-width](https://www.runoob.com/cssref/pr-dim-max-width.html) | 设置元素的最大宽度。 |
| [min-height](https://www.runoob.com/cssref/pr-dim-min-height.html) | 设置元素的最小高度。 |
| [min-width](https://www.runoob.com/cssref/pr-dim-min-width.html) | 设置元素的最小宽度。 |
| [width](https://www.runoob.com/cssref/pr-dim-width.html)     | 设置元素的宽度。     |

---------------------





### CSS Display(显示) 与 Visibility（可见性）

 **display属性设置一个元素应如何显示，visibility属性指定一个元素应可见还是隐藏。**



#### 隐藏元素 - display:none或visibility:hidden

隐藏一个元素可以通过把display属性设置为"none"，或把visibility属性设置为"hidden"。但是请注意，这两种方法会产生不同的结果。

**visibility:hidden可以隐藏某个元素，但隐藏的元素仍需占用与未隐藏之前一样的空间（实际存在，只是不可见）。**也就是说，该元素虽然被隐藏了，但仍然会影响布局。

```css
h1.hidden {visibility:hidden;}
```

**display:none可以隐藏某个元素，且隐藏的元素不会占用任何空间（相当于被注释掉了）**。也就是说，该元素不但被隐藏了，而且该元素原本占用的空间也会从页面布局中消失。

```css
h1.hidden {display:none;}
```



#### CSS Display - 块和内联元素

块元素是一个元素，占用了全部宽度，在前后都是换行符。

块元素的例子：

- `<h1>`
- `<p>`
- `<div>`

内联元素只需要必要的宽度，不强制换行。

内联元素的例子：

- `<span>`
- `<a>`



#### 如何改变一个元素显示

可以更改内联元素和块元素，反之亦然，可以使页面看起来是以一种特定的方式组合，并仍然遵循web标准。

下面的示例把列表项显示为内联元素：

```css
li {display:inline;}
```

下面的示例把span元素作为块元素：

```css
span {display:block;}
```

**注意：**变更元素的显示类型看该元素是如何显示，它是什么样的元素。例如：一个内联元素设置为display:block是不允许有它内部的嵌套块元素。

--------------



### CSS Position(定位)

**position 属性指定了元素的定位类型。**

position 属性的五个值：

- [static](https://www.runoob.com/css/css-positioning.html#position-static)
- [relative](https://www.runoob.com/css/css-positioning.html#position-relative)
- [fixed](https://www.runoob.com/css/css-positioning.html#position-fixed)
- [absolute](https://www.runoob.com/css/css-positioning.html#position-absolute)
- [sticky](https://www.runoob.com/css/css-positioning.html#position-sticky)

元素可以使用的顶部，底部，左侧和右侧属性定位。然而，这些属性无法工作，除非是先设定position属性。他们也有不同的工作方式，这取决于定位方法。



#### static 定位

HTML 元素的默认值，即没有定位，遵循正常的文档流对象。

静态定位的元素不会受到 top, bottom, left, right影响。

```css
div.static {
    position: static;
    border: 3px solid #73AD21;
}
```



#### fixed 定位

元素的位置相对于浏览器窗口是固定位置。

即使窗口是滚动的它也不会移动：

```css
p.pos_fixed
{
    position:fixed;
    top:30px;
    right:5px;
}
```

**注意：** Fixed 定位在 IE7 和 IE8 下需要描述 !DOCTYPE 才能支持。**Fixed定位使元素的位置与文档流无关，因此不占据空间。Fixed定位的元素和其他元素重叠。**



#### relative 定位

**相对定位元素的定位是相对其正常位置。**

```css
h2.pos_left
{
    position:relative;
    left:-20px;
}
h2.pos_right
{
    position:relative;
    left:20px;
}
```

移动相对定位元素，但它原本所占的空间不会改变（**即显示在移动后的位置，但实际占用的是原为的空间）。**

```css
h2.pos_top
{
    position:relative;
    top:-50px;
}
```

相对定位元素经常被用来作为绝对定位元素的容器块。



#### absolute 定位

绝对定位的元素的位置相对于最近的已定位父元素，如果元素没有已定位的父元素，那么它的位置相对于`<html>`:

```css
h2
{
    position:absolute;
    left:100px;
    top:150px;
}
```





#### sticky 定位

sticky 英文字面意思是粘，粘贴，所以可以把它称之为粘性定位。

**position: sticky;** **基于用户的滚动位置来定位。**

粘性定位的元素是依赖于用户的滚动，在 **position:relative** 与 **position:fixed** 定位之间切换。

它的行为就像 **position:relative;** 而当页面滚动超出目标区域时，它的表现就像 **position:fixed;**，它会固定在目标位置。

**元素定位表现为在跨越特定阈值前为相对定位，之后为固定定位。**

这个特定阈值指的是 top, right, bottom 或 left 之一，换言之，指定 top, right, bottom 或 left 四个阈值其中之一，才可使粘性定位生效。否则其行为与相对定位相同。

**注意:** Internet Explorer, Edge 15 及更早 IE 版本不支持 sticky 定位。 Safari 需要使用 -webkit- prefix (查看以下实例)

```css
div.sticky {
    position: -webkit-sticky; /* Safari */
    position: sticky;
    top: 0;
    background-color: green;
    border: 2px solid #4CAF50;
}
```



#### 重叠的元素

元素的定位与文档流无关，所以它们可以覆盖页面上的其它元素

z-index属性指定了一个元素的堆叠顺序（哪个元素应该放在前面，或后面）

一个元素可以有正数或负数的堆叠顺序：

```css
img
{
    position:absolute;
    left:0px;
    top:0px;
    z-index:-1;
}
```

具有更高堆叠顺序的元素总是在较低的堆叠顺序元素的前面。

**注意：** 如果两个定位元素重叠，没有指定z - index，最后定位在HTML代码中的元素将被显示在最前面。

-------------------



### CSS 布局 - Overflow

CSS overflow 属性用于**控制内容溢出元素框时显示的方式**。



**overflow属性有以下值：**

| 值      | 描述                                                     |
| :------ | :------------------------------------------------------- |
| visible | 默认值。内容不会被修剪，会呈现在元素框之外。             |
| hidden  | 内容会被修剪，并且其余内容是不可见的。                   |
| scroll  | 内容会被修剪，但是浏览器会显示滚动条以便查看其余的内容。 |
| auto    | 如果内容被修剪，则浏览器会显示滚动条以便查看其余的内容。 |
| inherit | 规定应该从父元素继承 overflow 属性的值。                 |



-------------------------



### CSS Float(浮动)

CSS 的 Float（浮动），会使元素向左或向右移动，其周围的元素也会重新排列。

Float（浮动），往往是用于图像，但它在布局时一样非常有用。



#### 元素怎样浮动

**元素的水平方向浮动，意味着元素只能左右移动**而不能上下移动。

**一个浮动元素会尽量向左或向右移动，直到它的外边缘碰到包含框或另一个浮动框的边框为止。**

浮动元素**之后的元素将围绕它。**

浮动元素**之前的元素将不会受到影响**。

如果图像是右浮动，下面的文本流将环绕在它左边：

```css
img
{
    float:right;
}
```



#### 彼此相邻的浮动元素

如果你把几个浮动的元素放到一起，如果有空间的话，它们将彼此相邻。如果没有，会自动被挤到下一行去。

在这里，我们对图片廊使用 float 属性：

```css
.thumbnail 
{
    float:left;
    width:110px;
    height:90px;
    margin:5px;
}
```



#### 清除浮动 - 使用 clear

元素浮动之后，周围的元素会重新排列，为了避免这种情况，使用 clear 属性。

clear 属性指定元素两侧不能出现浮动元素。

使用 clear 属性往文本中添加图片廊：

```css
.text_line
{
    clear:both;
}
```

----------------





### CSS 布局 - 水平 & 垂直对齐



#### 元素居中对齐

要水平居中对齐一个元素(如 `<div>`), 可以使用 **margin: auto;**。

设置到元素的宽度将防止它溢出到容器的边缘。

元素通过指定宽度，并将两边的空外边距平均分配：

```css
.center {
    margin: auto;
    width: 50%;
    border: 3px solid green;
    padding: 10px;
}
```



#### 文本居中对齐

如果仅仅是为了文本在元素内居中对齐，可以使用 **text-align: center;**



#### 图片居中对齐

要让图片居中对齐, 可以使用 **margin: auto;** 并将它放到 **块** 元素中:

```css
img {
    display: block;
    margin: auto;
    width: 40%;
}
```



#### 左右对齐 - 使用定位方式

我们可以使用 **position: absolute;** 属性来对齐元素:



#### 垂直居中对齐 - 使用 padding

CSS 中有很多方式可以实现垂直居中对齐。 一个简单的方式就是头部顶部使用 **padding**:

```css
.center {
    padding: 70px 0;
    border: 3px solid green;
}
```

---------------------



### CSS 组合选择符

CSS组合选择符包括各种简单选择符的组合方式。

**在 CSS3 中包含了四种组合方式:**

- 后代选择器(以空格` `分隔)
- 子元素选择器(以大于 `>` 号分隔）
- 相邻兄弟选择器（以加号`+`分隔）
- 普通兄弟选择器（以波浪号 `~`分隔）



#### 后代选择器

后代选择器用于选取某元素的后代元素（包含关系，div包含p）。以下实例选取所有`<p> `元素插入到 `<div>`元素中: 

```css
div p
{
	background-color:yellow;
}
<div>
<p>段落 1。 在 div 中。</p>
<p>段落 2。 在 div 中。</p>
</div>
```



#### 子元素选择器

与后代选择器相比，子元素选择器（Child selectors）只能选择作为某元素直接/一级子元素的元素**（可以把`>`看成是路径符号，其不能跨路径选择）**。以下实例选择了`<div>`元素中所有直接子元素 `<p> `：

```css
div>p
{
	background-color:yellow;
}
/* 差一级，可以生效 */
<div>
<p>I live in Duckburg.</p>
</div>

/* 差大于一级，不能生效 */
<div>
<span><p>I will not be styled.</p></span>
</div>
```



#### 相邻兄弟选择器

**相邻兄弟选择器（Adjacent sibling selector）可选择紧接在另一元素后的元素，且二者有相同父元素。**

如果需要选择紧接在另一个元素后的元素，而且二者有相同的父元素，可以使用相邻兄弟选择器（Adjacent sibling selector）。

以下实例选取了所有位于 `<div>` 元素后的第一个 `<p> `元素:

```css
div+p
{
	background-color:yellow;
}
<div>
<h2>DIV 内部标题</h2>
<p>DIV 内部段落。</p>
</div>

<p>DIV 之后的第一个 P 元素。</p>  /* 可以生效 */
```



#### 后续兄弟选择器

后续兄弟选择器**选取所有指定元素之后的相邻兄弟元素**。

以下实例选取了所有` <div> `元素之后的所有相邻兄弟元素` <p> `: 

```css
div~p
{
	background-color:yellow;
}
	
<p>之前段落，不会添加背景颜色。</p>
<div>
<p>段落 1。 在 div 中。</p>
</div>
/ * div后的所有p的都生效（同级） */
<p>段落 3。不在 div 中。</p>
<p>段落 4。不在 div 中。</p>
```

--------------------------------





### CSS 伪类(Pseudo-classes)

CSS伪类是用来添加一些选择器的特殊效果。

**伪类的语法：**

```css
selector:pseudo-class {property:value;}
```

**CSS类也可以使用伪类：**

```css
selector.class:pseudo-class {property:value;}
```



#### anchor伪类

在支持 CSS 的浏览器中，链接的不同状态都可以以不同的方式显示

```css
a:link {color:#FF0000;} /* 未访问的链接 */
a:visited {color:#00FF00;} /* 已访问的链接 */
a:hover {color:#FF00FF;} /* 鼠标划过链接 */
a:active {color:#0000FF;} /* 已选中的链接 */
```

**注意：** 在CSS定义中，a:hover 必须被置于 a:link 和 a:visited 之后，才是有效的。

**注意：** 在 CSS 定义中，a:active 必须被置于 a:hover 之后，才是有效的。

**注意：**伪类的名称不区分大小写。



#### 伪类和CSS类

伪类可以与 CSS 类配合使用：

```css
a.red:visited {color:#FF0000;}
<a class="red" href="css-syntax.html">CSS 语法</a>
```

如果在上面的例子的链接已被访问，它会显示为红色。



#### CSS :first-child 伪类

您可以使用 :first-child 伪类来选择父元素的第一个子元素。



#### 匹配第一个 `<p>` 元素

在下面的例子中，选择器匹配作**为任何元素的第一个子元素的` <p> `元素**：

```css
p:first-child
{
    color:blue;
}
```



#### 匹配所有`<p>` 元素中的第一个` <i> `元素

在下面的例子中，选择相匹配的所有`<p>`元素的第一个` <i> `元素：

```css
p > i:first-child
{
    color:blue;
}
```



#### CSS - :lang 伪类

**:lang 伪类使你有能力为不同的语言定义特殊的规则**

在下面的例子中，:lang 类为属性值为 no 的q元素定义引号的类型：

```css
/* 为其他标签添加前后缀 */
q:lang(no) {quotes: "~" "~";}
```



#### :first-line 伪元素

"first-line" 伪元素用于向文本的首行设置特殊样式。

在下面的例子中，浏览器会根据 "first-line" 伪元素中的样式对 p 元素的第一行文本进行格式化：

```css
p:first-line 
{
    color:#ff0000;
    font-variant:small-caps;
}
```

**注意：**"first-line" 伪元素只能用于块级元素。



#### :first-letter 伪元素

"first-letter" 伪元素用于向文本的首字母设置特殊样式：

```css
p:first-letter 
{
    color:#ff0000;
    font-size:xx-large;
}
```

**注意：** "first-letter" 伪元素只能用于块级元素。





#### 伪元素和CSS类

伪元素可以结合CSS类：

```css
p.article:first-letter {color:#ff0000;}

<p class="article">文章段落</p>
```

上面的例子会使所有 class 为 article 的段落的首字母变为红色。



#### :before 伪元素

":before" 伪元素可以在元素的内容前面插入新内容。

下面的例子在每个 `<h1>`元素前面插入一幅图片：

```css
h1:before 
{
    content:url(smiley.gif);
}
```



#### :after 伪元素

":after" 伪元素可以在元素的内容之后插入新内容。

下面的例子在每个` <h1>` 元素后面插入一幅图片：

```css
h1:after
{
    content:url(smiley.gif);
}
```



### CSS 导航栏

熟练使用导航栏，对于任何网站都非常重要。

使用CSS你可以转换成好看的导航栏而不是枯燥的HTML菜单。

**详细看：**[CSS 导航栏 | 菜鸟教程 (runoob.com)](https://www.runoob.com/css/css-navbar.html)



#### 导航栏=链接列表

作为标准的 HTML 基础一个导航栏是必须的。在我们的例子中我们将建立一个标准的 HTML 列表导航栏。导航条基本上是一个链接列表，所以使用 `<ul>` 和` <li>`元素非常有意义：

```css
<ul>
  <li><a href="#home">主页</a></li>
  <li><a href="#news">新闻</a></li>
  <li><a href="#contact">联系</a></li>
  <li><a href="#about">关于</a></li>
</ul>
```

现在，让我们从列表中删除边距和填充：

```css
ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
}
```

例子解析：

- list-style-type:none - 移除列表前小标志。一个导航栏并不需要列表标记
- 移除浏览器的默认设置将边距和填充设置为0

上面的例子中的代码是垂直和水平导航栏使用的标准代码。



#### 垂直导航栏

上面的代码，我们只需要 `<a>`元素的样式，建立一个垂直的导航栏：

```css
a
{
    display:block;
    width:60px;
}
```

示例说明：

- display:block - 显示块元素的链接，让整体变为可点击链接区域（不只是文本），它允许我们指定宽度
- width:60px - 块元素默认情况下是最大宽度。我们要指定一个60像素的宽度

---------------------



### CSS 属性 选择器



#### 属性选择器

下面的例子是把包含标题（title）的所有元素变为蓝色：

```css
[title]
{
    color:blue;
}
```



#### 属性和值选择器

下面的实例改变了标题title='runoob'元素的边框样式:

```css
[title=runoob]
{
    border:5px solid green;
}

<a title="runoob" href="http://www.runoob.com/">runoob</a>
<hr>
```



#### 属性和值的选择器 - 多值

下面是包含指定值的title属性的元素样式的例子，使用（~）分隔属性和值:

```css
[title~=hello] { color:blue; }
/* 包含Hello属性值的都会应用样式 */
```

下面是包含指定值的lang属性的元素样式的例子，使用（|）分隔属性和值:

```css
[lang|=en] { color:blue; }
/* 包含en值 */
```



#### 表单样式

属性选择器样式无需使用class或id的形式:

```css
input[type="text"]
{
    width:150px;
    display:block;
    margin-bottom:10px;
    background-color:yellow;
}
input[type="button"]
{
    width:120px;
    margin-left:35px;
    display:block;
}
```

-------------------------------



### CSS 表单



#### 输入框(input) 样式

使用 width 属性来设置输入框的宽度

```css
input {
  width: 100%;
}
```

以上实例中设置了所有` <input>` 元素的宽度为 100%，如果你只想设置指定类型的输入框可以使用以下属性选择器：

- `input[type=text]` - 选取文本输入框
- `input[type=password]` - 选择密码的输入框
- `input[type=number]` - 选择数字的输入框
- ...



#### 输入框填充

使用 **padding** 属性可以在输入框中添加内边距。

```css
input[type=text] {
  width: 100%;
  padding: 12px 20px;
  margin: 8px 0;
  box-sizing: border-box;
}
```

注意我们设置了 `box-sizing` 属性为 `border-box`。这样可以确保浏览器呈现出带有指定宽度和高度的输入框是把边框和内边距一起计算进去的。



#### 输入框(input) 边框

使用 `border` 属性可以修改 input 边框的大小或颜色，使用 `border-radius` 属性可以给 input 添加圆角：

```css
input[type=text] {
  border: 2px solid red;
  border-radius: 4px;
}
```



#### 输入框(input) 图标

如果你想在输入框中添加图标，可以使用 `background-image` 属性和用于定位的`background-position` 属性。注意设置图标的左边距，让图标有一定的空间：

```css
input[type=text] {
  background-color: white;
  background-image: url('searchicon.png');
  background-position: 10px 10px; 
  background-repeat: no-repeat;
  padding-left: 40px;
}
```

-------------------------------





### CSS 计数器

CSS 计数器通过一个变量来设置，根据规则递增变量。



#### 使用计数器自动编号

CSS 计数器根据规则来递增变量。

CSS 计数器使用到以下几个属性：

- `counter-reset` - 创建或者重置计数器
- `counter-increment` - 递增变量
- `content` - 插入生成的内容
- `counter()` 或 `counters()` 函数 - 将计数器的值添加到元素

要使用 CSS 计数器，得先用 counter-reset 创建：

以下实例在页面创建一个计数器 (在 body 选择器中)，每个 <h2> 元素的计数值都会递增，并在每个 <h2> 元素前添加 "Section <*计数值*>:"

```css
body {
  /* 创建一个名为session的计数器 */
  counter-reset: section;
}
 
h2::before {
  counter-increment: section;
  content: "Section " counter(section) ": ";
}
```



#### 嵌套计数器

以下实例在页面创建一个计数器，在每一个` <h1> `元素前添加计数值 "Section <*主标题计数值*>.", 嵌套的计数值则放在` <h2> `元素的前面，内容为 "<*主标题计数值*>.<*副标题计数值*>":

```css
ol {
  counter-reset: section;
  list-style-type: none;
}
 
li::before {
  counter-increment: section;
  content: counters(section,".") " ";
}
```





### CSS 网页布局

**具体**：[CSS 网页布局 | 菜鸟教程 (runoob.com)](https://www.runoob.com/css/css-website-layout.html)



#### 网页布局

网页布局有很多种方式，一般分为以下几个部分：**头部区域、菜单导航区域、内容区域、底部区域**。

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301112003545.jpeg" alt="img" style="zoom: 80%;" />

--------------------



### CSS !important 规则

CSS 中的 !important 规则用于增加样式的权重。

**!important** 与优先级无关，但它与最终的结果直接相关，使用一个 !important 规则时，此声明将覆盖任何其他声明。



#### 重要说明

使用 !important 是一个坏习惯，应该尽量避免，因为这破坏了样式表中的固有的级联规则 使得调试找 bug 变得更加困难了。

当两条相互冲突的带有 !important 规则的声明被应用到相同的元素上时，拥有更大优先级的声明将会被采用。

以下实例我们在查看 CSS 源码时就不是很清楚哪种颜色最重要：



#### 何时使用 !important

如果要在你的网站上设定一个全站样式的 CSS 样式可以使用 !important。

比如我们要让网站上所有按钮的样式都一样：

```css
.button {
  background-color: #8c8c8c !important;
  color: white !important;
  padding: 5px !important;
  border: 1px solid black !important;
}
 
#myDiv a {
  color: red;
  background-color: yellow;
}
```





