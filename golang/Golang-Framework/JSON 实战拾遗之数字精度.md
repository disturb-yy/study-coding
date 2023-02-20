> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.jianshu.com](https://www.jianshu.com/p/50fb2a0305a7) ![](http://upload-images.jianshu.io/upload_images/4209319-d628814663da1a2b.jpg) JSON 实战封面

正如《JSON 实战》书中所言 —— “JSON 已经成为 RESTful 接口设计中的事实标准”，作为一个写程序的，我们很难避免与 JSON 打交道。

JSON 从本质上讲就是一类字符串，所以在第二章（“在 JavaScript 中使用 JSON”）里一开始就向我们介绍了 JSON 的序列化 / 反序列化操作。用 JSON.stringify() 将信息序列化为 JSON（字符串）；用 JSON.parse() 将 JSON 反序列化为 JavaScript **可以理解**的数据结构。

例如我们在 node.js 控制台或者 chrome 控制台里输入：

```
    JSON.stringify({"数字": 12345678901234567890})


```

会得到：

```
'{"数字":12345678901234567000}'


```

再例如，输入：

```
JSON.parse('{"数字":12345678901234567890}')


```

会得到

```
{ '数字': 12345678901234567000 }


```

看看，JSON 的序列化和反序列化就是如此简单！

但是，如果我们仔细的看看序列化和反序列化的结果，就会喊出 **WTF** —— 我明明输入的数字是 **12345678901234567890**，怎么得到的却是 **12345678901234567000** ？😨

为什么会出现这样的结果？其实当你用的数字比较小的时候，是碰不上这个问题的，只有当你用的数字足够大（超过 **Number.MAX_SAFE_INTEGER**），或者足够小（小于 **Number.MIN_SAFE_INTEGER**）的时候，这个问题才会浮现，按照 [MDN](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Number/MIN_SAFE_INTEGER) 上的描述是这样的 ——

> 形成**这个数字**的原因是 JavaScript 在 IEEE 754 中使用 double-precision floating-point  
> format numbers 作为规定。在这个规定中能安全的表示数字的范围在 -(253 - 1) 到 253 - 1 之间.

怎么办?
----

是否要解决这个问题取决于你的业务场景。

如果你确信你的程序不会涉及到上面那么大的数，那就放心的使用 JSON.parse 和 JSON.stringify 好了（真是废话！😄）

### 后端之间传递

如果你只是在后端使用 JSON，那么选择方案还是挺多的，比如 node.js 里有一些库 [lossless-json](https://github.com/josdejong/lossless-json)、[json-bigint](https://github.com/sidorares/json-bigint) 可以解决精度的问题。再比如 python 就不存在这个问题（所以没人用 javascript 去做科学计算啊 😓）。

### 前后端之间传递

如果要在前端进行一些交互，比如让用户**输入数字**再转成 JSON 传递给后端，那么问题就出现了：怎么保证前端输入的数字在前端被转为 JSON 的时候不走样呢？我找了一圈之后发现答案其实挺简单，那就是你把在前端输入的数字就当成一个字符串，转成 JSON 就不会走样。

在 JS 控制台里输入：

```
JSON.stringify({"数字": "12345678901234567890"})


```

得到：

```
'{"数字":"12345678901234567890"}'


```

当这个写成字符串的 JSON 被传递到后端以后，根据后端用的语言不同，有不同的处理方法。

**例如**：
-------

### Python

python 做后端的只要直接把数字字符串转成需要的数字类型即可，例如:

```
int(json.loads('{"数字":"12345678901234567890"}')['数字'])


```

得到

```
12345678901234567890


```

### Node.js

可以利用 [decimal.js](https://github.com/MikeMcl/decimal.js/) 、[bignumber.js](https://github.com/MikeMcl/bignumber.js/!%5Bimage%5D(http://upload-images.jianshu.io/upload_images/4209319-76455863bcaf585b.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)) 这样的库把 字符串转为 "**big number**" 再进行计算：

```
   var BigNumber = require('bignumber.js');
   数字 = BigNumber('12345678901234567890');
   数字.plus('987654321')


```

### 小结

JSON 在 Javascript 的序列化与反序列化当中存在的数字精度其实是一个比较常见的问题，如果你在使用 JS 处理 JSON 数字时得到了一些莫名其妙的结果，请考虑一下是不是遇到了这个问题。