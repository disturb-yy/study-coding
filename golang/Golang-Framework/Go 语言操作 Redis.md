> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/redis/)

> 李文周的 Blog 中本篇文章介绍了 Redis 的常用数据类型，以及如何使用 Go 语言的 go-redis 库连接 redis（集群、哨兵等），执行 redis 基本命令、pipeline、事务和 watch 的用法。

在项目开发中 redis 的使用也比较频繁，本文介绍了 Go 语言中`go-redis`库的基本使用。

在项目开发中 redis 的使用也比较频繁，本文介绍了 Go 语言中`go-redis`库的基本使用。

Redis 介绍
--------

Redis 是一个开源的内存数据库，Redis 提供了多种不同类型的数据结构，很多业务场景下的问题都可以很自然地映射到这些数据结构上。除此之外，通过复制、持久化和客户端分片等特性，我们可以很方便地将 Redis 扩展成一个能够包含数百 GB 数据、每秒处理上百万次请求的系统。

### Redis 支持的数据结构

Redis 支持诸如字符串（string）、哈希（hashe）、列表（list）、集合（set）、带范围查询的排序集合（sorted set）、bitmap、hyperloglog、带半径查询的地理空间索引（geospatial index）和流（stream）等数据结构。

### Redis 应用场景

*   缓存系统，减轻主数据库（MySQL）的压力。
*   计数场景，比如微博、抖音中的关注数和粉丝数。
*   热门排行榜，需要排序的场景特别适合使用 ZSET。
*   利用 LIST 可以实现队列的功能。
*   利用 HyperLogLog 统计 UV、PV 等数据。
*   使用 geospatial index 进行地理位置相关查询。

### 准备 Redis 环境

读者可以选择在本机安装 redis 或使用云数据库，这里直接使用 Docker 启动一个 redis 环境，方便学习使用。

使用下面的命令启动一个名为 redis507 的 5.0.7 版本的 redis server 环境。

```
docker run --name redis507 -p 6379:6379 -d redis:5.0.7


```

**注意：**此处的版本、容器名和端口号可以根据自己需要设置。

启动一个 redis-cli 连接上面的 redis server。

```
docker run -it --network host --rm redis:5.0.7 redis-cli
```

go-redis 库
----------

### 安装

