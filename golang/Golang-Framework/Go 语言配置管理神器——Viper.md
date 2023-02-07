> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/viper/)

> 李文周的 Blog viper 配置管理 go config yaml json etcd Consul toml 配置 热加载

[Viper](https://github.com/spf13/viper) 是适用于 Go 应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

[Viper](https://github.com/spf13/viper) 是适用于 Go 应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

鉴于`viper`库本身的 README 已经写得十分详细，这里就将其翻译成中文，并在最后附上两个项目中使用`viper`的示例代码以供参考。

安装
--

```
go get github.com/spf13/viper


```

什么是 Viper？
----------

Viper 是适用于 Go 应用程序（包括`Twelve-Factor App`）的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。它支持以下特性：

*   设置默认值
*   从`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式的配置文件读取配置信息
*   实时监控和重新读取配置文件（可选）
*   从环境变量中读取
*   从远程配置系统（etcd 或 Consul）读取并监控配置变化
*   从命令行参数读取配置
*   从 buffer 读取配置
*   显式配置值

为什么选择 Viper?
------------

在构建现代应用程序时，你无需担心配置文件格式；你想要专注于构建出色的软件。Viper 的出现就是为了在这方面帮助你的。

Viper 能够为你执行下列操作：

1.  查找、加载和反序列化`JSON`、`TOML`、`YAML`、`HCL`、`INI`、`envfile`和`Java properties`格式的配置文件。
2.  提供一种机制为你的不同配置选项设置默认值。
3.  提供一种机制来通过命令行参数覆盖指定选项的值。
4.  提供别名系统，以便在不破坏现有代码的情况下轻松重命名参数。
5.  当用户提供了与默认值相同的命令行或配置文件时，可以很容易地分辨出它们之间的区别。

Viper 会按照下面的优先级。每个项目的优先级都高于它下面的项目:

*   显示调用`Set`设置值
*   命令行参数（flag）
*   环境变量
*   配置文件
*   key/value 存储
*   默认值

**重要：** 目前 Viper 配置的键（Key）是大小写不敏感的。目前正在讨论是否将这一选项设为可选。

把值存入 Viper
----------

### 建立默认值

一个好的配置系统应该支持默认值。键不需要默认值，但如果没有通过配置文件、环境变量、远程配置或命令行标志（flag）设置键，则默认值非常有用。

例如：

```
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})


```

### 读取配置文件

Viper 需要最少知道在哪里查找配置文件的配置。Viper 支持`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式的配置文件。Viper 可以搜索多个路径，但目前单个 Viper 实例只支持单个配置文件。Viper 不默认任何配置搜索路径，将默认决策留给应用程序。

下面是一个如何使用 Viper 搜索和读取配置文件的示例。不需要任何特定的路径，但是至少应该提供一个配置文件预期出现的路径。

```
viper.SetConfigFile("./config.yaml") 
viper.SetConfigName("config") 
viper.SetConfigType("yaml") 
viper.AddConfigPath("/etc/appname/")   
viper.AddConfigPath("$HOME/.appname")  
viper.AddConfigPath(".")               
err := viper.ReadInConfig() 
if err != nil { 
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}


```

在加载配置文件出错时，你可以像下面这样处理找不到配置文件的特定情况：

```
if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        
    } else {
        
    }
}




```

_注意 [自 1.6 起]：_ 你也可以有不带扩展名的文件，并以编程方式指定其格式。对于位于用户`$HOME`目录中的配置文件没有任何扩展名，如`.bashrc`。

**这里补充两个问题供读者解答并自行验证**

当你使用如下方式读取配置时，viper 会从`./conf`目录下查找任何以`config`为文件名的配置文件，如果同时存在`./conf/config.json`和`./conf/config.yaml`两个配置文件的话，`viper`会从哪个配置文件加载配置呢？

```
viper.SetConfigName("config")
viper.AddConfigPath("./conf")


```

在上面两个语句下搭配使用`viper.SetConfigType("yaml")`指定配置文件类型可以实现预期的效果吗？

### 写入配置文件

从配置文件中读取配置文件是有用的，但是有时你想要存储在运行时所做的所有修改。为此，可以使用下面一组命令，每个命令都有自己的用途:

