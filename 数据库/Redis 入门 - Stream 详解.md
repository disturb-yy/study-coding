> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.pdai.tech](https://www.pdai.tech/md/db/nosql-redis/db-redis-data-type-stream.html)

> Redis5.0 中还增加了一个数据类型 Stream，它借鉴了 Kafka 的设计，是一个新的强大的支持多播的可持久化的消息队列

> Redis5.0 中还增加了一个数据类型 Stream，它借鉴了 Kafka 的设计，是一个新的强大的支持多播的可持久化的消息队列。@pdai

*   [Redis 入门 - 数据类型：Stream 详解](#redis%E5%85%A5%E9%97%A8---%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8Bstream%E8%AF%A6%E8%A7%A3)
    *   [为什么会设计 Stream](#%E4%B8%BA%E4%BB%80%E4%B9%88%E4%BC%9A%E8%AE%BE%E8%AE%A1stream)
    *   [Stream 详解](#stream%E8%AF%A6%E8%A7%A3)
        *   [Stream 的结构](#stream%E7%9A%84%E7%BB%93%E6%9E%84)
        *   [增删改查](#%E5%A2%9E%E5%88%A0%E6%94%B9%E6%9F%A5)
        *   [独立消费](#%E7%8B%AC%E7%AB%8B%E6%B6%88%E8%B4%B9)
        *   [消费组消费](#%E6%B6%88%E8%B4%B9%E7%BB%84%E6%B6%88%E8%B4%B9)
        *   [信息监控](#%E4%BF%A1%E6%81%AF%E7%9B%91%E6%8E%A7)
    *   [更深入理解](#%E6%9B%B4%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3)
        *   [Stream 用在什么样场景](#stream%E7%94%A8%E5%9C%A8%E4%BB%80%E4%B9%88%E6%A0%B7%E5%9C%BA%E6%99%AF)
        *   [消息 ID 的设计是否考虑了时间回拨的问题？](#%E6%B6%88%E6%81%AFid%E7%9A%84%E8%AE%BE%E8%AE%A1%E6%98%AF%E5%90%A6%E8%80%83%E8%99%91%E4%BA%86%E6%97%B6%E9%97%B4%E5%9B%9E%E6%8B%A8%E7%9A%84%E9%97%AE%E9%A2%98)
        *   [消费者崩溃带来的会不会消息丢失问题?](#%E6%B6%88%E8%B4%B9%E8%80%85%E5%B4%A9%E6%BA%83%E5%B8%A6%E6%9D%A5%E7%9A%84%E4%BC%9A%E4%B8%8D%E4%BC%9A%E6%B6%88%E6%81%AF%E4%B8%A2%E5%A4%B1%E9%97%AE%E9%A2%98)
        *   [消费者彻底宕机后如何转移给其它消费者处理？](#%E6%B6%88%E8%B4%B9%E8%80%85%E5%BD%BB%E5%BA%95%E5%AE%95%E6%9C%BA%E5%90%8E%E5%A6%82%E4%BD%95%E8%BD%AC%E7%A7%BB%E7%BB%99%E5%85%B6%E5%AE%83%E6%B6%88%E8%B4%B9%E8%80%85%E5%A4%84%E7%90%86)
        *   [坏消息问题，Dead Letter，死信问题](#%E5%9D%8F%E6%B6%88%E6%81%AF%E9%97%AE%E9%A2%98dead-letter%E6%AD%BB%E4%BF%A1%E9%97%AE%E9%A2%98)
    *   [参考文章](#%E5%8F%82%E8%80%83%E6%96%87%E7%AB%A0)

[#](#为什么会设计stream) 为什么会设计 Stream
--------------------------------

> Redis5.0 中还增加了一个数据结构 Stream，从字面上看是流类型，但其实从功能上看，应该是 Redis 对消息队列（MQ，Message Queue）的完善实现。

用过 Redis 做消息队列的都了解，基于 Reids 的消息队列实现有很多种，例如：

*   **PUB/SUB，订阅 / 发布模式**
    *   但是发布订阅模式是**无法持久化**的，如果出现网络断开、Redis 宕机等，消息就会被丢弃；
*   基于 **List LPUSH+BRPOP** 或者 **基于 Sorted-Set** 的实现
    *   支持了持久化，但是不支持多播，分组消费等

为什么上面的结构无法满足广泛的 MQ 场景？ 这里便引出一个核心的问题：如果我们期望设计一种数据结构来实现消息队列，最重要的就是要理解**设计一个消息队列需要考虑什么**？初步的我们很容易想到

*   消息的生产
*   消息的消费
    *   单播和多播（多对多）
    *   阻塞和非阻塞读取
*   消息有序性
*   消息的持久化

其它还要考虑啥嗯？借助美团技术团队的一篇文章，[消息队列设计精要在新窗口打开](https://tech.meituan.com/2016/07/01/mq-design.html) 中的图

![](https://www.pdai.tech/images/db/redis/db-redis-stream-1.png)

**我们不妨看看 Redis 考虑了哪些设计**？

*   消息 ID 的序列化生成
*   消息遍历
*   消息的阻塞和非阻塞读取
*   消息的分组消费
*   未完成消息的处理
*   消息队列监控
*   ...

> 这也是我们需要理解 Stream 的点，但是结合上面的图，我们也应该理解 Redis Stream 也是一种超轻量 MQ 并没有完全实现消息队列所有设计要点，这决定着它适用的场景。

[#](#stream详解) Stream 详解
------------------------

> 经过梳理总结，我认为从以下几个大的方面去理解 Stream 是比较合适的，总结如下：@pdai

*   Stream 的结构设计
*   生产和消费
    *   基本的增删查改
    *   单一消费者的消费
    *   消费组的消费
*   监控状态

### [#](#stream的结构) Stream 的结构

每个 Stream 都有唯一的名称，它就是 Redis 的 key，在我们首次使用 xadd 指令追加消息时自动创建。

![](https://www.pdai.tech/images/db/redis/db-redis-stream-2.png)

上图解析：

*   `Consumer Group` ：消费组，使用 XGROUP CREATE 命令创建，一个消费组有多个消费者 (Consumer), 这些消费者之间是竞争关系。
*   `last_delivered_id` ：游标，每个消费组会有个游标 last_delivered_id，任意一个消费者读取了消息都会使游标 last_delivered_id 往前移动。
*   `pending_ids` ：消费者 (Consumer) 的状态变量，作用是维护消费者的未确认的 id。 pending_ids 记录了当前已经被客户端读取的消息，但是还没有 `ack` (Acknowledge character：确认字符）。如果客户端没有 ack，这个变量里面的消息 ID 会越来越多，一旦某个消息被 ack，它就开始减少。这个 pending_ids 变量在 Redis 官方被称之为 PEL，也就是 Pending Entries List，这是一个很核心的数据结构，它用来确保客户端至少消费了消息一次，而不会在网络传输的中途丢失了没处理。

此外我们还需要理解两点：

*   `消息ID`: 消息 ID 的形式是 timestampInMillis-sequence，例如 1527846880572-5，它表示当前的消息在毫米时间戳 1527846880572 时产生，并且是该毫秒内产生的第 5 条消息。消息 ID 可以由服务器自动生成，也可以由客户端自己指定，但是形式必须是整数 - 整数，而且必须是后面加入的消息的 ID 要大于前面的消息 ID。
*   `消息内容`: 消息内容就是键值对，形如 hash 结构的键值对，这没什么特别之处。

### [#](#增删改查) 增删改查

消息队列相关命令：

*   XADD - 添加消息到末尾
*   XTRIM - 对流进行修剪，限制长度
*   XDEL - 删除消息
*   XLEN - 获取流包含的元素数量，即消息长度
*   XRANGE - 获取消息列表，会自动过滤已经删除的消息
*   XREVRANGE - 反向获取消息列表，ID 从大到小
*   XREAD - 以阻塞或非阻塞方式获取消息列表

```
127.0.0.1:6379> xadd codehole * name laoqian age 30  
1527849609889-0  
127.0.0.1:6379> xadd codehole * name xiaoyu age 29
1527849629172-0
127.0.0.1:6379> xadd codehole * name xiaoqian age 1
1527849637634-0
127.0.0.1:6379> xlen codehole
(integer) 3
127.0.0.1:6379> xrange codehole - +  
127.0.0.1:6379> xrange codehole - +
1) 1) 1527849609889-0
   1) 1) "name"
      1) "laoqian"
      2) "age"
      3) "30"
2) 1) 1527849629172-0
   1) 1) "name"
      1) "xiaoyu"
      2) "age"
      3) "29"
3) 1) 1527849637634-0
   1) 1) "name"
      1) "xiaoqian"
      2) "age"
      3) "1"
127.0.0.1:6379> xrange codehole 1527849629172-0 +  
1) 1) 1527849629172-0
   2) 1) "name"
      2) "xiaoyu"
      3) "age"
      4) "29"
2) 1) 1527849637634-0
   2) 1) "name"
      2) "xiaoqian"
      3) "age"
      4) "1"
127.0.0.1:6379> xrange codehole - 1527849629172-0  
1) 1) 1527849609889-0
   2) 1) "name"
      2) "laoqian"
      3) "age"
      4) "30"
2) 1) 1527849629172-0
   2) 1) "name"
      2) "xiaoyu"
      3) "age"
      4) "29"
127.0.0.1:6379> xdel codehole 1527849609889-0
(integer) 1
127.0.0.1:6379> xlen codehole  
(integer) 2
127.0.0.1:6379> xrange codehole - +  
1) 1) 1527849629172-0
   2) 1) "name"
      2) "xiaoyu"
      3) "age"
      4) "29"
2) 1) 1527849637634-0
   2) 1) "name"
      2) "xiaoqian"
      3) "age"
      4) "1"
127.0.0.1:6379> del codehole  
(integer) 1


```

### [#](#独立消费) 独立消费

我们可以在不定义消费组的情况下进行 Stream 消息的独立消费，当 Stream 没有新消息时，甚至可以阻塞等待。Redis 设计了一个单独的消费指令 xread，可以将 Stream 当成普通的消息队列 (list) 来使用。使用 xread 时，我们可以完全忽略消费组 (Consumer Group) 的存在，就好比 Stream 就是一个普通的列表(list)。

```
127.0.0.1:6379> xread count 2 streams codehole 0-0
1) 1) "codehole"
   2) 1) 1) 1527851486781-0
         2) 1) "name"
            2) "laoqian"
            3) "age"
            4) "30"
      2) 1) 1527851493405-0
         2) 1) "name"
            2) "yurui"
            3) "age"
            4) "29"

127.0.0.1:6379> xread count 1 streams codehole $
(nil)

127.0.0.1:6379> xread block 0 count 1 streams codehole $

127.0.0.1:6379> xadd codehole * name youming age 60
1527852774092-0


127.0.0.1:6379> xread block 0 count 1 streams codehole $
1) 1) "codehole"
   2) 1) 1) 1527852774092-0
         2) 1) "name"
            2) "youming"
            3) "age"
            4) "60"
(93.11s)


```

客户端如果想要使用 xread 进行顺序消费，一定要记住当前消费到哪里了，也就是返回的消息 ID。下次继续调用 xread 时，将上次返回的最后一个消息 ID 作为参数传递进去，就可以继续消费后续的消息。

block 0 表示永远阻塞，直到消息到来，block 1000 表示阻塞 1s，如果 1s 内没有任何消息到来，就返回 nil

```
127.0.0.1:6379> xread block 1000 count 1 streams codehole $
(nil)
(1.07s)


```

### [#](#消费组消费) 消费组消费

*   **消费组消费图**

![](https://www.pdai.tech/images/db/redis/db-redis-stream-3.png)

*   相关命令：
    
    *   XGROUP CREATE - 创建消费者组
    *   XREADGROUP GROUP - 读取消费者组中的消息
    *   XACK - 将消息标记为 "已处理"
    *   XGROUP SETID - 为消费者组设置新的最后递送消息 ID
    *   XGROUP DELCONSUMER - 删除消费者
    *   XGROUP DESTROY - 删除消费者组
    *   XPENDING - 显示待处理消息的相关信息
    *   XCLAIM - 转移消息的归属权
    *   XINFO - 查看流和消费者组的相关信息；
    *   XINFO GROUPS - 打印消费者组的信息；
    *   XINFO STREAM - 打印流信息
*   **创建消费组**
    

Stream 通过 xgroup create 指令创建消费组 (Consumer Group)，需要传递起始消息 ID 参数用来初始化 last_delivered_id 变量。

```
127.0.0.1:6379> xgroup create codehole cg1 0-0  
OK

127.0.0.1:6379> xgroup create codehole cg2 $
OK
127.0.0.1:6379> xinfo stream codehole  
 1) length
 2) (integer) 3  
 3) radix-tree-keys
 4) (integer) 1
 5) radix-tree-nodes
 6) (integer) 2
 7) groups
 8) (integer) 2  
 9) first-entry  
10) 1) 1527851486781-0
    2) 1) "name"
       2) "laoqian"
       3) "age"
       4) "30"
11) last-entry  
12) 1) 1527851498956-0
    2) 1) "name"
       2) "xiaoqian"
       3) "age"
       4) "1"
127.0.0.1:6379> xinfo groups codehole  
1) 1) name
   2) "cg1"
   3) consumers
   4) (integer) 0  
   5) pending
   6) (integer) 0  
2) 1) name
   2) "cg2"
   3) consumers  
   4) (integer) 0
   5) pending
   6) (integer) 0  


```

*   **消费组消费**

Stream 提供了 xreadgroup 指令可以进行消费组的组内消费，需要提供消费组名称、消费者名称和起始消息 ID。它同 xread 一样，也可以阻塞等待新消息。读到新消息后，对应的消息 ID 就会进入消费者的 PEL(正在处理的消息) 结构里，客户端处理完毕后使用 xack 指令通知服务器，本条消息已经处理完毕，该消息 ID 就会从 PEL 中移除。

```
127.0.0.1:6379> xreadgroup GROUP cg1 c1 count 1 streams codehole >
1) 1) "codehole"
   2) 1) 1) 1527851486781-0
         2) 1) "name"
            2) "laoqian"
            3) "age"
            4) "30"
127.0.0.1:6379> xreadgroup GROUP cg1 c1 count 1 streams codehole >
1) 1) "codehole"
   2) 1) 1) 1527851493405-0
         2) 1) "name"
            2) "yurui"
            3) "age"
            4) "29"
127.0.0.1:6379> xreadgroup GROUP cg1 c1 count 2 streams codehole >
1) 1) "codehole"
   2) 1) 1) 1527851498956-0
         2) 1) "name"
            2) "xiaoqian"
            3) "age"
            4) "1"
      2) 1) 1527852774092-0
         2) 1) "name"
            2) "youming"
            3) "age"
            4) "60"

127.0.0.1:6379> xreadgroup GROUP cg1 c1 count 1 streams codehole >
(nil)

127.0.0.1:6379> xreadgroup GROUP cg1 c1 block 0 count 1 streams codehole >

127.0.0.1:6379> xadd codehole * name lanying age 61
1527854062442-0

127.0.0.1:6379> xreadgroup GROUP cg1 c1 block 0 count 1 streams codehole >
1) 1) "codehole"
   2) 1) 1) 1527854062442-0
         2) 1) "name"
            2) "lanying"
            3) "age"
            4) "61"
(36.54s)
127.0.0.1:6379> xinfo groups codehole  
1) 1) name
   2) "cg1"
   3) consumers
   4) (integer) 1  
   5) pending
   6) (integer) 5  
2) 1) name
   2) "cg2"
   3) consumers
   4) (integer) 0  
   5) pending
   6) (integer) 0

127.0.0.1:6379> xinfo consumers codehole cg1  
1) 1) name
   2) "c1"
   3) pending
   4) (integer) 5  
   5) idle
   6) (integer) 418715  

127.0.0.1:6379> xack codehole cg1 1527851486781-0
(integer) 1
127.0.0.1:6379> xinfo consumers codehole cg1
1) 1) name
   2) "c1"
   3) pending
   4) (integer) 4  
   5) idle
   6) (integer) 668504

127.0.0.1:6379> xack codehole cg1 1527851493405-0 1527851498956-0 1527852774092-0 1527854062442-0
(integer) 4
127.0.0.1:6379> xinfo consumers codehole cg1
1) 1) name
   2) "c1"
   3) pending
   4) (integer) 0  
   5) idle
   6) (integer) 745505


```

### [#](#信息监控) 信息监控

Stream 提供了 XINFO 来实现对服务器信息的监控，可以查询：

*   查看队列信息

```
127.0.0.1:6379> Xinfo stream mq
 1) "length"
 2) (integer) 7
 3) "radix-tree-keys"
 4) (integer) 1
 5) "radix-tree-nodes"
 6) (integer) 2
 7) "groups"
 8) (integer) 1
 9) "last-generated-id"
10) "1553585533795-9"
11) "first-entry"
12) 1) "1553585533795-3"
    2) 1) "msg"
       2) "4"
13) "last-entry"
14) 1) "1553585533795-9"
    2) 1) "msg"
       2) "10"


```

*   消费组信息

```
127.0.0.1:6379> Xinfo groups mq
1) 1) "name"
   2) "mqGroup"
   3) "consumers"
   4) (integer) 3
   5) "pending"
   6) (integer) 3
   7) "last-delivered-id"
   8) "1553585533795-4"


```

*   消费者组成员信息

```
127.0.0.1:6379> XINFO CONSUMERS mq mqGroup
1) 1) "name"
   2) "consumerA"
   3) "pending"
   4) (integer) 1
   5) "idle"
   6) (integer) 18949894
2) 1) "name"
   2) "consumerB"
   3) "pending"
   4) (integer) 1
   5) "idle"
   6) (integer) 3092719
3) 1) "name"
   2) "consumerC"
   3) "pending"
   4) (integer) 1
   5) "idle"
   6) (integer) 23683256


```

至此，消息队列的操作说明大体结束！

[#](#更深入理解) 更深入理解
-----------------

> 我们结合 MQ 中常见问题，看 Redis 是如何解决的，来进一步理解 Redis。

### [#](#stream用在什么样场景) Stream 用在什么样场景

可用作时通信等，大数据分析，异地数据备份等

![](https://www.pdai.tech/images/db/redis/db-redis-stream-4.png)

客户端可以平滑扩展，提高处理能力

![](https://www.pdai.tech/images/db/redis/db-redis-stream-5.png)

### [#](#消息id的设计是否考虑了时间回拨的问题) 消息 ID 的设计是否考虑了时间回拨的问题？

> 在 [分布式算法 - ID 算法](https://www.pdai.tech/md/algorithm/alg-domain-id-snowflake.html)设计中, 一个常见的问题就是时间回拨问题，那么 Redis 的消息 ID 设计中是否考虑到这个问题呢？

XADD 生成的 1553439850328-0，就是 Redis 生成的消息 ID，由两部分组成: **时间戳 - 序号**。时间戳是毫秒级单位，是生成消息的 Redis 服务器时间，它是个 64 位整型（int64）。序号是在这个毫秒时间点内的消息序号，它也是个 64 位整型。

可以通过 multi 批处理，来验证序号的递增：

```
127.0.0.1:6379> MULTI
OK
127.0.0.1:6379> XADD memberMessage * msg one
QUEUED
127.0.0.1:6379> XADD memberMessage * msg two
QUEUED
127.0.0.1:6379> XADD memberMessage * msg three
QUEUED
127.0.0.1:6379> XADD memberMessage * msg four
QUEUED
127.0.0.1:6379> XADD memberMessage * msg five
QUEUED
127.0.0.1:6379> EXEC
1) "1553441006884-0"
2) "1553441006884-1"
3) "1553441006884-2"
4) "1553441006884-3"
5) "1553441006884-4"


```

由于一个 redis 命令的执行很快，所以可以看到在同一时间戳内，是通过序号递增来表示消息的。

为了保证消息是有序的，因此 Redis 生成的 ID 是单调递增有序的。由于 ID 中包含时间戳部分，为了避免服务器时间错误而带来的问题（例如服务器时间延后了），Redis 的每个 Stream 类型数据都维护一个 latest_generated_id 属性，用于记录最后一个消息的 ID。**若发现当前时间戳退后（小于 latest_generated_id 所记录的），则采用时间戳不变而序号递增的方案来作为新消息 ID**（这也是序号为什么使用 int64 的原因，保证有足够多的的序号），从而保证 ID 的单调递增性质。

强烈建议使用 Redis 的方案生成消息 ID，因为这种时间戳 + 序号的单调递增的 ID 方案，几乎可以满足你全部的需求。但同时，记住 ID 是支持自定义的，别忘了！

### [#](#消费者崩溃带来的会不会消息丢失问题) 消费者崩溃带来的会不会消息丢失问题?

为了解决组内消息读取但处理期间消费者崩溃带来的消息丢失问题，STREAM 设计了 Pending 列表，用于记录读取但并未处理完毕的消息。命令 XPENDIING 用来获消费组或消费内消费者的未处理完毕的消息。演示如下：

```
127.0.0.1:6379> XPENDING mq mqGroup 
1) (integer) 5 
2) "1553585533795-0" 
3) "1553585533795-4" 
4) 1) 1) "consumerA" 
      2) "3"
   2) 1) "consumerB" 
      2) "1"
   3) 1) "consumerC" 
      2) "1"

127.0.0.1:6379> XPENDING mq mqGroup - + 10 
1) 1) "1553585533795-0" 
   2) "consumerA" 
   3) (integer) 1654355 
   4) (integer) 5 
2) 1) "1553585533795-1"
   2) "consumerA"
   3) (integer) 1654355
   4) (integer) 4


127.0.0.1:6379> XPENDING mq mqGroup - + 10 consumerA 
1) 1) "1553585533795-0"
   2) "consumerA"
   3) (integer) 1641083
   4) (integer) 5



```

每个 Pending 的消息有 4 个属性：

*   消息 ID
*   所属消费者
*   IDLE，已读取时长
*   delivery counter，消息被读取次数

上面的结果我们可以看到，我们之前读取的消息，都被记录在 Pending 列表中，说明全部读到的消息都没有处理，仅仅是读取了。那如何表示消费者处理完毕了消息呢？使用命令 XACK 完成告知消息处理完成，演示如下：

```
127.0.0.1:6379> XACK mq mqGroup 1553585533795-0 
(integer) 1

127.0.0.1:6379> XPENDING mq mqGroup 
1) (integer) 4 
2) "1553585533795-1"
3) "1553585533795-4"
4) 1) 1) "consumerA" 
      2) "2"
   2) 1) "consumerB"
      2) "1"
   3) 1) "consumerC"
      2) "1"
127.0.0.1:6379>


```

有了这样一个 Pending 机制，就意味着在某个消费者读取消息但未处理后，消息是不会丢失的。等待消费者再次上线后，可以读取该 Pending 列表，就可以继续处理该消息了，保证消息的有序和不丢失。

### [#](#消费者彻底宕机后如何转移给其它消费者处理) 消费者彻底宕机后如何转移给其它消费者处理？

> 还有一个问题，就是若某个消费者宕机之后，没有办法再上线了，那么就需要将该消费者 Pending 的消息，转义给其他的消费者处理，就是消息转移。

消息转移的操作时将某个消息转移到自己的 Pending 列表中。使用语法 XCLAIM 来实现，需要设置组、转移的目标消费者和消息 ID，同时需要提供 IDLE（已被读取时长），只有超过这个时长，才能被转移。演示如下：

```
127.0.0.1:6379> XPENDING mq mqGroup - + 10
1) 1) "1553585533795-1"
   2) "consumerA"
   3) (integer) 15907787
   4) (integer) 4


127.0.0.1:6379> XCLAIM mq mqGroup consumerB 3600000 1553585533795-1
1) 1) "1553585533795-1"
   2) 1) "msg"
      2) "2"


127.0.0.1:6379> XPENDING mq mqGroup - + 10
1) 1) "1553585533795-1"
   2) "consumerB"
   3) (integer) 84404 
   4) (integer) 5 


```

以上代码，完成了一次消息转移。转移除了要指定 ID 外，还需要指定 IDLE，保证是长时间未处理的才被转移。被转移的消息的 IDLE 会被重置，用以保证不会被重复转移，以为可能会出现将过期的消息同时转移给多个消费者的并发操作，设置了 IDLE，则可以避免后面的转移不会成功，因为 IDLE 不满足条件。例如下面的连续两条转移，第二条不会成功。

```
127.0.0.1:6379> XCLAIM mq mqGroup consumerB 3600000 1553585533795-1
127.0.0.1:6379> XCLAIM mq mqGroup consumerC 3600000 1553585533795-1


```

这就是消息转移。至此我们使用了一个 Pending 消息的 ID，所属消费者和 IDLE 的属性，还有一个属性就是消息被读取次数，delivery counter，该属性的作用由于统计消息被读取的次数，包括被转移也算。这个属性主要用在判定是否为错误数据上。

### [#](#坏消息问题-dead-letter-死信问题) 坏消息问题，Dead Letter，死信问题

正如上面所说，如果某个消息，不能被消费者处理，也就是不能被 XACK，这是要长时间处于 Pending 列表中，即使被反复的转移给各个消费者也是如此。此时该消息的 delivery counter 就会累加（上一节的例子可以看到），当累加到某个我们预设的临界值时，我们就认为是坏消息（也叫死信，DeadLetter，无法投递的消息），由于有了判定条件，我们将坏消息处理掉即可，删除即可。删除一个消息，使用 XDEL 语法，演示如下：

```
127.0.0.1:6379> XDEL mq 1553585533795-1
(integer) 1

127.0.0.1:6379> XRANGE mq - +
1) 1) "1553585533795-0"
   2) 1) "msg"
      2) "1"
2) 1) "1553585533795-2"
   2) 1) "msg"
      2) "3"


```

注意本例中，并没有删除 Pending 中的消息因此你查看 Pending，消息还会在。可以执行 XACK 标识其处理完毕！

[#](#参考文章) 参考文章
---------------

本文主要梳理总结自：

*   https://www.runoob.com/redis/redis-stream.html
*   https://www.zhihu.com/question/279540635
*   https://zhuanlan.zhihu.com/p/60501638