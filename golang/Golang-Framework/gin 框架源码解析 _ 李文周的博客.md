> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/gin-sourcecode/)

> 李文周的 Blog gin 框架 go web 源码 路由 中间件 分析 trie radix tree 对象池 sync.Pool c.Next c.Abort

通过阅读 gin 框架的源码来探究 gin 框架路由与中间件的秘密。

gin 框架路由详解
----------

gin 框架使用的是定制版本的 [httprouter](https://github.com/julienschmidt/httprouter)，其路由的原理是大量使用公共前缀的树结构，它基本上是一个紧凑的 [Trie tree](https://baike.sogou.com/v66237892.htm)（或者只是 [Radix Tree](https://baike.sogou.com/v73626121.htm)）。具有公共前缀的节点也共享一个公共父节点。

### Radix Tree

基数树（Radix Tree）又称为 PAT 位树（Patricia Trie or crit bit tree），是一种更节省空间的前缀树（Trie Tree）。对于基数树的每个节点，如果该节点是唯一的子树的话，就和父节点合并。下图为一个基数树示例：

![](https://www.liwenzhou.com/images/Go/gin/radix_tree.png)

`Radix Tree`可以被认为是一棵简洁版的前缀树。我们注册路由的过程就是构造前缀树的过程，具有公共前缀的节点也共享一个公共父节点。假设我们现在注册有以下路由信息：

```
r := gin.Default()

r.GET("/", func1)
r.GET("/search/", func2)
r.GET("/support/", func3)
r.GET("/blog/", func4)
r.GET("/blog/:post/", func5)
r.GET("/about-us/", func6)
r.GET("/about-us/team/", func7)
r.GET("/contact/", func8)


```

那么我们会得到一个`GET`方法对应的路由树，具体结构如下：

```
Priority   Path             Handle
9          \                *<1>
3          ├s               nil
2          |├earch\         *<2>
1          |└upport\        *<3>
2          ├blog\           *<4>
1          |    └:post      nil
1          |         └\     *<5>
2          ├about-us\       *<6>
1          |        └team\  *<7>
1          └contact\        *<8>


```

上面最右边那一列每个`*<数字>`表示 Handle 处理函数的内存地址 (一个指针)。从根节点遍历到叶子节点我们就能得到完整的路由表。

例如：`blog/:post`其中`:post`只是实际文章名称的占位符 (参数)。与`hash-maps`不同，这种树结构还允许我们使用像`:post`参数这种动态部分，因为我们实际上是根据路由模式进行匹配，而不仅仅是比较哈希值。

由于 URL 路径具有层次结构，并且只使用有限的一组字符 (字节值)，所以很可能有许多常见的前缀。这使我们可以很容易地将路由简化为更小的问题。此外，**路由器为每种请求方法管理一棵单独的树**。一方面，它比在每个节点中都保存一个 method-> handle map 更加节省空间，它还使我们甚至可以在开始在前缀树中查找之前大大减少路由问题。

为了获得更好的可伸缩性，**每个树级别上的子节点都按`Priority(优先级)`排序，其中优先级（最左列）就是在子节点 (子节点、子子节点等等) 中注册的句柄的数量。**==即子节点多的路径优先匹配==，这样做有两个好处:

1.  首先优先匹配被大多数路由路径包含的节点。这样可以让尽可能多的路由快速被定位。
    
2.  类似于成本补偿。最长的路径可以被优先匹配，补偿体现在最长的路径需要花费更长的时间来定位，如果最长路径的节点能被优先匹配（即每次拿子节点都命中），那么路由匹配所花的时间不一定比短路径的路由长。下面展示了节点（每个`-`可以看做一个节点）匹配的路径：从左到右，从上到下。
    

```
   ├------------
   ├---------
   ├-----
   ├----
   ├--
   ├--
   └-


```

### 路由树节点

路由树是由一个个节点构成的，gin 框架路由树的节点由`node`结构体表示，它有以下字段：

```
type node struct {
   
	path      string
	
	
	
	indices   string
	
	children  []*node
	
	handlers  HandlersChain
	
	priority  uint32
	
	
	
	
	
	nType     nodeType
	
	maxParams uint8
	
	wildChild bool
	
	fullPath  string
}


```

### 请求方法树

在 gin 的路由中，每一个`HTTP Method`(GET、POST、PUT、DELETE…) 都对应了一棵 `radix tree`，我们注册路由的时候会调用下面的`addRoute`函数：

```
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
   
   
   
	root := engine.trees.get(method)
	if root == nil {
	
	   
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)
}


```

从上面的代码中我们可以看到在注册路由的时候都是先根据请求方法获取对应的树，也就是 gin 框架会为每一个请求方法创建一棵对应的树。只不过需要注意到一个细节是 gin 框架中保存请求方法对应树关系并不是使用的 map 而是使用的切片，`engine.trees`的类型是`methodTrees`，其定义如下：

```
type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree  


```

而获取请求方法对应树的 get 方法定义如下：

```
func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}


```

为什么使用切片而不是 map 来存储`请求方法->树`的结构呢？我猜是出于节省内存的考虑吧，毕竟 HTTP 请求方法的数量是固定的，而且常用的就那几种，所以即使使用切片存储查询起来效率也足够了。顺着这个思路，我们可以看一下 gin 框架中`engine`的初始化方法中，确实对`tress`字段做了一次内存申请：

```
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		
		
		trees:                  make(methodTrees, 0, 9),
		
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}


```

### 注册路由

注册路由的逻辑主要有`addRoute`函数和`insertChild`方法。

#### addRoute

```
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++
	numParams := countParams(path)  

	
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(numParams, path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk:
	for {
		
		if numParams > n.maxParams {
			n.maxParams = numParams
		}

		
		
		
		i := longestCommonPrefix(path, n.path)

		
		
		
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],  
				wildChild: n.wildChild,
				indices:   n.indices,
				children:  n.children,
				handlers:  n.handlers,
				priority:  n.priority - 1, 
				fullPath:  n.fullPath,
			}

			
			for _, v := range child.children {
				if v.maxParams > child.maxParams {
					child.maxParams = v.maxParams
				}
			}

			n.children = []*node{&child}
			
			n.indices = string([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

		
		if i < len(path) {
			path = path[i:]

			if n.wildChild {  
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++

				
				if numParams > n.maxParams {
					n.maxParams = numParams
				}
				numParams--

				
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
					
					if len(n.path) >= len(path) || path[len(n.path)] == '/' {
						continue walk
					}
				}

				pathSeg := path
				if n.nType != catchAll {
					pathSeg = strings.SplitN(path, "/", 2)[0]
				}
				prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
				panic("'" + pathSeg +
					"' in new path '" + fullPath +
					"' conflicts with existing wildcard '" + n.path +
					"' in existing prefix '" + prefix +
					"'")
			}
			
			c := path[0]

			
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

			
			
			
			for i, max := 0, len(n.indices); i < max; i++ {
				if c == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}

			
			if c != ':' && c != '*' {
				
				
				n.indices += string([]byte{c})
				child := &node{
					maxParams: numParams,
					fullPath:  fullPath,
				}
				
				n.children = append(n.children, child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
			}
			n.insertChild(numParams, path, fullPath, handlers)
			return
		}

		
		if n.handlers != nil {
			panic("handlers are already registered for path '" + fullPath + "'")
		}
		n.handlers = handlers
		return
	}
}


```

其实上面的代码很好理解，大家可以参照动画尝试将以下情形代入上面的代码逻辑，体味整个路由树构造的详细过程：

1.  第一次注册路由，例如注册 search
2.  继续注册一条没有公共前缀的路由，例如 blog
3.  注册一条与先前注册的路由有公共前缀的路由，例如 support

![](https://www.liwenzhou.com/images/Go/gin/addroute.gif)

#### insertChild

```
func (n *node) insertChild(numParams uint8, path string, fullPath string, handlers HandlersChain) {
  
	for numParams > 0 {
		
		wildcard, i, valid := findWildcard(path)
		if i < 0 { 
			break
		}

		
		if !valid {
			panic("only one wildcard per path segment is allowed, has: '" +
				wildcard + "' in path '" + fullPath + "'")
		}

		
		if len(wildcard) < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		
		
		if len(n.children) > 0 {
			panic("wildcard segment '" + wildcard +
				"' conflicts with existing children in path '" + fullPath + "'")
		}

		if wildcard[0] == ':' { 
			if i > 0 {
				
				n.path = path[:i]
				path = path[i:]
			}

			n.wildChild = true
			child := &node{
				nType:     param,
				path:      wildcard,
				maxParams: numParams,
				fullPath:  fullPath,
			}
			n.children = []*node{child}
			n = child
			n.priority++
			numParams--

			
			
			if len(wildcard) < len(path) {
				path = path[len(wildcard):]

				child := &node{
					maxParams: numParams,
					priority:  1,
					fullPath:  fullPath,
				}
				n.children = []*node{child}
				n = child  
				continue
			}

			
			n.handlers = handlers
			return
		}

		
		if i+len(wildcard) != len(path) || numParams > 1 {
			panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
		}

		if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
			panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
		}

		
		i--
		if path[i] != '/' {
			panic("no / before catch-all in path '" + fullPath + "'")
		}

		n.path = path[:i]
		
		
		child := &node{
			wildChild: true,
			nType:     catchAll,
			maxParams: 1,
			fullPath:  fullPath,
		}
		
		if n.maxParams < 1 {
			n.maxParams = 1
		}
		n.children = []*node{child}
		n.indices = string('/')
		n = child
		n.priority++

		
		child = &node{
			path:      path[i:],
			nType:     catchAll,
			maxParams: 1,
			handlers:  handlers,
			priority:  1,
			fullPath:  fullPath,
		}
		n.children = []*node{child}

		return
	}

	
	n.path = path
	n.handlers = handlers
	n.fullPath = fullPath
}


```

`insertChild`函数是根据`path`本身进行分割，将`/`分开的部分分别作为节点保存，形成一棵树结构。参数匹配中的`:`和`*`的区别是，前者是匹配一个字段而后者是匹配后面所有的路径。

### 路由匹配

我们先来看 gin 框架处理请求的入口函数`ServeHTTP`：

```
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  
	c := engine.pool.Get().(*Context)
  
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)  

	engine.pool.Put(c)  
}


```

函数很长，这里省略了部分代码，只保留相关逻辑代码：

```
func (engine *Engine) handleHTTPRequest(c *Context) {
	

	
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		
		value := root.getValue(rPath, c.Params, unescape)
		if value.handlers != nil {
			c.handlers = value.handlers
			c.Params = value.params
			c.fullPath = value.fullPath
			c.Next()  
			c.writermem.WriteHeaderNow()
			return
		}
	
	
	c.handlers = engine.allNoRoute
	serveError(c, http.StatusNotFound, default404Body)
}


```

路由匹配是由节点的 `getValue`方法实现的。`getValue`根据给定的路径 (键) 返回`nodeValue`值，保存注册的处理函数和匹配到的路径参数数据。

如果找不到任何处理函数，则会尝试 TSR(尾随斜杠重定向)。

代码虽然很长，但还算比较工整。大家可以借助注释看一下路由查找及参数匹配的逻辑。

```
type nodeValue struct {
	handlers HandlersChain
	params   Params  
	tsr      bool
	fullPath string
}



func (n *node) getValue(path string, po Params, unescape bool) (value nodeValue) {
	value.params = po
walk: 
	for {
		prefix := n.path
		if path == prefix {
			
			
			if value.handlers = n.handlers; value.handlers != nil {
				value.fullPath = n.fullPath
				return
			}

			if path == "/" && n.wildChild && n.nType != root {
				value.tsr = true
				return
			}

			
			indices := n.indices
			for i, max := 0, len(indices); i < max; i++ {
				if indices[i] == '/' {
					n = n.children[i]
					value.tsr = (len(n.path) == 1 && n.handlers != nil) ||
						(n.nType == catchAll && n.children[0].handlers != nil)
					return
				}
			}

			return
		}

		if len(path) > len(prefix) && path[:len(prefix)] == prefix {
			path = path[len(prefix):]
			
			
			if !n.wildChild {
				c := path[0]
				indices := n.indices
				for i, max := 0, len(indices); i < max; i++ {
					if c == indices[i] {
						n = n.children[i] 
						continue walk
					}
				}

				
				
				
				value.tsr = path == "/" && n.handlers != nil
				return
			}

			
			n = n.children[0]
			switch n.nType {
			case param:
				
				end := 0
				for end < len(path) && path[end] != '/' {
					end++
				}

				
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] 
				value.params[i].Key = n.path[1:]
				val := path[:end]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(val); err != nil {
						value.params[i].Value = val 
					}
				} else {
					value.params[i].Value = val
				}

				
				if end < len(path) {
					if len(n.children) > 0 {
						path = path[end:]
						n = n.children[0]
						continue walk
					}

					
					value.tsr = len(path) == end+1
					return
				}

				if value.handlers = n.handlers; value.handlers != nil {
					value.fullPath = n.fullPath
					return
				}
				if len(n.children) == 1 {
					
					
					n = n.children[0]
					value.tsr = n.path == "/" && n.handlers != nil
				}
				return

			case catchAll:
				
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] 
				value.params[i].Key = n.path[2:]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(path); err != nil {
						value.params[i].Value = path 
					}
				} else {
					value.params[i].Value = path
				}

				value.handlers = n.handlers
				value.fullPath = n.fullPath
				return

			default:
				panic("invalid node type")
			}
		}

		
		
		value.tsr = (path == "/") ||
			(len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
				path == prefix[:len(prefix)-1] && n.handlers != nil)
		return
	}
}


```

gin 框架中间件详解
-----------

gin 框架涉及中间件相关有 4 个常用的方法，它们分别是`c.Next()`、`c.Abort()`、`c.Set()`、`c.Get()`。

### 中间件的注册

gin 框架中的中间件设计很巧妙，我们可以首先从我们最常用的`r := gin.Default()`的`Default`函数开始看，它内部构造一个新的`engine`之后就通过`Use()`函数注册了`Logger`中间件和`Recovery`中间件：

```
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())  
	return engine
}


```

继续往下查看一下`Use()`函数的代码：

```
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)  
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}


```

从下方的代码可以看出，注册中间件其实就是将中间件函数追加到`group.Handlers`中：

```
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, middleware...)
	return group.returnObj()
}


```

而我们注册路由时会将对应路由的函数和之前的中间件函数结合到一起：

```
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)  
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}


```

其中结合操作的函数内容如下，注意观察这里是如何实现拼接两个切片得到一个新切片的。

```
const abortIndex int8 = math.MaxInt8 / 2

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {  
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}


```

也就是说，我们会将一个路由的中间件函数和处理函数结合到一起组成一条处理函数链条`HandlersChain`，而它本质上就是一个由`HandlerFunc`组成的切片：

```
type HandlersChain []HandlerFunc


```

### 中间件的执行

我们在上面路由匹配的时候见过如下逻辑：

```
value := root.getValue(rPath, c.Params, unescape)
if value.handlers != nil {
  c.handlers = value.handlers
  c.Params = value.params
  c.fullPath = value.fullPath
  c.Next()  
  c.writermem.WriteHeaderNow()
  return
}


```

其中`c.Next()`就是很关键的一步，它的代码很简单：

```
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}


```

从上面的代码可以看到，这里通过索引遍历`HandlersChain`链条，从而实现依次调用该路由的每一个函数（中间件或处理请求的函数）。

![](https://www.liwenzhou.com/images/Go/gin/gin_middleware1.png)

我们可以在中间件函数中通过再次调用`c.Next()`实现嵌套调用（func1 中调用 func2；func2 中调用 func3），

![](https://www.liwenzhou.com/images/Go/gin/gin_middleware2.png)

或者通过调用`c.Abort()`中断整个调用链条，从当前函数返回。

```
func (c *Context) Abort() {
	c.index = abortIndex  
}


```

### c.Set()/c.Get()

`c.Set()`和`c.Get()`这两个方法多用于在多个函数之间通过`c`传递数据的，比如我们可以在认证中间件中获取当前请求的相关信息（userID 等）通过`c.Set()`存入`c`，然后在后续处理业务逻辑的函数中通过`c.Get()`来获取当前请求的用户。`c`就像是一根绳子，将该次请求相关的所有的函数都串起来了。 ![](https://www.liwenzhou.com/images/Go/gin/gin_middleware3.png)

总结
--

1.  gin 框架路由使用前缀树，路由注册的过程是构造前缀树的过程，路由匹配的过程就是查找前缀树的过程。
2.  gin 框架的中间件函数和处理函数是以切片形式的调用链条存在的，我们可以顺序调用也可以借助`c.Next()`方法实现嵌套调用。
3.  借助`c.Set()`和`c.Get()`方法我们能够在不同的中间件函数中传递数据。

![](https://www.liwenzhou.com/images/wxgzh_qrcode.png)