# TCP端口扫描器

**功能**

​	查询某个ip地址的`TCP`端口`21-120`的开启状态



### 非并发的TCP扫描器

```go
// 程序功能
// 扫描ip地址的前100个端口
// 2023年1月4日

package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 21; i < 120; i++ {
		// 服务器地址+端口号
		address := fmt.Sprintf("124.222.178.124:%d", i)
		// 返回连接和一个错误
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s 关闭了\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了!!!\n", address)
	}
}

```



### 并发的TCP扫描器

​	在`for`循环中调用`goroutine`，注意直接使用循环变量可能会存在**资源竞争**，因此最后使用传参的方式进行调用。

```go
// 程序功能
// 扫描ip地址的前100个端口
// 2023年1月4日

package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 21; i < 65535; i++ {
		wg.Add(1)
		go func(j int) { 
			defer wg.Done()
			// 服务器地址+端口号
			address := fmt.Sprintf("124.222.178.124:%d", j)
			// 返回连接和一个错误
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了!!!\n", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds", elapsed)
}
```



### 并发的TCP扫描器 - WORKER池

​	生成100个`goroutine`进行工作

<img src="C:\Users\83573\AppData\Roaming\Typora\typora-user-images\image-20230104220040965.png" alt="image-20230104220040965" style="zoom:50%;" />

```go
package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func work(ports chan int, results chan int) {
	// 等待从channel中获取数据，如果没有数据则会阻塞等待
	for p := range ports {
		address := fmt.Sprintf("124.222.178.124:%d", p)

		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			// fmt.Printf("%s 关闭了\n", address)
			results <- p * (-1)
			continue
		}
		conn.Close()
		results <- p
		// wg.Done()
	}
}

func main() {
	start := time.Now()
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var closeports []int
	// var wg sync.WaitGroup

	// 循环100次，生成100个goroutine
	for i := 0; i < cap(ports); i++ {
		go work(ports, results)
	}

	go func() {
		for i := 1; i < 65536; i++ {
			// wg.Add(1)
			// 往channel中传入数据
			ports <- i
		}
	}()

	for i := 1; i < 65536; i++ {
		port := <-results
		if port < 0 {
			openports = append(openports, -port)
		} else {
			closeports = append(closeports, port)
		}
	}

	// wg.Wait()
	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closeports)
	for _, port := range closeports {
		fmt.Printf("%d 关闭了\n", port)
	}
	for _, port := range openports {
		fmt.Printf("%d 打开了\n", port)
	}
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds", elapsed)
}
```



### TCP扫描器的命令行版本

```go
// 程序功能
// 用户通过命令行输入ip+端口号
// 程序返回该ip端口号的开启状态
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 解析命令行获取ip和端口号
	address := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])
	// 进行TCP扫描
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("%s关闭了", address)
		return
	}
	conn.Close()
	fmt.Printf("%s打开了", address)
}
```

