> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.pdai.tech](https://www.pdai.tech/md/db/nosql-redis/db-redis-x-redis-object.html)

> 我们在前文已经阐述了 Redis 5 种基础数据类型详解，分别是字符串 (string)、列表 (list)、哈希 (hash)、集合 (set)、有序集合 (zset)，以及 5.0 版本中 Redis Stream 结构详解；那么这些基础类型的底层是如何实现的呢？Redis 的每种对象其实都由 ** 对象结构 (redisObject)** 与 ** 对应编码的数据结构 ** 组合而成, 本文主要介绍 ** 对象结构 (redisObject)** 部分。

> 我们在前文已经阐述了 [Redis 5 种基础数据类型详解](https://www.pdai.tech/md/db/nosql-redis/db-redis-data-types.html)，分别是字符串 (string)、列表 (list)、哈希 (hash)、集合 (set)、有序集合 (zset)，以及 5.0 版本中 [Redis Stream 结构详解](https://www.pdai.tech/md/db/nosql-redis/db-redis-data-type-stream.html)；那么这些基础类型的底层是如何实现的呢？Redis 的每种对象其实都由**对象结构 (redisObject)** 与 **对应编码的数据结构**组合而成, 本文主要介绍**对象结构 (redisObject)** 部分。@pdai

*   [Redis 进阶 - 数据结构：对象机制详解](#redis%E8%BF%9B%E9%98%B6---%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%AF%B9%E8%B1%A1%E6%9C%BA%E5%88%B6%E8%AF%A6%E8%A7%A3)
    *   [引入: 从哪里开始学习底层？](#%E5%BC%95%E5%85%A5%E4%BB%8E%E5%93%AA%E9%87%8C%E5%BC%80%E5%A7%8B%E5%AD%A6%E4%B9%A0%E5%BA%95%E5%B1%82)
    *   [为什么 Redis 会设计 redisObject 对象](#%E4%B8%BA%E4%BB%80%E4%B9%88redis%E4%BC%9A%E8%AE%BE%E8%AE%A1redisobject%E5%AF%B9%E8%B1%A1)
    *   [redisObject 数据结构](#redisobject%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)
    *   [命令的类型检查和多态](#%E5%91%BD%E4%BB%A4%E7%9A%84%E7%B1%BB%E5%9E%8B%E6%A3%80%E6%9F%A5%E5%92%8C%E5%A4%9A%E6%80%81)
    *   [对象共享](#%E5%AF%B9%E8%B1%A1%E5%85%B1%E4%BA%AB)
    *   [引用计数以及对象的消毁](#%E5%BC%95%E7%94%A8%E8%AE%A1%E6%95%B0%E4%BB%A5%E5%8F%8A%E5%AF%B9%E8%B1%A1%E7%9A%84%E6%B6%88%E6%AF%81)
    *   [小结](#%E5%B0%8F%E7%BB%93)
    *   [参考文章](#%E5%8F%82%E8%80%83%E6%96%87%E7%AB%A0)

[#](#引入-从哪里开始学习底层) 引入: 从哪里开始学习底层？
---------------------------------

> 我在整理 Redis 底层设计时，发现网上其实是有很多资料的，但是缺少一种自上而下的承接。这里我将收集很多网上的资料并重新组织，来帮助你更好的理解 Redis 底层设计。@pdai

我们在前文已经阐述了 [Redis 5 种基础数据类型详解](https://www.pdai.tech/md/db/nosql-redis/db-redis-data-types.html)和 [Redis 入门 - 数据结构：Stream 详解](https://www.pdai.tech/md/db/nosql-redis/db-redis-data-type-stream.html)；那么这些基础类型的底层是如何实现的呢？

带着这个问题我们来着手理解底层设计，首先看下图：

![](https://www.pdai.tech/images/db/redis/db-redis-object-2-2.png)

它反映了 Redis 的每种对象其实都由**对象结构 (redisObject)** 与 **对应编码的数据结构**组合而成，而每种对象类型对应若干编码方式，不同的编码方式所对应的底层数据结构是不同的。

所以，我们需要从几个个角度来着手底层研究：

*   **对象设计机制**: 对象结构 (redisObject)
*   **编码类型和底层数据结构**: 对应编码的数据结构

[#](#为什么redis会设计redisobject对象) 为什么 Redis 会设计 redisObject 对象
-----------------------------------------------------------

> 为什么 Redis 会设计 redisObject 对象？

在 redis 的命令中，用于对键进行处理的命令占了很大一部分，而对于键所保存的值的类型（键的类型），键能执行的命令又各不相同。如： `LPUSH` 和 `LLEN` 只能用于列表键, 而 `SADD` 和 `SRANDMEMBER` 只能用于集合键, 等等; 另外一些命令, 比如 `DEL`、 `TTL` 和 `TYPE`, 可以用于任何类型的键；但是要正确实现这些命令, 必须为不同类型的键设置不同的处理方式: 比如说, 删除一个列表键和删除一个字符串键的操作过程就不太一样。

以上的描述说明, **Redis 必须让每个键都带有类型信息, 使得程序可以检查键的类型, 并为它选择合适的处理方式**.

比如说， 集合类型就可以由字典和整数集合两种不同的数据结构实现， 但是， 当用户执行 ZADD 命令时， 他 / 她应该不必关心集合使用的是什么编码， 只要 Redis 能按照 ZADD 命令的指示， 将新元素添加到集合就可以了。

这说明, **操作数据类型的命令除了要对键的类型进行检查之外, 还需要根据数据类型的不同编码进行多态处理**.

为了解决以上问题, **Redis 构建了自己的类型系统**, 这个系统的主要功能包括:

*   redisObject 对象.
*   基于 redisObject 对象的类型检查.
*   基于 redisObject 对象的显式多态函数.
*   对 redisObject 进行分配、共享和销毁的机制.

[#](#redisobject数据结构) redisObject 数据结构
--------------------------------------

redisObject 是 Redis 类型系统的核心, 数据库中的每个键、值, 以及 Redis 本身处理的参数, 都表示为这种数据类型.

```
typedef struct redisObject {

    
    unsigned type:4;

    
    unsigned encoding:4;

    
    unsigned lru:LRU_BITS; 

    
    int refcount;

    
    void *ptr;

} robj;


```

下图对应上面的结构

![](https://www.pdai.tech/images/db/redis/db-redis-object-1.png)

**其中 type、encoding 和 ptr 是最重要的三个属性**。

*   **type 记录了对象所保存的值的类型**，它的值可能是以下常量中的一个：

```
#define OBJ_STRING 0 
#define OBJ_LIST 1 
#define OBJ_SET 2 
#define OBJ_ZSET 3 
#define OBJ_HASH 4 


```

*   **encoding 记录了对象所保存的值的编码**，它的值可能是以下常量中的一个：

```
#define OBJ_ENCODING_RAW 0     
#define OBJ_ENCODING_INT 1     
#define OBJ_ENCODING_HT 2      
#define OBJ_ENCODING_ZIPMAP 3  
#define OBJ_ENCODING_LINKEDLIST 4 
#define OBJ_ENCODING_ZIPLIST 5 
#define OBJ_ENCODING_INTSET 6  
#define OBJ_ENCODING_SKIPLIST 7  
#define OBJ_ENCODING_EMBSTR 8  
#define OBJ_ENCODING_QUICKLIST 9 
#define OBJ_ENCODING_STREAM 10 


```

*   **ptr 是一个指针，指向实际保存值的数据结构**，这个数据结构由 type 和 encoding 属性决定。举个例子， 如果一个 redisObject 的 type 属性为`OBJ_LIST` ， encoding 属性为`OBJ_ENCODING_QUICKLIST` ，那么这个对象就是一个 Redis 列表（List)，它的值保存在一个 QuickList 的数据结构内，而 ptr 指针就指向 quicklist 的对象；

下图展示了 redisObject 、Redis 所有数据类型、Redis 所有编码方式以及底层数据结构之间的关系（pdai：从 6.0 版本中梳理而来）：

![](https://www.pdai.tech/images/db/redis/db-redis-object-2-2.png)

> 注意：`OBJ_ENCODING_ZIPMAP`没有出现在图中，因为在 redis2.6 开始，它不再是任何数据类型的底层结构 (虽然还有 zipmap.c 的代码); `OBJ_ENCODING_LINKEDLIST`也不支持了，相关代码也删除了。

*   **lru 属性: 记录了对象最后一次被命令程序访问的时间**

**空转时长**：当前时间减去键的值对象的 lru 时间，就是该键的空转时长。Object idletime 命令可以打印出给定键的空转时长

如果服务器打开了 maxmemory 选项，并且服务器用于回收内存的算法为 volatile-lru 或者 allkeys-lru，那么当服务器占用的内存数超过了 maxmemory 选项所设置的上限值时，空转时长较高的那部分键会优先被服务器释放，从而回收内存。

[#](#命令的类型检查和多态) 命令的类型检查和多态
---------------------------

> 那么 Redis 是如何处理一条命令的呢？

**当执行一个处理数据类型命令的时候，redis 执行以下步骤**

*   根据给定的 key，在数据库字典中查找和他相对应的 redisObject，如果没找到，就返回 NULL；
*   检查 redisObject 的 type 属性和执行命令所需的类型是否相符，如果不相符，返回类型错误；
*   根据 redisObject 的 encoding 属性所指定的编码，选择合适的操作函数来处理底层的数据结构；
*   返回数据结构的操作结果作为命令的返回值。

比如现在执行 LPOP 命令：

![](https://www.pdai.tech/images/db/redis/db-redis-object-3.png)

[#](#对象共享) 对象共享
---------------

> redis 一般会把一些常见的值放到一个共享对象中，这样可使程序避免了重复分配的麻烦，也节约了一些 CPU 时间。

**redis 预分配的值对象如下**：

*   各种命令的返回值，比如成功时返回的 OK，错误时返回的 ERROR，命令入队事务时返回的 QUEUE，等等
*   包括 0 在内，小于 REDIS_SHARED_INTEGERS 的所有整数（REDIS_SHARED_INTEGERS 的默认值是 10000）

![](https://www.pdai.tech/images/db/redis/db-redis-object-4.png)

> 注意：共享对象只能被字典和双向链表这类能带有指针的数据结构使用。像整数集合和压缩列表这些只能保存字符串、整数等自勉之的内存数据结构

**为什么 redis 不共享列表对象、哈希对象、集合对象、有序集合对象，只共享字符串对象**？

*   列表对象、哈希对象、集合对象、有序集合对象，本身可以包含字符串对象，复杂度较高。
*   如果共享对象是保存字符串对象，那么验证操作的复杂度为 O(1)
*   如果共享对象是保存字符串值的字符串对象，那么验证操作的复杂度为 O(N)
*   如果共享对象是包含多个值的对象，其中值本身又是字符串对象，即其它对象中嵌套了字符串对象，比如列表对象、哈希对象，那么验证操作的复杂度将会是 O(N 的平方)

如果对复杂度较高的对象创建共享对象，需要消耗很大的 CPU，用这种消耗去换取内存空间，是不合适的

[#](#引用计数以及对象的消毁) 引用计数以及对象的消毁
-----------------------------

> redisObject 中有 refcount 属性，是对象的引用计数，显然计数 0 那么就是可以回收。

*   每个 redisObject 结构都带有一个 refcount 属性，指示这个对象被引用了多少次；
*   当新创建一个对象时，它的 refcount 属性被设置为 1；
*   当对一个对象进行共享时，redis 将这个对象的 refcount 加一；
*   当使用完一个对象后，或者消除对一个对象的引用之后，程序将对象的 refcount 减一；
*   当对象的 refcount 降至 0 时，这个 RedisObject 结构，以及它引用的数据结构的内存都会被释放。

[#](#小结) 小结
-----------

*   redis 使用自己实现的对象机制（redisObject) 来实现类型判断、命令多态和基于引用次数的垃圾回收；
*   redis 会预分配一些常用的数据对象，并通过共享这些对象来减少内存占用，和避免频繁的为小对象分配内存。

[#](#参考文章) 参考文章
---------------

*   https://www.cnblogs.com/gaopengfirst/p/10072680.html
*   https://www.cnblogs.com/neooelric/p/9621736.html
*   https://juejin.cn/post/6844904192042074126