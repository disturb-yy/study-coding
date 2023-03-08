# docker容器连接另一个容器

```go
version: "3.7"
services:
  redis5014:
    image: "redis:5.0.14"
    ports:
      - "26379:6379"

// 连接redis，使用cache容器的主机名和端口号6379
rdb := redis.NewClient(&redis.Options{
        Addr:     "redis5014:6379",
        Password: "",
        DB:       0,
})
```