Go 社区中目前有很多成熟的 redis client 库，比如 [[https://github.com/gomodule/redigo](https://github.com/gomodule/redigo) 和 [https://github.com/go-redis/redis](https://github.com/go-redis/redis)，读者可以自行选择适合自己的库。本书使用 go-redis 这个库来操作 Redis 数据库。

使用以下命令下安装 go-redis 库。

```
go get github.com/go-redis/redis/v8
```

### 连接

#### 普通连接模式

go-redis 库中使用 redis.NewClient 函数连接 Redis 服务器。

```
rdb := redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", 
	DB:       0,  
	PoolSize: 20, 
})


```

除此之外，还可以使用 redis.ParseURL 函数从表示数据源的字符串中解析得到 Redis 服务器的配置信息。

```
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}

rdb := redis.NewClient(opt)
```

#### TLS 连接模式

如果使用的是 TLS 连接方式，则需要使用 tls.Config 配置。

```
rdb := redis.NewClient(&redis.Options{
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
	},
})


```

#### Redis Sentinel 模式

使用下面的命令连接到由 Redis Sentinel 管理的 Redis 服务器。

```
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "master-name",
    SentinelAddrs: []string{":9126", ":9127", ":9128"},
})


```

#### Redis Cluster 模式

使用下面的命令连接到 Redis Cluster，go-redis 支持按延迟或随机路由命令。

```
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"}, 
})
```

基本使用
----

### 执行命令

下面的示例代码演示了 go-redis 库的基本使用。

```
func doCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val, err)

	
	cmder := rdb.Get(ctx, "key")
	fmt.Println(cmder.Val()) 
	fmt.Println(cmder.Err()) 

	
	err = rdb.Set(ctx, "key", 10, time.Hour).Err()

	
	value := rdb.Get(ctx, "key").Val()
	fmt.Println(value)
}


```

### 执行任意命令

go-redis 还提供了一个执行任意命令或自定义命令的 Do 方法，特别是一些 go-redis 库暂时不支持的命令都可以使用该方法执行。具体使用方法如下。

```
func doDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	
	err := rdb.Do(ctx, "set", "key", 10, "EX", 3600).Err()
	fmt.Println(err)

	
	val, err := rdb.Do(ctx, "get", "key").Result()
	fmt.Println(val, err)
}


```

### redis.Nil

go-redis 库提供了一个 redis.Nil 错误来表示 Key 不存在的错误。因此在使用 go-redis 时需要注意对返回错误的判断。在某些场景下我们应该区别处理 redis.Nil 和其他不为 nil 的错误。

```
func getValueFromRedis(key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		
		return "", err
	}
	return val, nil
}


```

其他示例
----

### zset 示例

下面的示例代码演示了如何使用 go-redis 库操作 zset。

```
func zsetDemo() {
	
	zsetKey := "language_rank"
	
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}


```

执行上面的函数将得到如下输出结果。

```
zadd success
Golang's score is 100.000000 now.
Golang 100
C/C++ 99
Java 98
Python 95
JavaScript 97
Java 98
C/C++ 99
Golang 100


```

### 扫描或遍历所有 key

你可以使用`KEYS prefix:*` 命令按前缀获取所有 key。

```
vals, err := rdb.Keys(ctx, "prefix*").Result()


```

但是如果需要扫描数百万的 key ，那速度就会比较慢。这种场景下你可以使用 Scan 命令来遍历所有符合要求的 key。

```
func scanKeysDemo1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		
		keys, cursor, err = rdb.Scan(ctx, cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { 
			break
		}
	}
}


```

Go-redis 允许将上面的代码简化为如下示例。

```
func scanKeysDemo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}


```

例如，我们可以写出一个将所有匹配指定模式的 key 删除的示例。

```
func delKeysByMatch(match string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	iter := rdb.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}


```

此外，对于 Redis 中的 set、hash、zset 数据类型，go-redis 也支持类似的遍历方法。

```
iter := rdb.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
iter := rdb.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
iter := rdb.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator(


```

Pipeline
--------

Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。区别于一个接一个地执行 100 个命令，你可以将这些命令放入 pipeline 中，然后使用 1 次读写操作像执行单个命令一样执行它们。这样做的好处是节省了执行命令的网络往返时间（RTT）。

y 在下面的示例代码中演示了使用 pipeline 通过一个 write + read 操作来执行多个命令。

```
pipe := rdb.Pipeline()

incr := pipe.Incr(ctx, "pipeline_counter")
pipe.Expire(ctx, "pipeline_counter", time.Hour)

cmds, err := pipe.Exec(ctx)
if err != nil {
	panic(err)
}


fmt.Println(incr.Val())


```

上面的代码相当于将以下两个命令一次发给 Redis Server 端执行，与不使用 Pipeline 相比能减少一次 RTT。

```
INCR pipeline_counter
EXPIRE pipeline_counts 3600


```

或者，你也可以使用`Pipelined` 方法，它会在函数退出时调用 Exec。

```
var incr *redis.IntCmd

cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
	incr = pipe.Incr(ctx, "pipelined_counter")
	pipe.Expire(ctx, "pipelined_counter", time.Hour)
	return nil
})
if err != nil {
	panic(err)
}


fmt.Println(incr.Val())


```

我们可以遍历 pipeline 命令的返回值依次获取每个命令的结果。下方的示例代码中使用 pipiline 一次执行了 100 个 Get 命令，在 pipeline 执行后遍历取出 100 个命令的执行结果。

```
cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
	for i := 0; i < 100; i++ {
		pipe.Get(ctx, fmt.Sprintf("key%d", i))
	}
	return nil
})
if err != nil {
	panic(err)
}

for _, cmd := range cmds {
    fmt.Println(cmd.(*redis.StringCmd).Val())
}


```

在那些我们需要一次性执行多个命令的场景下，就可以考虑使用 pipeline 来优化。

事务
--

Redis 是单线程执行命令的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，例如在它们之间交替执行。但是，`Multi/exec`能够确保在`multi/exec`两个语句之间的命令之间没有其他客户端正在执行命令。

在这种场景我们需要使用 TxPipeline 或 TxPipelined 方法将 pipeline 命令使用 `MULTI` 和`EXEC`包裹起来。

```
pipe := rdb.TxPipeline()
incr := pipe.Incr(ctx, "tx_pipeline_counter")
pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
_, err := pipe.Exec(ctx)
fmt.Println(incr.Val(), err)


var incr2 *redis.IntCmd
_, err = rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
	incr2 = pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	return nil
})
fmt.Println(incr2.Val(), err)


```

上面代码相当于在一个 RTT 下执行了下面的 redis 命令：

```
MULTI
INCR pipeline_counter
EXPIRE pipeline_counts 3600
EXEC


```

### Watch

我们通常搭配 `WATCH`命令来执行事务操作。从使用`WATCH`命令监视某个 key 开始，直到执行`EXEC`命令的这段时间里，如果有其他用户抢先对被监视的 key 进行了替换、更新、删除等操作，那么当用户尝试执行`EXEC`的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。

Watch 方法接收一个函数和一个或多个 key 作为参数。

```
Watch(fn func(*Tx) error, keys ...string) error


```

下面的代码片段演示了 Watch 方法搭配 TxPipelined 的使用示例。

```
func watchDemo(ctx context.Context, key string) error {
	return rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		
		
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
}


```

将上面的函数执行并打印其返回值，如果我们在程序运行后的 5 秒内修改了被 watch 的 key 的值，那么该事务操作失败，返回`redis: transaction failed`错误。

最后我们来看一个 go-redis 官方文档中使用 `GET` 、`SET`和`WATCH`命令实现一个 INCR 命令的完整示例。

```
const routineCount = 100

increment := func(key string) error {
	txf := func(tx *redis.Tx) error {
		
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		
		n++

		
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			
			pipe.Set(key, n, 0)
			return nil
		})
		return err
	}

	for retries := routineCount; retries > 0; retries-- {
		err := rdb.Watch(txf, key)
		if err != redis.TxFailedErr {
			return err
		}
		
	}
	return errors.New("increment reached maximum number of retries")
}

var wg sync.WaitGroup
wg.Add(routineCount)
for i := 0; i < routineCount; i++ {
	go func() {
		defer wg.Done()

		if err := increment("counter3"); err != nil {
			fmt.Println("increment error:", err)
		}
	}()
}
wg.Wait()

n, err := rdb.Get("counter3").Int()
fmt.Println("ended with", n, err)


```

在这个示例中使用了 `redis.TxFailedErr` 来检查事务是否失败。

更多详情请查阅[文档](https://pkg.go.dev/github.com/go-redis/redis)。

![](https://www.liwenzhou.com/images/wxgzh_qrcode.png)