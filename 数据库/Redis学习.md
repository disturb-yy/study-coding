

# 知识体系

![img](https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301311802863.png)



# 相关学习

[♥Redis教程 - Redis知识体系详解♥ | Java 全栈知识体系 (pdai.tech)](https://www.pdai.tech/md/db/nosql-redis/db-redis-overview.html)







# Redis 入门



## Redis数据结构简介

首先对redis来说，**所有的key（键）都是字符串。我们在谈基础数据结构时，讨论的是存储值的数据类型**，主要包括常见的5种数据类型，分别是：String、List、Set、Zset、Hash。

Redis 通常被称为数据结构服务器，因为值（value）可以是字符串(String)、哈希(Hash)、列表(list)、集合(sets)和有序集合(sorted sets)等类型。

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301312036873.jpeg" alt="img" style="zoom: 80%;" />









## [# ](./Redis 入门 - 5 种基础数据类型详解.md)数据类型：5种基础数据类型详解

==Redis所有的key（键）都是字符串。**我们在谈基础数据结构时，讨论的是存储值的数据类型**，主要包括常见的5种数据类型，分别是：String、List、Set、Zset、Hash。==



### String 字符串

Redis 字符串存储字节序列，包括文本、序列化对象和二进制数组。 因此，**字符串是最基本的 Redis 数据类型**。 它们通常用于缓存，但它们支持其他功能，这些功能也允许您实现计数器并执行按位运算。



#### 限制

默认情况下，单个 Redis 字符串的最大大小为 512 MB。



#### 基本命令

##### 获取和设置字符串

- [`SET`](https://redis.io/commands/set) 存储字符串值。
- 仅当字符串值尚不存在时，[`SETNX`](https://redis.io/commands/setnx) 才会存储该值。对于实现锁很有用。（==如果不存在才保存值==）
- [`GET`](https://redis.io/commands/get) 检索字符串值。
- [`MGET`](https://redis.io/commands/mget) 在单个操作中检索多个字符串值。

##### 管理计数器

- [`INCRBY`](https://redis.io/commands/incrby)以原子方式递增（并在传递负数时递减）存储在给定键上的计数器。
- 浮点计数器存在另一个命令：[INCRBYFLOAT](https://redis.io/commands/incrbyfloat)。

##### 按位运算

若要对字符串执行按位运算，请参阅[位图数据类型](https://redis.io/docs/data-types/bitmaps)文档。

请参阅[字符串命令的完整列表](https://redis.io/commands/?group=string)。

##### 性能

大多数字符串操作都是 O（1），这意味着它们非常高效。 但是，请注意 [`SUBSTR`](https://redis.io/commands/substr)、[`GETRANGE`](https://redis.io/commands/getrange) 和 [`SETRANGE`](https://redis.io/commands/setrange) 命令，它们可以是 O（n）。 处理大型字符串时，这些随机访问字符串命令可能会导致性能问题。

##### 选择

如果要将结构化数据存储为序列化字符串，则可能还需要考虑 [Redis 哈希](https://redis.io/docs/data-types/hashes)或 [RedisJSON。](https://redis.io/docs/stack/json)

******



### List 列表



**Redis 列表是字符串值的链接列表。** Redis 列表经常用于：

- 实现堆栈和队列。
- 为后台工作程序系统构建队列管理。



#### 限制

Redis 列表的最大长度为 2^32 - 1 （4,294,967,295） 个元素。



#### 基本命令

- [`LPUSH`](https://redis.io/commands/lpush) 在列表的头部添加一个新元素; [`RPUSH`](https://redis.io/commands/rpush)添加到尾部。
- [`LPOP`](https://redis.io/commands/lpop) 从列表的头部删除并返回一个元素;[`RPOP`](https://redis.io/commands/rpop) 执行相同的操作，但从列表的尾部。
- [`LLEN`](https://redis.io/commands/llen) 返回列表的长度。
- [`LMOVE`](https://redis.io/commands/lmove) 以原子方式将元素从一个列表移动到另一个列表。
- [`LTRIM`](https://redis.io/commands/ltrim) 将列表缩减到指定的元素范围。

==可以使用基本命令来生成先进先出的队列和先进后出的栈==



#### 阻止命令

列表支持多个阻止命令==（如channel的效果）==。 例如：

- [`BLPOP`](https://redis.io/commands/blpop) 从列表的头部删除并返回一个元素。 **如果列表为空，则该命令将一直阻止，直到元素可用或达到指定的超时。**
- [`BLMOVE`](https://redis.io/commands/blmove) 以原子方式将元素从源列表移动到目标列表。 如果源列表为空，则该命令将阻塞，直到新元素可用。

请参阅[完整的列表命令系列](https://redis.io/commands/?group=list)。



#### 性能

访问其头部或尾部的列表操作是 O（1），这意味着它们非常高效。 但是，操作列表中元素的命令通常是 O（n）。 这些示例包括 [`LINDEX`](https://redis.io/commands/lindex)、[`LINSERT`](https://redis.io/commands/linsert) 和 [`LSET。`](https://redis.io/commands/lset) 运行这些命令时要小心，主要是在对大型列表进行操作时。



#### 选择

当您需要存储和处理一系列不确定的事件时，请考虑[使用 Redis 流](https://redis.io/docs/data-types/streams)作为列表的替代方法。

-----------------------



### sets 集合

**Redis 集合** 是唯一字符串（成员）的无序集合（sets的类型类似数据结构中的set）。 您可以使用 Redis Sets 有效地：

- 跟踪唯一项目（例如，跟踪访问给定博客文章的所有唯一 IP 地址）。
- 表示关系（例如，具有给定角色的所有用户的集合）。
- 执行常见的集合操作，例如交集、并集和差分。



#### 限制

Redis 集合 的最大大小为 2^32 - 1 （4,294,967,295） 个成员。



#### 基本命令

- [`SADD`](https://redis.io/commands/sadd) 将新成员添加到集合中。（Set Add）
- [`SREM`](https://redis.io/commands/srem)从集合中删除指定的成员。（Set Removed）
- [`SISMEMBER`](https://redis.io/commands/sismember)测试字符串的集合成员资格。
- [`SINTER`](https://redis.io/commands/sinter) 返回两个或多个集合共有的成员集（即交集）。
- [`SCARD`](https://redis.io/commands/scard) 返回集合的大小（也称为基数）。  (Set Card)

请参阅[设置命令的完整列表](https://redis.io/commands/?group=set)。



#### 性能

大多数集合操作（包括添加、删除和检查项是否为集合成员）都是 O（1）。 这意味着它们非常高效。 但是，对于具有数十万或更多成员的大型集，在运行 [`SMEMBERS`](https://redis.io/commands/smembers) 命令时应格外小心。 此命令为 O（n），并在单个响应中返回整个集合。 作为替代方法，请考虑 [`SSCAN，`](https://redis.io/commands/sscan)它允许您以迭代方式检索集合的所有成员。



#### 选择

对大型数据集（或流数据）进行的成员资格检查可能会占用大量内存。 如果您担心内存使用并且不需要完美的精度，请考虑使用 [Bloom 过滤器或 Cuckoo 过滤器](https://redis.io/docs/stack/bloom)作为集合的替代方案。

Redis 集经常用作一种索引。 如果需要索引和查询数据，请考虑 [RediSearch](https://redis.io/docs/stack/search) 和 [RedisJSON。](https://redis.io/docs/stack/json)

--------------------





### hashes 散列

**Redis hashes 是结构化为字段值对==集合==的记录类型（即hashes的）。** 您可以使用哈希来表示基本对象和存储计数器分组等。



#### 基本命令

- [`HSET`](https://redis.io/commands/hset) 在哈希上设置一个或多个字段的值。
- [`HGET`](https://redis.io/commands/hget) 返回给定字段的值。
- [`HMGET`](https://redis.io/commands/hmget) 返回一个或多个给定字段的值。
- [`HINCRBY`](https://redis.io/commands/hincrby) 将给定字段的值按提供的整数递增。

请参阅[哈希命令的完整列表](https://redis.io/commands/?group=hash)。



#### 性能

大多数 Redis 哈希命令都是 O（1）。

一些命令（如 [`HKEYS、`](https://redis.io/commands/hkeys)[`HVALS` 和](https://redis.io/commands/hvals) [`HGETALL`](https://redis.io/commands/hgetall)）是 O（n），其中 n 是字段值对的数量。



## [#](./Redis 入门 - 3 种特殊类型详解.md) 数据类型：3 种特殊类型详解

Redis 除了上文中 5 种基础数据类型，还有三种特殊的数据类型，分别是 HyperLogLogs（基数统计）， Bitmaps (位图) 和 geospatial（地理位置）。

*   **HyperLogLogs 基数统计用来解决什么问题**？

这个结构可以非常省内存的去统计各种计数，比如注册 IP 数、每日访问 IP 数、页面实时 UV、在线用户数，共同好友数等。

*   **Bitmap用来解决什么问题**？

比如：统计用户信息，活跃，不活跃！ 登录，未登录！ 打卡，不打卡！ **两个状态的，都可以使用 Bitmaps**！

- **geospatial 用来推算地理位置的信息: 两地之间的距离, 方圆几里的人**



-------------------





## [#](./Redis 入门 - Stream 详解.md) 数据类型：Stream 详解

Redis5.0 中还增加了一个数据类型 Stream，它借鉴了 Kafka 的设计，是一个新的强大的支持多播的可持久化的消息队列。





-----------------------



# Redis 进阶



## [#](./Redis 进阶 -对象机制详解md.md) 对象机制详解