*   WriteConfig - 将当前的`viper`配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
*   SafeWriteConfig - 将当前的`viper`配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
*   WriteConfigAs - 将当前的`viper`配置写入给定的文件路径。将覆盖给定的文件 (如果它存在的话)。
*   SafeWriteConfigAs - 将当前的`viper`配置写入给定的文件路径。不会覆盖给定的文件 (如果它存在的话)。

根据经验，标记为`safe`的所有方法都不会覆盖任何文件，而是直接创建（如果不存在），而默认行为是创建或截断。

一个小示例：

```
viper.WriteConfig() 
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") 
viper.SafeWriteConfigAs("/path/to/my/.other_config")


```

### 监控并重新读取配置文件

Viper 支持在运行时实时读取配置文件的功能。

需要重新启动服务器以使配置生效的日子已经一去不复返了，viper 驱动的应用程序可以在运行时读取配置文件的更新，而不会错过任何消息。

只需告诉 viper 实例 watchConfig。可选地，你可以为 Viper 提供一个回调函数，以便在每次发生更改时运行。

**确保在调用`WatchConfig()`之前添加了所有的配置路径。**

```
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  
	fmt.Println("Config file changed:", e.Name)
})


```

### 从 io.Reader 读取配置

Viper 预先定义了许多配置源，如文件、环境变量、标志和远程 K/V 存储，但你不受其约束。你还可以实现自己所需的配置源并将其提供给 viper。

```
viper.SetConfigType("yaml") 


var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

viper.Get("name") 


```

### 覆盖设置

这些可能来自命令行标志，也可能来自你自己的应用程序逻辑。

```
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)


```

### 注册和使用别名

别名允许多个键引用单个值

```
viper.RegisterAlias("loud", "Verbose")  

viper.Set("verbose", true) 
viper.Set("loud", true)   

viper.GetBool("loud") 
viper.GetBool("verbose") 


```

### 使用环境变量

Viper 完全支持环境变量。这使`Twelve-Factor App`开箱即用。有五种方法可以帮助与 ENV 协作:

*   `AutomaticEnv()`
*   `BindEnv(string...) : error`
*   `SetEnvPrefix(string)`
*   `SetEnvKeyReplacer(string...) *strings.Replacer`
*   `AllowEmptyEnv(bool)`

_使用 ENV 变量时，务必要意识到 Viper 将 ENV 变量视为区分大小写。_

Viper 提供了一种机制来确保 ENV 变量是惟一的。通过使用`SetEnvPrefix`，你可以告诉 Viper 在读取环境变量时使用前缀。`BindEnv`和`AutomaticEnv`都将使用这个前缀。

`BindEnv`使用一个或两个参数。第一个参数是键名称，第二个是环境变量的名称。环境变量的名称区分大小写。如果没有提供 ENV 变量名，那么 Viper 将自动假设 ENV 变量与以下格式匹配：前缀 + “_” + 键名全部大写。当你显式提供 ENV 变量名（第二个参数）时，它 **不会** 自动添加前缀。例如，如果第二个参数是 “id”，Viper 将查找环境变量 “ID”。

在使用 ENV 变量时，需要注意的一件重要事情是，每次访问该值时都将读取它。Viper 在调用`BindEnv`时不固定该值。

`AutomaticEnv`是一个强大的助手，尤其是与`SetEnvPrefix`结合使用时。调用时，Viper 会在发出`viper.Get`请求时随时检查环境变量。它将应用以下规则。它将检查环境变量的名称是否与键匹配（如果设置了`EnvPrefix`）。

`SetEnvKeyReplacer`允许你使用`strings.Replacer`对象在一定程度上重写 Env 键。如果你希望在`Get()`调用中使用`-`或者其他什么符号，但是环境变量里使用`_`分隔符，那么这个功能是非常有用的。可以在`viper_test.go`中找到它的使用示例。

或者，你可以使用带有`NewWithOptions`工厂函数的`EnvKeyReplacer`。与`SetEnvKeyReplacer`不同，它接受`StringReplacer`接口，允许你编写自定义字符串替换逻辑。

