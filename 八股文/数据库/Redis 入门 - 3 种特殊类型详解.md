> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.pdai.tech](https://www.pdai.tech/md/db/nosql-redis/db-redis-data-type-special.html)

> Redis 除了上文中 5 种基础数据类型，还有三种特殊的数据类型，分别是 HyperLogLogs（基数统计）， Bitmaps (位图) 和 geospatial（地理位置）

> Redis 除了上文中 5 种基础数据类型，还有三种特殊的数据类型，分别是 **HyperLogLogs**（基数统计）， **Bitmaps** (位图) 和 **geospatial** （地理位置）。@pdai

*   [Redis 入门 - 数据类型：3 种特殊类型详解](#redis%E5%85%A5%E9%97%A8---%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B3%E7%A7%8D%E7%89%B9%E6%AE%8A%E7%B1%BB%E5%9E%8B%E8%AF%A6%E8%A7%A3)
    *   [HyperLogLogs（基数统计）](#hyperloglogs%E5%9F%BA%E6%95%B0%E7%BB%9F%E8%AE%A1)
    *   [Bitmap （位存储）](#bitmap-%E4%BD%8D%E5%AD%98%E5%82%A8)
    *   [geospatial (地理位置)](#geospatial-%E5%9C%B0%E7%90%86%E4%BD%8D%E7%BD%AE)
        *   [geoadd](#geoadd)
        *   [geopos](#geopos)
        *   [geodist](#geodist)
        *   [georadius](#georadius)
        *   [georadiusbymember](#georadiusbymember)
        *   [geohash(较少使用)](#geohash%E8%BE%83%E5%B0%91%E4%BD%BF%E7%94%A8)
        *   [底层](#%E5%BA%95%E5%B1%82)
    *   [参考文章](#%E5%8F%82%E8%80%83%E6%96%87%E7%AB%A0)

[#](#hyperloglogs-基数统计) HyperLogLogs（基数统计）
------------------------------------------

> Redis 2.8.9 版本更新了 Hyperloglog 数据结构！

*   **什么是基数？**

举个例子，A = {1, 2, 3, 4, 5}， B = {3, 5, 6, 7, 9}；那么基数（不重复的元素）= 1, 2, 4, 6, 7, 9； （允许容错，即可以接受一定误差）

*   **HyperLogLogs 基数统计用来解决什么问题**？

这个结构可以非常省内存的去统计各种计数，比如注册 IP 数、每日访问 IP 数、页面实时 UV、在线用户数，共同好友数等。

*   **它的优势体现在哪**？

一个大型的网站，每天 IP 比如有 100 万，粗算一个 IP 消耗 15 字节，那么 100 万个 IP 就是 15M。而 HyperLogLog 在 Redis 中每个键占用的内容都是 12K，理论存储近似接近 2^64 个值，不管存储的内容是什么，它一个基于基数估算的算法，只能比较准确的估算出基数，可以使用少量固定的内存去存储并识别集合中的唯一元素。而且这个估算的基数并不一定准确，是一个带有 0.81% 标准错误的近似值（对于可以接受一定容错的业务场景，比如 IP 数统计，UV 等，是可以忽略不计的）。

*   **相关命令使用**

```
127.0.0.1:6379> pfadd key1 a b c d e f g h i	
(integer) 1
127.0.0.1:6379> pfcount key1					
(integer) 9
127.0.0.1:6379> pfadd key2 c j k l m e g a		
(integer) 1
127.0.0.1:6379> pfcount key2
(integer) 8
127.0.0.1:6379> pfmerge key3 key1 key2			
OK
127.0.0.1:6379> pfcount key3
(integer) 13


```

[#](#bitmap-位存储) Bitmap （位存储）
-----------------------------

> Bitmap 即位图数据结构，都是操作二进制位来进行记录，只有 0 和 1 两个状态。

*   **用来解决什么问题**？

比如：统计用户信息，活跃，不活跃！ 登录，未登录！ 打卡，不打卡！ **两个状态的，都可以使用 Bitmaps**！

如果存储一年的打卡状态需要多少内存呢？ 365 天 = 365 bit 1 字节 = 8bit 46 个字节左右！

*   **相关命令使用**

使用 bitmap 来记录 周一到周日的打卡！ 周一：1 周二：0 周三：0 周四：1 ......

```
127.0.0.1:6379> setbit sign 0 1
(integer) 0
127.0.0.1:6379> setbit sign 1 1
(integer) 0
127.0.0.1:6379> setbit sign 2 0
(integer) 0
127.0.0.1:6379> setbit sign 3 1
(integer) 0
127.0.0.1:6379> setbit sign 4 0
(integer) 0
127.0.0.1:6379> setbit sign 5 0
(integer) 0
127.0.0.1:6379> setbit sign 6 1
(integer) 0


```

查看某一天是否有打卡！

```
127.0.0.1:6379> getbit sign 3
(integer) 1
127.0.0.1:6379> getbit sign 5
(integer) 0


```

统计操作，统计 打卡的天数！

```
127.0.0.1:6379> bitcount sign 
(integer) 3


```

[#](#geospatial-地理位置) geospatial (地理位置)
---------------------------------------

> Redis 的 Geo 在 Redis 3.2 版本就推出了! 这个功能可以推算地理位置的信息: 两地之间的距离, 方圆几里的人

### [#](#geoadd) geoadd

> 添加地理位置

```
127.0.0.1:6379> geoadd china:city 118.76 32.04 nanjing 112.55 37.86 taiyuan 123.43 41.80 shenyang
(integer) 3
127.0.0.1:6379> geoadd china:city 144.05 22.52 shengzhen 120.16 30.24 hangzhou 108.96 34.26 xian
(integer) 3


```

**规则**

两级无法直接添加，我们一般会下载城市数据 (这个网址可以查询 GEO： http://www.jsons.cn/lngcode)！

*   有效的经度从 - 180 度到 180 度。
*   有效的纬度从 - 85.05112878 度到 85.05112878 度。

```
127.0.0.1:6379> geoadd china:city 39.90 116.40 beijin
(error) ERR invalid longitude,latitude pair 39.900000,116.400000


```

### [#](#geopos) geopos

> 获取指定的成员的经度和纬度

```
127.0.0.1:6379> geopos china:city taiyuan manjing
1) 1) "112.54999905824661255"
   1) "37.86000073876942196"
2) 1) "118.75999957323074341"
   1) "32.03999960287850968"


```

获得当前定位, 一定是一个坐标值!

### [#](#geodist) geodist

> 如果不存在, 返回空

单位如下

*   m
*   km
*   mi 英里
*   ft 英尺

```
127.0.0.1:6379> geodist china:city taiyuan shenyang m
"1026439.1070"
127.0.0.1:6379> geodist china:city taiyuan shenyang km
"1026.4391"


```

### [#](#georadius) georadius

> 附近的人 ==> 获得所有附近的人的地址, 定位, 通过半径来查询

获得**与指定坐标**一定半径范围内的其他人

```
127.0.0.1:6379> georadius china:city 110 30 1000 km			以 100,30 这个坐标为中心, 寻找半径为1000km的城市
1) "xian"
2) "hangzhou"
3) "manjing"
4) "taiyuan"
127.0.0.1:6379> georadius china:city 110 30 500 km
1) "xian"
127.0.0.1:6379> georadius china:city 110 30 500 km withdist
1) 1) "xian"
   2) "483.8340"
127.0.0.1:6379> georadius china:city 110 30 1000 km withcoord withdist count 2
1) 1) "xian"
   2) "483.8340"
   3) 1) "108.96000176668167114"
      2) "34.25999964418929977"
2) 1) "manjing"
   2) "864.9816"
   3) 1) "118.75999957323074341"
      2) "32.03999960287850968"


```

参数 key 经度 纬度 半径 单位 [显示结果的经度和纬度] [显示结果的距离] [显示的结果的数量]

### [#](#georadiusbymember) georadiusbymember

> 显示与**指定成员一定半径范围**内的其他成员

```
127.0.0.1:6379> georadiusbymember china:city taiyuan 1000 km
1) "manjing"
2) "taiyuan"
3) "xian"
127.0.0.1:6379> georadiusbymember china:city taiyuan 1000 km withcoord withdist count 2
1) 1) "taiyuan"
   2) "0.0000"
   3) 1) "112.54999905824661255"
      2) "37.86000073876942196"
2) 1) "xian"
   2) "514.2264"
   3) 1) "108.96000176668167114"
      2) "34.25999964418929977"


```

参数与 georadius 一样

### [#](#geohash-较少使用) geohash(较少使用)

> 该命令返回 11 个字符的 hash 字符串

```
127.0.0.1:6379> geohash china:city taiyuan shenyang
1) "ww8p3hhqmp0"
2) "wxrvb9qyxk0"


```

将二维的经纬度转换为一维的字符串, 如果两个字符串越接近, 则距离越近

### [#](#底层) 底层

> geo 底层的实现原理实际上就是 Zset, 我们可以通过 Zset 命令来操作 geo

```
127.0.0.1:6379> type china:city
zset


```

查看全部元素 删除指定的元素

```
127.0.0.1:6379> zrange china:city 0 -1 withscores
 1) "xian"
 2) "4040115445396757"
 3) "hangzhou"
 4) "4054133997236782"
 5) "manjing"
 6) "4066006694128997"
 7) "taiyuan"
 8) "4068216047500484"
 9) "shenyang"
1)  "4072519231994779"
2)  "shengzhen"
3)  "4154606886655324"
127.0.0.1:6379> zrem china:city manjing
(integer) 1
127.0.0.1:6379> zrange china:city 0 -1
1) "xian"
2) "hangzhou"
3) "taiyuan"
4) "shenyang"
5) "shengzhen"


```

[#](#参考文章) 参考文章
---------------

*   http://www.jsons.cn/lngcode
*   https://www.cnblogs.com/junlinsky/p/13528452.html
*   https://www.cnblogs.com/touyel/p/12728096.html
*   https://www.cnblogs.com/junlinsky/p/13528452.html
*   https://www.cnblogs.com/wang-sky/p/13857787.html