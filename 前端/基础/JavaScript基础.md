# 





## 概述



### JavaScript 简介

JavaScript 是互联网上最流行的脚本语言，这门语言可用于 HTML 和 web，更可广泛用于服务器、PC、笔记本电脑、平板电脑和智能手机等设备。



#### JavaScript 是脚本语言

- JavaScript 是一种轻量级的编程语言。

- JavaScript 是可插入 HTML 页面的编程代码。

- JavaScript 插入 HTML 页面后，可由所有的现代浏览器执行。



### 运行模式

JavaScript是一种属于网络的高级脚本语言,已经被广泛用于Web应用开发,常用来为网页添加各式各样的动态功能,为用户提供更流畅美观的浏览效果。通常JavaScript脚本是通过嵌入在HTML中来实现自身的功能的。

1. 是一种解释性脚本语言（代码不进行[预编译](https://baike.baidu.com/item/预编译?fromModule=lemma_inlink)）。
2. 主要用来向[HTML](https://baike.baidu.com/item/HTML?fromModule=lemma_inlink)（[标准通用标记语言](https://baike.baidu.com/item/标准通用标记语言?fromModule=lemma_inlink)下的一个应用）页面添加交互行为。
3. 可以直接嵌入HTML页面，但写成单独的[js](https://baike.baidu.com/item/js/10687961?fromModule=lemma_inlink)文件有利于结构和行为的[分离](https://baike.baidu.com/item/分离?fromModule=lemma_inlink)。
4. 跨平台特性，在绝大多数浏览器的支持下，可以在多种平台下运行（如[Windows](https://baike.baidu.com/item/Windows?fromModule=lemma_inlink)、[Linux](https://baike.baidu.com/item/Linux?fromModule=lemma_inlink)、[Mac](https://baike.baidu.com/item/Mac/173?fromModule=lemma_inlink)、[Android](https://baike.baidu.com/item/Android/60243?fromModule=lemma_inlink)、[iOS](https://baike.baidu.com/item/iOS/45705?fromModule=lemma_inlink)等）。
5. JavaScript脚本语言同其他语言一样，有它自身的基本数据类型，表达式和[算术运算符](https://baike.baidu.com/item/算术运算符/9324947?fromModule=lemma_inlink)及程序的基本程序框架。JavaScript提供了四种基本的数据类型和两种特殊数据类型用来处理数据和文字。而变量提供存放信息的地方，表达式则可以完成较复杂的信息处理。



### 语言特点

**JavaScript脚本语言具有以下特点:**

（1）[脚本语言](https://baike.baidu.com/item/脚本语言?fromModule=lemma_inlink)。JavaScript是一种解释型的脚本语言，C、[C++](https://baike.baidu.com/item/C%2B%2B?fromModule=lemma_inlink)等语言先[编译](https://baike.baidu.com/item/编译?fromModule=lemma_inlink)后执行，而JavaScript是在程序的运行过程中逐行进行解释。

（2）基于对象。JavaScript是一种基于对象的脚本语言，它不仅可以创建对象，也能使用现有的对象。

（3）简单。JavaScript语言中采用的是弱类型的变量类型，对使用的数据类型未做出严格的要求，是基于Java基本语句和控制的脚本语言，其设计简单紧凑。

（4）动态性。JavaScript是一种采用事件驱动的脚本语言，它不需要经过Web服务器就可以对用户的输入做出响应。在访问一个网页时，鼠标在网页中进行鼠标点击或上下移、窗口移动等操作JavaScript都可直接对这些事件给出相应的响应。

（5）跨平台性。JavaScript脚本语言不依赖于操作系统，仅需要浏览器的支持。因此一个JavaScript脚本在编写后可以带到任意机器上使用，前提是机器上的浏览器支 持JavaScript脚本语言，JavaScript已被大多数的浏览器所支持。 [6] 不同于服务器端脚本语言，例如[PHP](https://baike.baidu.com/item/PHP/9337?fromModule=lemma_inlink)与[ASP](https://baike.baidu.com/item/ASP/128906?fromModule=lemma_inlink)，JavaScript主要被作为客户端脚本语言在用户的浏览器上运行，不需要服务器的支持。所以在早期程序员比较倾向于使用JavaScript以减少对服务器的负担，而与此同时也带来另一个问题，安全性。

-----------------------------





## JavaScript 基础



### JavaScript 实例

学习 100 多个 JavaScript 实例！

在实例页面中，您可以点击 "尝试一下" 来查看 JavaScript 在线实例。

- [JavaScript 实例](https://www.runoob.com/js/js-examples.html)
- [JavaScript 对象实例](https://www.runoob.com/js/js-ex-objects.html)
- [JavaScript 浏览器支持实例](https://www.runoob.com/js/js-ex-browser.html)
- [JavaScript HTML DOM 实例](https://www.runoob.com/js/js-ex-dom.html)



### JavaScript 测验

在菜鸟教程中测试您的 JavaScript 技能！

[JavaScript 测验](https://www.runoob.com/quiz/javascript-quiz.html)



### JavaScript 参考手册

在菜鸟教程中，我们为您提供完整的 JavaScript 对象、浏览器对象、HTML DOM 对象参考手册。

以下手册包含了每个对象、属性、方法的实例。

- [JavaScript 内置对象](https://www.runoob.com/jsref/jsref-tutorial.html)
- [Browser 对象](https://www.runoob.com/jsref/jsref-tutorial.html)
- [HTML DOM 对象](https://www.runoob.com/jsref/jsref-tutorial.html)



-------------------------





## JavaScript 用法

HTML 中的 Javascript 脚本代码必须位于` <script>` 与 `</script>` 标签之间。

Javascript 脚本代码可被放置在 HTML 页面的` <body> `和 `<head>`*部分中。



### `<script>` 标签

如需在 HTML 页面中插入 JavaScript，请使用` <script>` 标签。

`<script>` 和 `</script>` 会告诉 JavaScript 在何处开始和结束。

`<script>` 和 `</script>` 之间的代码行包含了 JavaScript:

```js
<script>
alert("我的第一个 JavaScript");
</script>
```



###  JavaScript 函数和事件

上面例子中的 JavaScript 语句，会在页面加载时执行。

通常，我们需要在某个事件发生时执行代码，比如当用户点击按钮时。

如果我们把 JavaScript 代码放入函数中，就可以在事件发生时调用该函数。



### 在` <head>` 或者` <body>` 的JavaScript

您可以在 HTML 文档中放入不限数量的脚本。

脚本可位于 HTML 的` <body>` 或 `<head> `部分中，或者同时存在于两个部分中。

通常的做法是把函数放入` <head>` 部分中，或者放在页面底部。这样就可以把它们安置到同一处位置，不会干扰页面的内容。



### 外部的 JavaScript（导入外部JS）

也可以把脚本保存到外部文件中。外部文件通常包含被多个网页使用的代码。

外部 JavaScript 文件的文件扩展名是 .js。

如需使用外部文件，请在 `<script> `标签的 "src" 属性中设置该 .js 文件：

```js
<!DOCTYPE html>
<html>
<body>
<script src="myScript.js"></script>
</body>
</html>
```

你可以将脚本放置于 `<head> `或者` <body>`中，放在` <script> `标签中的脚本与外部引用的脚本运行效果完全一致。

 **外部脚本不能包含 `<script>` 标签。**



----------------------------------



## JavaScript 输出

JavaScript 没有任何打印或者输出的函数。



### JavaScript 显示数据

JavaScript 可以通过不同的方式来输出数据：

- 使用 **window.alert()** 弹出警告框。
- 使用 **document.write()** 方法将**内容写到 HTML 文档中**。（直接显示）
- 使用 **innerHTML** **写入到 HTML 元素**。（写到对应的元素里面）
- 使用 **console.log()** 写入到浏览器的控制台。



### 使用 window.alert()

你可以弹出警告框来显示数据：

```js
<script>
window.alert(5 + 6);
</script>	
```



### 操作 HTML 元素

**如需从 JavaScript 访问某个 HTML 元素，您可以使用 document.getElementById(*id*) 方法。**

请使用 "id" 属性来标识 HTML 元素，并 innerHTML 来获取或插入元素内容：

**理解：通过“id”获取该元素，然后调用innerHTML 对该元素进行修改**

```js
<p id="demo">我的第一个段落</p>

<script>
document.getElementById("demo").innerHTML = "段落已修改。";
</script>
```

**document.getElementById("demo")** 是使用 id 属性来查找 HTML 元素的 JavaScript 代码 。

**innerHTML = "段落已修改。"** 是用于修改元素的 HTML 内容(innerHTML)的 JavaScript 代码。





### 写到 HTML 文档

出于测试目的，您可以将JavaScript直接写在HTML 文档中：

```js
<script>
document.write(Date());
</script>
```



### 写到控制台

如果您的浏览器支持调试，你可以使用 **console.log()** 方法在浏览器中显示 JavaScript 值。

浏览器中使用 F12 来启用调试模式， 在调试窗口中点击 "Console" 菜单。

```javascript
<script>
a = 5;
b = 6;
c = a + b;
console.log(c);
</script>
```

-------------------------------------------





## JavaScript 语法

JavaScript 是一个程序语言。语法规则定义了语言结构。**JavaScript 是一个脚本语言**。它是一个轻量级，但功能强大的编程语言。



### JavaScript 字面量

在编程语言中，**一般固定值称为字面量，如 3.14**。

**数字（Number）字面量** 可以是整数或者是小数，或者是科学计数(e)。

**字符串（String）字面量** 可以使用<u>单引号或双引号</u>:

**表达式字面量** 用于计算：

```js
5 + 6

5 * 10
```

**数组（Array）字面量** 定义一个数组：

```js
[40, 100, 1, 5, 25, 10]
```

**对象（Object）字面量** 定义一个对象：

```js
{firstName:"John", lastName:"Doe", age:50, eyeColor:"blue"}
```

**函数（Function）字面量** 定义一个函数：

```js
function myFunction(a, b) { return a * b;}
```





### JavaScript 变量

在编程语言中，变量用于存储数据值。**JavaScript 使用关键字 var 来定义变量， 使用等号来为变量赋值：**

```js
var x, length
x = 5
length = 6
```



### JavaScript 操作符

JavaScript使用 **算术运算符** 来计算值:

```js
(5 + 6) * 10
```

**JavaScript语言有多种类型的运算符：**

| 类型                   | 实例      | 描述                   |
| :--------------------- | :-------- | :--------------------- |
| 赋值，算术和位运算符   | = + - * / | 在 JS 运算符中描述     |
| 条件，比较及逻辑运算符 | == != < > | 在 JS 比较运算符中描述 |



### JavaScript 语句

在 HTML 中，JavaScript 语句用于向浏览器发出命令。**语句是用分号分隔：**

```js
x = 5 + 6;
y = x * 10;
```



### JavaScript 关键字

JavaScript 关键字用于标识要执行的操作。和其他任何编程语言一样，JavaScript 保留了一些关键字为自己所用。

**更多的关键字：**[JavaScript 语法 | 菜鸟教程 (runoob.com)](https://www.runoob.com/js/js-syntax.html)



### JavaScript 注释

不是所有的 JavaScript 语句都是"命令"。双斜杠 **//** 后的内容将会被浏览器忽略：



### JavaScript 数据类型

JavaScript 有多种数据类型：数字，字符串，数组，对象等等

```js
var length = 16;                                  // Number 通过数字字面量赋值
var points = x * 10;                              // Number 通过表达式字面量赋值
var lastName = "Johnson";                         // String 通过字符串字面量赋值
var cars = ["Saab", "Volvo", "BMW"];              // Array  通过数组字面量赋值
var person = {firstName:"John", lastName:"Doe"};  // Object 通过对象字面量赋值
```



### JavaScript 函数

JavaScript 语句可以写在函数内，函数可以重复引用：**引用一个函数** = 调用函数(执行函数内的语句)。

```js
function myFunction(a, b) {
    return a * b;                                // 返回 a 乘以 b 的结果
}
```



### JavaScript 字母大小写

**JavaScript 对大小写是敏感的。**

当编写 JavaScript 语句时，请留意是否关闭大小写切换键。函数 **getElementById** 与 **getElementbyID** 是不同的。同样，变量 **myVariable** 与 **MyVariable** 也是不同的。

**JavaScript 中，常见的是驼峰法的命名规则，如 lastName (而不是lastname)。**



### JavaScript 字符集

**JavaScript 使用 Unicode 字符集。**Unicode 覆盖了所有的字符，包含标点等字符。

[HTML UTF-8 参考手册 | 菜鸟教程 (runoob.com)](https://www.runoob.com/charsets/ref-html-utf8.html)

-----------------------------





## JavaScript 语句

JavaScript 语句向浏览器发出的命令。**语句的作用是告诉浏览器该做什么。**



### 分号 ;

**分号用于分隔 JavaScript 语句。通常我们在每条可执行的语句结尾添加分号。**使用分号的另一用处是在一行中编写多条语句。

**在 JavaScript 中，用分号来结束语句是可选的。**



### JavaScript 代码

JavaScript 代码是 JavaScript 语句的序列。**浏览器按照编写顺序依次执行每条语句。**



### JavaScript 代码块

JavaScript 可以分批地组合起来。**代码块以左花括号开始，以右花括号结束。代码块的作用是一并地执行语句序列。**

```js
function myFunction()
{
    document.getElementById("demo").innerHTML="你好Dolly";
    document.getElementById("myDIV").innerHTML="你最近怎么样?";
}
```



### JavaScript 语句标识符

JavaScript 语句通常以一个 **语句标识符** 为开始，并执行该语句。**语句标识符是保留关键字不能作为变量名使用。**

| 语句         | 描述                                                         |
| :----------- | :----------------------------------------------------------- |
| break        | 用于跳出循环。                                               |
| catch        | 语句块，在 try 语句块执行出错时执行 catch 语句块。           |
| continue     | 跳过循环中的一个迭代。                                       |
| do ... while | 执行一个语句块，在条件语句为 true 时继续执行该语句块。       |
| for          | 在条件语句为 true 时，可以将代码块执行指定的次数。           |
| for ... in   | 用于遍历数组或者对象的属性（对数组或者对象的属性进行循环操作）。 |
| function     | 定义一个函数                                                 |
| if ... else  | 用于基于不同的条件来执行不同的动作。                         |
| return       | 退出函数                                                     |
| switch       | 用于基于不同的条件来执行不同的动作。                         |
| throw        | 抛出（生成）错误 。                                          |
| try          | 实现错误处理，与 catch 一同使用。                            |
| var          | 声明一个变量。                                               |
| while        | 当条件语句为 true 时，执行语句块。                           |



### 对代码行进行折行

您可以在文本字符串中使用反斜杠对代码行进行换行。下面的例子会正确地显示：

```js
document.write("你好 \
世界!");
```



---------------------



## JavaScript 注释

JavaScript 注释可用于提高代码的可读性。



### 单行注释

单行注释以 **//** 开头



### 多行注释

多行注释以 **/\*** 开始，以 ***/** 结尾。



-----------------------------------



## JavaScript 变量

变量是用于存储信息的"容器"。



### JavaScript 变量

与代数一样，JavaScript 变量可用于存放值（比如 x=5）和表达式（比如 z=x+y）。

**变量可以使用短名称（比如 x 和 y），也可以使用描述性更好的名称（比如 age, sum, totalvolume）。**

- 变量**必须以字母开头**
- 变量**也能以 $ 和 _ 符号开头**（不过我们不推荐这么做）
- 变量名称**对大小写敏感**（y 和 Y 是不同的变量）



### JavaScript 数据类型

JavaScript 变量还能保存其他数据类型，比如文本值 (name="Bill Gates")。在 JavaScript 中，类似 "Bill Gates" 这样一条文本被称为字符串。

JavaScript 变量有很多种类型，但是现在，我们只关注数字和字符串。**当您向变量分配文本值时，应该用双引号或单引号包围这个值。**

当您向变量赋的值是数值时，不要使用引号。如果您用引号包围数值，该值会被作为文本来处理。



### 声明（创建） JavaScript 变量

在 JavaScript 中创建变量通常称为"声明"变量。**我们使用 var 关键词来声明变量：变量声明之后，该变量是空的（它没有值）。**如需向变量赋值，请使用等号：

```JS
var carname;
carname="Volvo";
```

不过，您也可以在声明变量时对其赋值

```JS
var carname="Volvo";
```



### 一条语句，多个变量

您可以在一条语句中声明很多变量。该语句以 var 开头，并使用逗号分隔变量即可：

```JS
var lastname="Doe", age=30, job="carpenter";
```

一条语句中声明的多个变量**不可以**同时赋同一个值:

```JS
var x,y,z=1;
// x,y 为 undefined， z 为 1。
```



### Value = undefined

在计算机程序中，经常会声明无值的变量。**未使用值来声明的变量，其值实际上是 undefined。**



### 重新声明 JavaScript 变量

如果重新声明 JavaScript 变量，该变量的值不会丢失：

在以下两条语句执行后，变量 carname 的值依然是 "Volvo"：

```JS
var carname="Volvo";
var carname;
```



----------------------------



## JavaScript 数据类型

**值类型(基本类型)**：字符串（String）、数字(Number)、布尔(Boolean)、空（Null）、未定义（Undefined）、Symbol。

**引用数据类型（对象类型）**：对象(Object)、数组(Array)、函数(Function)，还有两个特殊的对象：正则（RegExp）和日期（Date）。

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301121211086.png" alt="img" style="zoom: 50%;" />

*Symbol 是 ES6 引入了一种新的原始数据类型，表示独一无二的值。*



### JavaScript 拥有动态类型

JavaScript 拥有动态类型。这意味着相同的变量可用作不同的类型：

```JS
var x;               // x 为 undefined
var x = 5;           // 现在 x 为数字
var x = "John";      // 现在 x 为字符串
```

变量的数据类型可以使用 **typeof** 操作符来查看：

```JS
typeof "John"                // 返回 string
typeof 3.14                  // 返回 number
typeof false                 // 返回 boolean
typeof [1,2,3,4]             // 返回 object
typeof {name:'John', age:34} // 返回 object
```



### JavaScript 字符串

字符串是存储字符（比如 "Bill Gates"）的变量。**字符串可以是引号中的任意文本。您可以使用单引号或双引号：**

**您可以在字符串中使用引号，只要不匹配包围字符串的引号即可：**

```JS
var answer="It's alright";
var answer="He is called 'Johnny'";
var answer='He is called "Johnny"';
```



### JavaScript 数字

JavaScript 只有一种数字类型。数字可以带小数点，也可以不带：

```JS
var x1=34.00;      //使用小数点来写
var x2=34;         //不使用小数点来写
// 极大或极小的数字可以通过科学（指数）计数法来书写：
var y=123e5;      // 12300000
var z=123e-5;     // 0.00123
```



### JavaScript 布尔

布尔（逻辑）只能有两个值：true 或 false。



### JavaScript 数组

下面的代码创建名为 cars 的数组：

```JS
var cars=new Array();
cars[0]="Saab";
var cars=new Array("Saab","Volvo","BMW");
var cars=["Saab","Volvo","BMW"];
```



### JavaScript 对象

对象由花括号分隔。在括号内部，对象的属性以名称和值对的形式 (name : value) 来定义。属性由逗号分隔：

```JS
var person={firstname:"John", lastname:"Doe", id:5566};
// 上面例子中的对象 (person) 有三个属性：firstname、lastname 以及 id。
// 空格和折行无关紧要。声明可横跨多行：
var person={
firstname : "John",
lastname  : "Doe",
id        :  5566
};
```

对象属性有两种寻址方式：

```JS
name=person.lastname;
name=person["lastname"];
```



### Undefined 和 Null

Undefined 这个值表示变量不含有值。**可以通过将变量的值设置为 null 来清空变量。**



### 声明变量类型

当您声明新变量时，可以使用关键词 "new" 来声明其类型：

```JS
var carname=new String;
var x=      new Number;
var y=      new Boolean;
var cars=   new Array;
var person= new Object;
```

 **JavaScript 变量均为对象。当您声明一个变量时，就创建了一个新的对象。**



-------------------------------



## JavaScript 对象

**JavaScript 对象是拥有属性和方法的数据。**



### JavaScript 对象

在 JavaScript中，几乎所有的事物都是对象。

以下代码为变量 **car** 设置值为 "Fiat" :

```JS
var car = "Fiat";
```

对象也是一个变量，但对象可以包含多个值（多个变量），每个值以 **name:value** 对呈现。

```JS
var car = {name:"Fiat", model:500, color:"white"};
```





### 对象定义

你可以使用字符来定义和创建 JavaScript 对象:

```JS
var person = {firstName:"John", lastName:"Doe", age:50, eyeColor:"blue"};
```



### 对象属性

可以说 "JavaScript 对象是变量的容器"。但是，我们通常认为 "JavaScript 对象是键值对的容器"。键值对通常写法为 **name : value** (键与值以冒号分割)。键值对在 JavaScript 对象通常称为 **对象属性**。



### 访问对象属性

你可以通过两种方式访问对象属性:

```JS
person.lastName;
person["lastName"];
```



### 对象方法

对象的方法定义了一个函数，并作为对象的属性存储。**对象方法通过添加 () 调用 (作为一个函数)。**该实例访问了 person 对象的 fullName() 方法:

```JS
name = person.fullName();
```

如果你要访问 person 对象的 fullName 属性，它将作为一个定义函数的字符串返回：

```JS
name = person.fullName;
```

**JavaScript 对象是属性和方法的容器。**



### 访问对象方法

你可以使用以下语法创建对象方法：

```JS
methodName : function() {
    // 代码 
}
```



### 创建含方法的对象

```js
var person = {
    firstName: "John",
    lastName : "Doe",
    id : 5566,
    fullName : function() 
	{
       return this.firstName + " " + this.lastName;
    }
};
```



-------------------------------------



## JavaScript 函数

函数是由事件驱动的或者当它被调用时执行的可重复使用的代码块。



### JavaScript 函数语法

函数就是包裹在花括号中的代码块，前面使用了关键词 function：

```js
function functionname()
{
    // 执行代码
}
```

当调用该函数时，会执行函数内的代码。**可以在某事件发生时直接调用函数（比如当用户点击按钮时），并且可由 JavaScript 在任何位置进行调用。**



### 调用带参数的函数

在调用函数时，您可以向其传递值，这些值被称为参数。这些参数可以在函数中使用**。您可以发送任意多的参数，由逗号 (,) 分隔：**

```js
myFunction(argument1,argument2)
```

当您声明函数时，请把参数作为变量来声明：

```js
function myFunction(var1,var2)
{
	// 代码
}
```

**变量和参数必须以一致的顺序出现。第一个变量就是第一个被传递的参数的给定的值，以此类推。**



### 带有返回值的函数

**有时，我们会希望函数将值返回调用它的地方。通过使用 return 语句就可以实现。**在使用 return 语句时，函数会停止执行，并返回指定的值。

```js
// 函数会返回值 5。
function myFunction()
{
    var x=5;
    return x;
}
```

**注意： 整个 JavaScript 并不会停止执行，仅仅是函数。JavaScript 将继续执行代码，从调用函数的地方。**



函数调用将被返回值取代：

```js
var myVar=myFunction();
```

myVar 变量的值是 5，也就是函数 "myFunction()" 所返回的值。即使不把它保存为变量，您也可以使用返回值：

```js
document.getElementById("demo").innerHTML=myFunction();
```





### 局部 JavaScript 变量

**在 JavaScript 函数内部声明的变量（使用 var）是*局部*变量，所以只能在函数内部访问它**。（该变量的作用域是局部的）。

您可以在不同的函数中使用名称相同的局部变量，因为只有声明过该变量的函数才能识别出该变量。

只要函数运行完毕，本地变量就会被删除。



### 全局 JavaScript 变量

在函数外声明的变量是*全局*变量，网页上的所有脚本和函数都能访问它。



### JavaScript 变量的生存期

JavaScript 变量的生命期从它们被声明的时间开始。

局部变量会在函数运行以后被删除。

**全局变量会在页面关闭后被删除。**



### 向未声明的 JavaScript 变量分配值

如果您把值赋给尚未声明的变量，该变量将被自动作为 window 的一个属性。

这条语句：

```js
carname="Volvo";
```

将声明 window 的一个属性 carname。**非严格模式下给未声明变量赋值创建的全局变量，是全局对象的可配置属性，可以删除。**

```js
var var1 = 1; // 不可配置全局属性
var2 = 2; // 没有使用 var 声明，可配置全局属性

console.log(this.var1); // 1
console.log(window.var1); // 1
console.log(window.var2); // 2

delete var1; // false 无法删除
console.log(var1); //1

delete var2; 
console.log(delete var2); // true
console.log(var2); // 已经删除 报错变量未定义
```



----------------------



## JavaScript 作用域

作用域是可访问变量的集合。



### JavaScript 作用域

在 JavaScript 中, 对象和函数同样也是变量。

**在 JavaScript 中, 作用域为可访问变量，对象，函数的集合。**

JavaScript 函数作用域: 作用域在函数内修改。





### JavaScript 局部作用域

变量在函数内声明，变量为局部变量，具有局部作用域。

局部变量：只能在函数内部访问。



### JavaScript 全局变量

变量在函数外定义，即为全局变量。

全局变量有 **全局作用域**: 网页中所有脚本和函数均可使用。 



**如果变量在函数内没有声明（没有使用 var 关键字），该变量为全局变量。**

**以下实例中 carName 在函数内，但是为全局变量。**

```js
// 此处可调用 carName 变量
 
function myFunction() {
    carName = "Volvo";
    // 此处可调用 carName 变量
}
```

**因为把值赋给尚未声明的变量，该变量将被自动作为 window 的一个属性，其为全局变量。**



### JavaScript 变量生命周期

JavaScript 变量生命周期在它声明时初始化。==局部变量在函数执行完毕后销毁。全局变量在页面关闭后销毁。==



### HTML 中的全局变量

在 HTML 中, 全局变量是 window 对象，所以window 对象可以调用函数内创建的未带var的变量。

**注意：**所有数据变量都属于 window 对象。

```js
//此处可使用 window.carName
 
function myFunction() {
    carName = "Volvo";
}
```





## JavaScript 事件

------

HTML 事件是发生在 HTML 元素上的事情。当在 HTML 页面中使用 JavaScript 时， JavaScript 可以触发这些事件。



### HTML 事件

HTML 事件可以是浏览器行为，也可以是用户行为。

以下是 HTML 事件的实例：

- HTML 页面完成加载
- HTML input 字段改变时
- HTML 按钮被点击

通常，当事件发生时，你可以做些事情。在事件触发时 JavaScript 可以执行一些代码。

**HTML 元素中可以添加事件属性，使用 JavaScript 代码来添加 HTML 元素。**

单引号:

```js
<some-HTML-element some-event='JavaScript 代码'>
```

双引号:

```js
<some-HTML-element some-event="JavaScript 代码">
```

在以下实例中，按钮元素中添加了 onclick 属性 (并加上代码):

```js
<button onclick="getElementById('demo').innerHTML=Date()">现在的时间是?</button>
```

在下一个实例中，**代码将修改自身元素的内容 (使用 this.innerHTML):**

```js
<button onclick="this.innerHTML=Date()">现在的时间是?</button>
```



### 常见的HTML事件

下面是一些常见的HTML事件的列表:

| 事件        | 描述                                 |
| :---------- | :----------------------------------- |
| onchange    | HTML 元素改变                        |
| onclick     | 用户点击 HTML 元素                   |
| onmouseover | 鼠标指针移动到指定的元素上时发生     |
| onmouseout  | 用户从一个 HTML 元素上移开鼠标时发生 |
| onkeydown   | 用户按下键盘按键                     |
| onload      | 浏览器已完成页面的加载               |



### JavaScript 可以做什么?

事件可以用于处理表单验证，用户输入，用户行为及浏览器动作:

- 页面加载时触发事件
- 页面关闭时触发事件
- 用户点击按钮执行动作
- 验证用户输入内容的合法性
- 等等 ...

可以使用多种方法来执行 JavaScript 事件代码：

- HTML 事件属性可以直接执行 JavaScript 代码
- HTML 事件属性可以调用 JavaScript 函数
- 你可以为 HTML 元素指定自己的事件处理程序
- 你可以阻止事件的发生。
- 等等 ...





## JavaScript 字符串

------

JavaScript 字符串用于存储和处理文本。



### JavaScript 字符串

字符串可以存储一系列字符，如 "John Doe"。

**字符串可以是插入到引号中的任何字符。你可以使用单引号或双引号：**

你可以使用索引位置来访问字符串中的每个字符：

```js
var character = carname[7];
```

字符串的索引从 0 开始，这意味着第一个字符索引值为 [0],第二个为 [1], 以此类推。

你也可以在字符串添加转义字符来使用引号：



### 字符串长度

可以使用内置属性 **length** 来计算字符串的长度：

```js
var txt = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
var sln = txt.length;
```



### 特殊字符

| 代码 | 输出        |
| :--- | :---------- |
| \'   | 单引号      |
| \"   | 双引号      |
| \\   | 反斜杠      |
| \n   | 换行        |
| \r   | 回车        |
| \t   | tab(制表符) |
| \b   | 退格符      |
| \f   | 换页符      |



### 字符串可以是对象

通常， JavaScript 字符串是原始值，可以使用字符创建： **var firstName = "John"**

但我们也可以使用 new 关键字将字符串定义为一个对象： **var firstName = new String("John")**

 ```js
 var x = "John";
 var y = new String("John");
 typeof x // 返回 String
 typeof y // 返回 Object
 ```

**不要创建 String 对象。它会拖慢执行速度，并可能产生其他副作用：**

```js
var x = "John";             
var y = new String("John");
(x === y) // 结果为 false，因为 x 是字符串，y 是对象
```

**=== 为绝对相等，即数据类型与值都必须相等。**



### 字符串属性和方法

**原始值字符串，如 "John", 没有属性和方法(因为他们不是对象)。**

原始值可以使用 JavaScript 的属性和方法，**因为 JavaScript 在执行方法和属性时可以把原始值当作对象。**

## 字符串属性

| 属性        | 描述                       |
| :---------- | :------------------------- |
| constructor | 返回创建字符串属性的函数   |
| length      | 返回字符串的长度           |
| prototype   | 允许您向对象添加属性和方法 |





## JavaScript 运算符

------

**运算符 = 用于赋值。**

**运算符 + 用于加值。**



### 用于字符串的 + 运算符

\+ 运算符用于把文本值或字符串变量加起来（连接起来）。

如需把两个或多个字符串变量连接起来，请使用 **+** 运算符。



### 对字符串和数字进行加法运算

两个数字相加，返回数字相加的和，如果数字与字符串相加，返回字符串





### JavaScript 比较 和 逻辑运算符

------

**比较和逻辑运算符用于测试 *true* 或者 *false*。**

| ===  | 绝对等于（值和类型均相等）                         | x==="5" | *false* | [实例 »](https://www.runoob.com/try/try.php?filename=tryjs_comparison3) |
| ---- | -------------------------------------------------- | ------- | ------- | ------------------------------------------------------------ |
|      |                                                    | x===5   | *true*  |                                                              |
| !==  | 不绝对等于（值和类型有一个不相等，或两个都不相等） | x!=="5" | *true*  |                                                              |



**可以在条件语句中使用比较运算符对值进行比较，然后根据结果来采取行动：**

```js
if (age<18) x="Too young";
```



### 逻辑运算符

逻辑运算符用于测定变量或值之间的逻辑。

给定 x=6 以及 y=3，下表解释了逻辑运算符：

| 运算符 | 描述 | 例子                      |
| :----- | :--- | :------------------------ |
| &&     | and  | (x < 10 && y > 1) 为 true |
| \|\|   | or   | `(x==5 || y==5) `为 false |
| !      | not  | !(x==y) 为 true           |



### 条件运算符

JavaScript 还包含了基于某些条件对变量进行赋值的条件运算符。

```js
variablename=(condition)?value1:value2 
```

如果变量 age 中的值小于 18，则向变量 voteable 赋值 "年龄太小"，否则赋值 "年龄已达到"。

```js
voteable=(age<18)?"年龄太小":"年龄已达到";
```





## JavaScript if...Else 语句

------

条件语句用于基于不同的条件来执行不同的动作。



### 条件语句

通常在写代码时，您总是需要为不同的决定来执行不同的动作。您可以在代码中使用条件语句来完成该任务。

在 JavaScript 中，我们可使用以下条件语句：

- **if 语句** - 只有当指定条件为 true 时，使用该语句来执行代码
- **if...else 语句** - 当条件为 true 时执行代码，当条件为 false 时执行其他代码
- **if...else if....else 语句**- 使用该语句来选择多个代码块之一来执行
- **switch 语句** - 使用该语句来选择多个代码块之一来执行



------

### if 语句

只有当指定条件为 true 时，该语句才会执行代码。

```js
if (condition)
{
    当条件为 true 时执行的代码
}
```



### if...else 语句

请使用 if....else 语句在条件为 true 时执行代码，在条件为 false 时执行其他代码。

```js
if (condition)
{
    当条件为 true 时执行的代码
}
else
{
    当条件不为 true 时执行的代码
}
```



### if...else if...else 语句

使用 if....else if...else 语句来选择多个代码块之一来执行。

```js
if (condition1)
{
    当条件 1 为 true 时执行的代码
}
else if (condition2)
{
    当条件 2 为 true 时执行的代码
}
else
{
  当条件 1 和 条件 2 都不为 true 时执行的代码
}
```



### JavaScript switch 语句

请使用 switch 语句来选择要执行的多个代码块之一。

```js
switch(n)
{
    case 1:
        执行代码块 1
        break;
    case 2:
        执行代码块 2
        break;
    default:
        与 case 1 和 case 2 不同时执行的代码
}
```





## JavaScript for 循环

------

循环可以将代码块执行指定的次数。



### 不同类型的循环

JavaScript 支持不同类型的循环：

- **for** - 循环代码块一定的次数
- **for/in** - 循环遍历对象的属性
- **while** - 当指定的条件为 true 时循环指定的代码块
- **do/while** - 同样当指定的条件为 true 时循环指定的代码块



### For 循环

for 循环是您在希望创建循环时常会用到的工具。

```js
for (语句 1; 语句 2; 语句 3)
{
    被执行的代码块
}
```

**语句 1** （代码块）开始前执行

**语句 2** 定义运行循环（代码块）的条件

**语句 3** 在循环（代码块）已被执行之后执行



### For/In 循环

JavaScript for/in 语句循环**遍历对象的属性**：

```js
var person={fname:"Bill",lname:"Gates",age:56}; 
 
for (x in person)  // x 为属性名，即"fname"
{
    txt=txt + person[x];
}
```



### while 循环

while 循环会在指定条件为真时循环执行代码块。

```js
while (i<5)
{
    x=x + "The number is " + i + "<br>";
    i++;
}
```



### do/while 循环

do/while 循环是 while 循环的变体。该循环会在检查条件是否为真之前执行一次代码块，然后如果条件为真的话，就会重复这个循环。

```js
do
{
    x=x + "The number is " + i + "<br>";
    i++;
}
while (i<5);
```



### break 和 continue 语句

break 语句用于跳出循环。

continue 用于跳过循环中的一个迭代。



## JavaScript 标签

正如您在 switch 语句那一章中看到的，可以对 JavaScript 语句进行标记。

如需标记 JavaScript 语句，请在语句之前加上冒号：

```js
label:
statements
```

**break 和 continue 语句仅仅是能够跳出代码块的语句。**

语法:

```js
break labelname; 
 
continue labelname;
```

continue 语句（带有或不带标签引用）只能用在循环中。

break 语句（不带标签引用），只能用在循环或 switch 中。

通过标签引用，break 语句可用于跳出任何 JavaScript 代码块：





## typeof, null, 和 undefined



### typeof 操作符

你可以使用 typeof 操作符来检测变量的数据类型。

```js
typeof "John"                // 返回 string
typeof 3.14                  // 返回 number
typeof false                 // 返回 boolean
typeof [1,2,3,4]             // 返回 object
typeof {name:'John', age:34} // 返回 object
```

  在JavaScript中，数组是一种特殊的对象类型。 因此 typeof [1,2,3,4] 返回 object。 



### null

在 JavaScript 中 null 表示 "什么都没有"。**null是一个只有一个值的特殊类型。表示一个空对象引用。**

**用 typeof 检测 null 返回是object。**

**你可以设置为 null 来清空对象或者设置为 undefined 来清空对象:**

```js
var person = null;           // 值为 null(空), 但类型为对象
var person = undefined;     // 值为 undefined, 类型为 undefined
```



### undefined

在 JavaScript 中, **undefined** 是一个没有设置值的变量。

**typeof** 一个没有值的变量会返回 **undefined**。

```js
var person;                  // 值为 undefined(空), 类型是undefined
```

任何变量都可以通过设置值为 **undefined** 来清空。 类型为 **undefined**.

```js
person = undefined;          // 值为 undefined, 类型是undefined
```



### undefined 和 null 的区别

**null 和 undefined 的值相等，但类型不等：**

```js
typeof undefined             // undefined
typeof null                  // object
null === undefined           // false
null == undefined            // true
```





## JavaScript 类型转换

------

Number() 转换为数字， String() 转换为字符串， Boolean() 转换为布尔值。



### typeof 操作符

你可以使用 **typeof** 操作符来查看 JavaScript 变量的数据类型。



### constructor 属性

**constructor** 属性返回所有 JavaScript 变量的构造函数。

```js
"John".constructor                 // 返回函数 String()  { [native code] }
(3.14).constructor                 // 返回函数 Number()  { [native code] }
false.constructor                  // 返回函数 Boolean() { [native code] }
[1,2,3,4].constructor              // 返回函数 Array()   { [native code] }
{name:'John', age:34}.constructor  // 返回函数 Object()  { [native code] }
new Date().constructor             // 返回函数 Date()    { [native code] }
function () {}.constructor         // 返回函数 Function(){ [native code] }
```

你可以使用 constructor 属性来查看对象是否为数组 (包含字符串 "Array"):



### JavaScript 类型转换

JavaScript 变量可以转换为新变量或其他数据类型：

- 通过使用 JavaScript 函数
- 通过 JavaScript 自身自动转换



### 将数字转换为字符串

全局方法 **String()** 可以将数字转换为字符串。

该方法可用于任何类型的数字，字母，变量，表达式：

Number 方法 **toString()** 也是有同样的效果。



### 将布尔值转换为字符串

全局方法 **String()** 可以将布尔值转换为字符串。

Boolean 方法 **toString()** 也有相同的效果。



### 将日期转换为字符串

Date() 返回字符串。

全局方法 String() 可以将日期对象转换为字符串。

Date 方法 **toString()** 也有相同的效果。



### 将字符串转换为数字

全局方法 **Number()** 可以将字符串转换为数字。

字符串包含数字(如 "3.14") 转换为数字 (如 3.14).

空字符串转换为 0。

其他的字符串会转换为 NaN (不是个数字)。



### 一元运算符 +

**Operator +** 可用于将变量转换为数字：

```js
var y = "5";      // y 是一个字符串
var x = + y;      // x 是一个数字
```

如果变量不能转换，它仍然会是一个数字，但值为 NaN (不是一个数字):

```js
var y = "John";   // y 是一个字符串
var x = + y;      // x 是一个数字 (NaN)
```



### 将布尔值转换为数字

全局方法 **Number()** 可将布尔值转换为数字。



### 将日期转换为数字

全局方法 **Number()** 可将日期转换为数字。

日期方法 **getTime()** 也有相同的效果。



### 自动转换类型

当 JavaScript 尝试操作一个 "错误" 的数据类型时，会自动转换为 "正确" 的数据类型。



### 自动转换为字符串

当你尝试输出一个对象或一个变量时 JavaScript 会自动调用变量的 toString() 方法：





## JavaScript 正则表达式

正则表达式（英语：Regular Expression，在代码中常简写为regex、regexp或RE）使用单个字符串来描述、匹配一系列符合某个句法规则的字符串搜索模式。搜索模式可用于文本搜索和文本替换。



###  语法

```js
/正则表达式主体/修饰符(可选)
```

实例

```js
var patt = /runoob/i
```

实例解析：

**/runoob/i** 是一个正则表达式。

**runoob** 是一个**正则表达式主体** (用于检索)。

**i** 是一个**修饰符** (搜索不区分大小写)。



### 使用字符串方法

在 JavaScript 中，正则表达式通常用于两个字符串方法 : search() 和 replace()。

**search()** 方法用于检索字符串中指定的子字符串，或检索与正则表达式相匹配的子字符串，并返回子串的起始位置。

**replace()** 方法用于在字符串中用一些字符串替换另一些字符串，或替换一个与正则表达式匹配的子串。



#### search() 方法使用正则表达式

```js
// 使用正则表达式搜索 "Runoob" 字符串，且不区分大小写：
var str = "Visit Runoob!"; 
var n = str.search(/Runoob/i);  // 返回6
```

#### search() 方法使用字符串

search 方法可使用字符串作为参数。字符串参数会转换为正则表达式：

```js
var str = "Visit Runoob!"; 
var n = str.search("Runoob");
```



#### replace() 方法使用正则表达式

```js
var str = document.getElementById("demo").innerHTML; 
var txt = str.replace(/microsoft/i,"Runoob");
```

#### replace() 方法使用字符串

replace() 方法将接收字符串作为参数：

```js
var str = document.getElementById("demo").innerHTML; 
var txt = str.replace("Microsoft","Runoob");
```



### 正则表达式修饰符

**修饰符** 可以在全局搜索中不区分大小写:

| 修饰符 | 描述                                                     |
| :----- | :------------------------------------------------------- |
| i      | 执行对大小写不敏感的匹配。                               |
| g      | 执行全局匹配（查找所有匹配而非在找到第一个匹配后停止）。 |
| m      | 执行多行匹配。                                           |



### 使用 RegExp 对象

在 JavaScript 中，RegExp 对象是一个预定义了属性和方法的正则表达式对象。



#### 使用 test()

test() 方法是一个正则表达式方法。

test() 方法用于检测一个字符串是否匹配某个模式，如果字符串中含有匹配的文本，则返回 true，否则返回 false。

以下实例用于搜索字符串中的字符 "e"：

```js
var patt = /e/;
patt.test("The best things in life are free!");
// 字符串中含有 "e"，所以该实例输出为true
```



#### 使用 exec()

exec() 方法是一个正则表达式方法。

exec() 方法用于检索字符串中的正则表达式的匹配。

该函**数返回一个数组**，其中存放匹配的结果。如果未找到匹配，则返回值为 null。

以下实例用于搜索字符串中的字母 "e":

```js
/e/.exec("The best things in life are free!");
// 字符串中含有 "e"，所以该实例输出为: e
```





## JavaScript 错误 - throw、try 和 catch

------

**try** 语句测试代码块的错误。

**catch** 语句处理错误。

**throw** 语句创建自定义错误。

**finally** 语句在 try 和 catch 语句之后，无论是否有触发异常，该语句都会执行。



###  JavaScript 错误

当 JavaScript 引擎执行 JavaScript 代码时，会发生各种错误。

可能是语法错误，通常是程序员造成的编码错误或错别字。

可能是拼写错误或语言中缺少的功能（可能由于浏览器差异）。

可能是由于来自服务器或用户的错误输出而导致的错误。

当然，也可能是由于许多其他不可预知的因素。



### JavaScript 抛出（throw）错误

当错误发生时，当事情出问题时，JavaScript 引擎通常会停止，并生成一个错误消息。

描述这种情况的技术术语是：JavaScript 将**抛出**一个错误。



### JavaScript try 和 catch

**try** 语句允许我们定义在执行时进行错误测试的代码块。

**catch** 语句允许我们定义当 try 代码块发生错误时，所执行的代码块。

JavaScript 语句 **try** 和 **catch** 是成对出现的。

```js
var txt=""; 
function message() 
{ 
    try { 
        adddlert("Welcome guest!"); 
    } catch(err) { 
        txt="本页有一个错误。\n\n"; 
        txt+="错误描述：" + err.message + "\n\n"; 
        txt+="点击确定继续。\n\n"; 
        alert(txt); 
    } 
}
```



### finally 语句

finally 语句不论之前的 try 和 catch 中是否产生异常都会执行该代码块。



### Throw 语句

throw 语句允许我们创建自定义错误。

正确的技术术语是：创建或**抛出异常**（exception）。

如果把 throw 与 try 和 catch 一起使用，那么您能够控制程序流，并生成自定义的错误消息。

```js
function myFunction() {
    var message, x;
    message = document.getElementById("message");
    message.innerHTML = "";
    x = document.getElementById("demo").value;
    try { 
        if(x == "")  throw "值为空";
        if(isNaN(x)) throw "不是数字";
        x = Number(x);
        if(x < 5)    throw "太小";
        if(x > 10)   throw "太大";
    }
    catch(err) {
        message.innerHTML = "错误: " + err;
    }
}
```





## JavaScript 严格模式(use strict)

JavaScript 严格模式（strict mode）即在严格的条件下运行。

#### 严格模式声明

严格模式通过在脚本或函数的头部添加 **use strict**; 表达式来声明。

```js
x = 3.14;       // 不报错
myFunction();

function myFunction() {
   "use strict";
    y = 3.14;   // 报错 (y 未定义)
}
```





## JavaScript 使用误区

------------------------------------

[JavaScript 使用误区 | 菜鸟教程 (runoob.com)](https://www.runoob.com/js/js-mistakes.html)



### 比较运算符常见错误

在常规的比较中，数据类型是被忽略的，以下 if 条件语句返回 true：

```js
var x = 10;
var y = "10";
if (x == y)   // true
if (x === y)  // false
```

这种错误经常会在 switch 语句中出现，**switch 语句会使用恒等计算符(===)进行比较**:

```js
var x = 10;
switch(x) {
    case "10": alert("Hello");  // 不会执行
}
```





## JavaScript 表单

HTML 表单验证可以通过 JavaScript 来完成。

以下实例代码用于判断表单字段(fname)值是否存在， 如果不存在，就弹出信息，阻止表单提交：

```js
function validateForm() {
    var x = document.forms["myForm"]["fname"].value;
    if (x == null || x == "") {
        alert("需要输入名字。");
        return false;
    }
}

<form name="myForm" action="demo_form.php" onsubmit="return validateForm()" method="post">
名字: <input type="text" name="fname">
<input type="submit" value="提交">
</form>
```





### 数据验证

数据验证用于确保用户输入的数据是有效的。

典型的数据验证有：

- 必需字段是否有输入?
- 用户是否输入了合法的数据?
- 在数字字段是否输入了文本?

大多数情况下，数据验证用于确保用户正确输入数据。

数据验证可以使用不同方法来定义，并通过多种方式来调用。

**服务端数据验证**是在数据提交到服务器上后再验证。

**客户端数据验证**是在数据发送到服务器前，在浏览器上完成验证。



### HTML 约束验证

HTML5 新增了 HTML 表单的验证方式：约束验证（constraint validation）。

约束验证是表单被提交时浏览器用来实现验证的一种算法。

HTML 约束验证基于：

- **HTML 输入属性**
- **CSS 伪类选择器**
- **DOM 属性和方法**





## JavaScript this 关键字

面向对象语言中 this 表示当前对象的一个引用。

但在 JavaScript 中 this 不是固定不变的，它会随着执行环境的改变而改变。

- 在方法中，this 表示该方法所属的对象。
- 如果单独使用，this 表示全局对象。
- 在函数中，this 表示全局对象。
- 在函数中，在严格模式下，this 是未定义的(undefined)。
- 在事件中，this 表示接收事件的元素。
- 类似 call() 和 apply() 方法可以将 this 引用到任何对象。



### 方法中的 this

在对象方法中， this 指向调用它所在方法的对象。

在上面一个实例中，this 表示 person 对象。

fullName 方法所属的对象就是 person。

```js
fullName : function() {
  return this.firstName + " " + this.lastName;
}
```



### 单独使用 this

单独使用 this，则它指向全局(Global)对象。

在浏览器中，window 就是该全局对象为 [**object Window**]:



### 事件中的 this

在 HTML 事件句柄中，this 指向了接收事件的 HTML 元素：

```js
<button onclick="this.style.display='none'">
点我后我就消失了
</button>
```



### 对象方法中绑定

下面实例中，this 是 person 对象，person 对象是函数的所有者：

```js
var person = {
  firstName: "John",
  lastName : "Doe",
  id       : 5566,
  fullName : function() {
    return this.firstName + " " + this.lastName;
  }
};
```

说明: **this.firstName** 表示 **this** (person) 对象的 **firstName** 属性。



### 显式函数绑定

在 JavaScript 中函数也是对象，对象则有方法，apply 和 call 就是函数对象的方法。这两个方法异常强大，他们允许切换函数执行的上下文环境（context），即 this 绑定的对象。





## JavaScript let 和 const

--------------------------

let 声明的变量只在 let 命令所在的代码块内有效。

const 声明一个只读的常量，一旦声明，常量的值就不能改变。



### JavaScript 块级作用域(Block Scope)

let 声明的变量只在 let 命令所在的代码块 **{}** 内有效，在 **{}** 之外不能访问。

```js
{ 
    let x = 2;
}
// 这里不能使用 x 变量
```

循环体中使用`let`

```js
let i = 5;
for (let i = 0; i < 10; i++) {
    // 一些代码...
}
// 这里输出 i 为 5
```



HTML 代码中使用全局变量

在 JavaScript 中, 全局作用域是针对 JavaScript 环境。

在 HTML 中, 全局作用域是针对 window 对象。

使用 **var** 关键字声明的全局作用域变量属于 window 对象:

```js
var carName = "Volvo";
// 可以使用 window.carName 访问变量
```

使用 **let** 关键字声明的全局作用域变量不属于 window 对象:

```js
let carName = "Volvo";
// 不能使用 window.carName 访问变量
```



### const 关键字

const 用于声明一个或多个常量，声明时必须进行初始化，且初始化后值不可再修改：

```js
const PI = 3.141592653589793;
PI = 3.14;      // 报错
PI = PI + 10;   // 报错
```

`const`定义常量与使用`let` 定义的变量相似：

- 二者都是块级作用域
- 都不能和它所在作用域内的其他变量或函数拥有相同的名称

两者还有以下两点区别：

- `const`声明的常量必须初始化，而`let`声明的变量不用
- const 定义常量的值不能通过再赋值修改，也不能再次声明。而 let 定义的变量值可以修改。



-----------------------



## JSON操作



### JSON 字符串转换为 JavaScript 对象

首先，创建 JavaScript 字符串，字符串为 JSON 格式的数据：

```js
var text = '{ "sites" : [' +
    '{ "name":"Runoob" , "url":"www.runoob.com" },' +
    '{ "name":"Google" , "url":"www.google.com" },' +
    '{ "name":"Taobao" , "url":"www.taobao.com" } ]}';
    
obj = JSON.parse(text);
document.getElementById("demo").innerHTML = obj.sites[1].name + " " + obj.sites[1].url;
```

**更多的JSON教程**：[JSON 教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/json/json-tutorial.html)





## javascript:void(0) 含义 —— 无返回值的函数

我们经常会使用到 **javascript:void(0)** 这样的代码，那么在 JavaScript 中 **javascript:void(0)** 代表的是什么意思呢？

**javascript:void(0)** 中最关键的是 **void** 关键字， **void** 是 JavaScript 中非常重要的关键字，该操作符指定要计算一个表达式但是不返回值。

语法格式如下：

```js
void func()
javascript:void func()
```

或

```js
void(func())
javascript:void(func())
```





### href="#"与href="javascript:void(0)"的区别

**#** 包含了一个位置信息，默认的锚是**#top** 也就是网页的上端。

而javascript:void(0), 仅仅表示一个死链接。

在页面很长的时候会使用 **#** 来定位页面的具体位置，格式为：**# + id**。

如果你要定义一个死链接请使用 javascript:void(0) 。

```js
<a href="javascript:void(0);">点我没有反应的!</a>
<a href="#pos">点我定位到指定位置!</a>
<br>
...
<br>
<p id="pos">尾部定位点</p>
```