默认情况下，空环境变量被认为是未设置的，并将返回到下一个配置源。若要将空环境变量视为已设置，请使用`AllowEmptyEnv`方法。

#### Env 示例：

```
SetEnvPrefix("spf") 
BindEnv("id")

os.Setenv("SPF_ID", "13") 

id := Get("id") 


```

### 使用 Flags

Viper 具有绑定到标志的能力。具体来说，Viper 支持 [Cobra](https://github.com/spf13/cobra) 库中使用的`Pflag`。

与`BindEnv`类似，该值不是在调用绑定方法时设置的，而是在访问该方法时设置的。这意味着你可以根据需要尽早进行绑定，即使在`init()`函数中也是如此。

对于单个标志，`BindPFlag()`方法提供此功能。

例如：

```
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))


```

你还可以绑定一组现有的 pflags （pflag.FlagSet）：

举个例子：

```
pflag.Int("flagname", 1234, "help message for flagname")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)

i := viper.GetInt("flagname") 


```

在 Viper 中使用 pflag 并不阻碍其他包中使用标准库中的 flag 包。pflag 包可以通过导入这些 flags 来处理 flag 包定义的 flags。这是通过调用 pflag 包提供的便利函数`AddGoFlagSet()`来实现的。

例如：

```
package main

import (
	"flag"
	"github.com/spf13/pflag"
)

func main() {

	
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") 

	...
}


```

#### flag 接口

如果你不使用`Pflag`，Viper 提供了两个 Go 接口来绑定其他 flag 系统。

`FlagValue`表示单个 flag。这是一个关于如何实现这个接口的非常简单的例子：

```
type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }


```

一旦你的 flag 实现了这个接口，你可以很方便地告诉 Viper 绑定它：

```
viper.BindFlagValue("my-flag-name", myFlag{})


```

`FlagValueSet`代表一组 flags 。这是一个关于如何实现这个接口的非常简单的例子:

```
type myFlagSet struct {
	flags []myFlag
}

func (f myFlagSet) VisitAll(fn func(FlagValue)) {
	for _, flag := range flags {
		fn(flag)
	}
}


```

一旦你的 flag set 实现了这个接口，你就可以很方便地告诉 Viper 绑定它：

```
fSet := myFlagSet{
	flags: []myFlag{myFlag{}, myFlag{}},
}
viper.BindFlagValues("my-flags", fSet)


```

### 远程 Key/Value 存储支持

在 Viper 中启用远程支持，需要在代码中匿名导入`viper/remote`这个包。

`import _ "github.com/spf13/viper/remote"`

Viper 将读取从 Key/Value 存储（例如 etcd 或 Consul）中的路径检索到的配置字符串（如`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式）。这些值的优先级高于默认值，但是会被从磁盘、flag 或环境变量检索到的配置值覆盖。（译注：也就是说 Viper 加载配置值的优先级为：磁盘上的配置文件 > 命令行标志位 > 环境变量 > 远程 Key/Value 存储 > 默认值。）

Viper 使用 [crypt](https://github.com/bketelsen/crypt) 从 K/V 存储中检索配置，这意味着如果你有正确的 gpg 密匙，你可以将配置值加密存储并自动解密。加密是可选的。

你可以将远程配置与本地配置结合使用，也可以独立使用。

`crypt`有一个命令行助手，你可以使用它将配置放入 K/V 存储中。`crypt`默认使用在 [http://127.0.0.1:4001](http://127.0.0.1:4001/) 的 etcd。

```
$ go get github.com/bketelsen/crypt/bin/crypt
$ crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json


```

确认值已经设置：

```
$ crypt get -plaintext /config/hugo.json


```

有关如何设置加密值或如何使用 Consul 的示例，请参见`crypt`文档。

### 远程 Key/Value 存储示例 - 未加密

#### etcd

```
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") 
err := viper.ReadRemoteConfig()


```

#### Consul

你需要 Consul Key/Value 存储中设置一个 Key 保存包含所需配置的 JSON 值。例如，创建一个 key`MY_CONSUL_KEY`将下面的值存入 Consul key/value 存储：

```
{
    "port": 8080,
    "hostname": "liwenzhou.com"
}


```

```
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json") 
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port")) 
fmt.Println(viper.Get("hostname")) 


```

#### Firestore

```
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json") 
err := viper.ReadRemoteConfig()


```

当然，你也可以使用`SecureRemoteProvider`。

### 远程 Key/Value 存储示例 - 加密

```
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") 
err := viper.ReadRemoteConfig()


```

### 监控 etcd 中的更改 - 未加密

```
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") 


err := runtime_viper.ReadRemoteConfig()


runtime_viper.Unmarshal(&runtime_conf)


go func(){
	for {
	    time.Sleep(time.Second * 5) 

	    
	    err := runtime_viper.WatchRemoteConfig()
	    if err != nil {
	        log.Errorf("unable to read remote config: %v", err)
	        continue
	    }

	    
	    runtime_viper.Unmarshal(&runtime_conf)
	}
}()


```

从 Viper 获取值
-----------

在 Viper 中，有几种方法可以根据值的类型获取值。存在以下功能和方法:

*   `Get(key string) : interface{}`
*   `GetBool(key string) : bool`
*   `GetFloat64(key string) : float64`
*   `GetInt(key string) : int`
*   `GetIntSlice(key string) : []int`
*   `GetString(key string) : string`
*   `GetStringMap(key string) : map[string]interface{}`
*   `GetStringMapString(key string) : map[string]string`
*   `GetStringSlice(key string) : []string`
*   `GetTime(key string) : time.Time`
*   `GetDuration(key string) : time.Duration`
*   `IsSet(key string) : bool`
*   `AllSettings() : map[string]interface{}`

需要认识到的一件重要事情是，每一个 Get 方法在找不到值的时候都会返回零值。为了检查给定的键是否存在，提供了`IsSet()`方法。

例如：

```
viper.GetString("logfile") 
if viper.GetBool("verbose") {
    fmt.Println("verbose enabled")
}


```

### 访问嵌套的键

访问器方法也接受深度嵌套键的格式化路径。例如，如果加载下面的 JSON 文件：

```
{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}


```

Viper 可以通过传入`.`分隔的路径来访问嵌套字段：

```
GetString("datastore.metric.host") 


```

这遵守上面建立的优先规则；搜索路径将遍历其余配置注册表，直到找到为止。(译注：因为 Viper 支持从多种配置来源，例如磁盘上的配置文件> 命令行标志位 > 环境变量 > 远程 Key/Value 存储 > 默认值，我们在查找一个配置的时候如果在当前配置源中没找到，就会继续从后续的配置源查找，直到找到为止。)

例如，在给定此配置文件的情况下，`datastore.metric.host`和`datastore.metric.port`均已定义（并且可以被覆盖）。如果另外在默认值中定义了`datastore.metric.protocol`，Viper 也会找到它。

然而，如果`datastore.metric`被直接赋值覆盖（被 flag，环境变量，`set()`方法等等…），那么`datastore.metric`的所有子键都将变为未定义状态，它们被高优先级配置级别 “遮蔽”（shadowed）了。

最后，如果存在与分隔的键路径匹配的键，则返回其值。例如：

```
{
    "datastore.metric.host": "0.0.0.0",
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

GetString("datastore.metric.host") 


```

### 提取子树

从 Viper 中提取子树。

例如，`viper`实例现在代表了以下配置：

```
app:
  cache1:
    max-items: 100
    item-size: 64
  cache2:
    max-items: 200
    item-size: 80


```

执行后：

```
subv := viper.Sub("app.cache1")


```

`subv`现在就代表：

```
max-items: 100
item-size: 64


```

假设我们现在有这么一个函数：

```
func NewCache(cfg *Viper) *Cache {...}


```

它基于`subv`格式的配置信息创建缓存。现在，可以轻松地分别创建这两个缓存，如下所示：

```
cfg1 := viper.Sub("app.cache1")
cache1 := NewCache(cfg1)

cfg2 := viper.Sub("app.cache2")
cache2 := NewCache(cfg2)


```

### 反序列化

你还可以选择将所有或特定的值解析到结构体、map 等。

有两种方法可以做到这一点：

*   `Unmarshal(rawVal interface{}) : error`
*   `UnmarshalKey(key string, rawVal interface{}) : error`

举个例子：

```
type config struct {
	Port int
	Name string
	PathMap string `mapstructure:"path_map"`
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}


```

如果你想要解析那些键本身就包含`.`(默认的键分隔符）的配置，你需要修改分隔符：

```
v := viper.NewWithOptions(viper.KeyDelimiter("::"))

v.SetDefault("chart::values", map[string]interface{}{
    "ingress": map[string]interface{}{
        "annotations": map[string]interface{}{
            "traefik.frontend.rule.type":                 "PathPrefix",
            "traefik.ingress.kubernetes.io/ssl-redirect": "true",
        },
    },
})

type config struct {
	Chart struct{
        Values map[string]interface{}
    }
}

var C config

v.Unmarshal(&C)


```

Viper 还支持解析到嵌入的结构体：

```
type config struct {
	Module struct {
		Enabled bool

		moduleConfig `mapstructure:",squash"`
	}
}


type moduleConfig struct {
	Token string
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}


```

Viper 在后台使用 [github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) 来解析值，其默认情况下使用`mapstructure`tag。

**注意** 当我们需要将 viper 读取的配置反序列到我们定义的结构体变量中时，一定要使用`mapstructure`tag 哦！

### 序列化成字符串

你可能需要将 viper 中保存的所有设置序列化到一个字符串中，而不是将它们写入到一个文件中。你可以将自己喜欢的格式的序列化器与`AllSettings()`返回的配置一起使用。

```
import (
    yaml "gopkg.in/yaml.v2"
    
)

func yamlStringSettings() string {
    c := viper.AllSettings()
    bs, err := yaml.Marshal(c)
    if err != nil {
        log.Fatalf("unable to marshal config to YAML: %v", err)
    }
    return string(bs)
}


```

使用单个还是多个 Viper 实例?
------------------

Viper 是开箱即用的。你不需要配置或初始化即可开始使用 Viper。由于大多数应用程序都希望使用单个中央存储库管理它们的配置信息，所以 viper 包提供了这个功能。它类似于单例模式。

在上面的所有示例中，它们都以其单例风格的方法演示了如何使用 viper。

### 使用多个 viper 实例

你还可以在应用程序中创建许多不同的 viper 实例。每个都有自己独特的一组配置和值。每个人都可以从不同的配置文件，key value 存储区等读取数据。每个都可以从不同的配置文件、键值存储等中读取。viper 包支持的所有功能都被镜像为 viper 实例的方法。

例如：

```
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")




```

当使用多个 viper 实例时，由用户来管理不同的 viper 实例。

使用 Viper 示例
-----------

假设我们的项目现在有一个`./conf/config.yaml`配置文件，内容如下：

```
port: 8123
version: "v1.2.3"


```

接下来通过示例代码演示两种在项目中使用`viper`管理项目配置信息的方式。

### 直接使用 viper 管理配置

这里用一个 demo 演示如何在 gin 框架搭建的 web 项目中使用`viper`，使用 viper 加载配置文件中的信息，并在代码中直接使用`viper.GetXXX()`方法获取对应的配置值。

```
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./conf/config.yaml") 
	err := viper.ReadInConfig()        
	if err != nil {                    
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	
	viper.WatchConfig()

	r := gin.Default()
	
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})

	if err := r.Run(
		fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		panic(err)
	}
}


```

### 使用结构体变量保存配置信息

除了上面的用法外，我们还可以在项目中定义与配置文件对应的结构体，`viper`加载完配置信息后使用结构体变量保存配置信息。

```
package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

var Conf = new(Config)

func main() {
	viper.SetConfigFile("./conf/config.yaml") 
	err := viper.ReadInConfig()               
	if err != nil {                           
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	
	viper.WatchConfig()
	
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	r := gin.Default()
	
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, Conf.Version)
	})

	if err := r.Run(fmt.Sprintf(":%d", Conf.Port)); err != nil {
		panic(err)
	}
}


```

参考链接
----

[https://github.com/spf13/viper/blob/master/README.md](https://github.com/spf13/viper/blob/master/README.md)

![](https://www.liwenzhou.com/images/wxgzh_qrcode.png)