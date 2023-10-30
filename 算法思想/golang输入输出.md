# 读取输入

当使用Go语言编写程序时，有多种方式可以读取用户的输入。以下是一些常见的读取输入的方式：

### fmt.Scan 和 fmt.Scanln

这些函数用于从标准输入读取用户的输入，并将其存储到指定的变量中。它们以空格或换行符作为输入的分隔符。

```go
package main

import "fmt"

func main() {
    var name string
    var age int

    // 读取字符串
    fmt.Print("请输入您的姓名：")
    fmt.Scan(&name)

    // 读取整数
    fmt.Print("请输入您的年龄：")
    fmt.Scan(&age)

    fmt.Println("您的姓名是：", name)
    fmt.Println("您的年龄是：", age)
}
```

2. bufio.NewReader 和 bufio.Scanner：这些类型提供了更灵活的读取方式，可以读取一行或一段文本，并按需进行分割。

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // 使用 bufio.NewReader 读取一行文本
    fmt.Print("请输入一行文本：")
    reader := bufio.NewReader(os.Stdin)
    line, _ := reader.ReadString('\n')
    fmt.Println("您输入的文本是：", line)

    // 使用 bufio.Scanner 逐词读取文本
    fmt.Print("请输入一段文本：")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    fmt.Println("您输入的文本是：", text)
}
```

3. os.Args：这个方法通过命令行参数获取输入。os.Args 切片的第一个元素是程序本身的名称，之后的元素是用户在命令行中提供的参数。

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) > 1 {
        fmt.Println("您输入的参数是：", os.Args[1:])
    } else {
        fmt.Println("请在命令行中输入参数")
    }
}
```

这些是一些常见的读取输入的方式，您可以根据具体的需求选择适合的方法。无论使用哪种方式，都要注意对输入进行适当的验证和错误处理，以确保程序的健壮性。















# 打印输出