> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 [www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/sqlx/#autoid-0-4-1)

> 李文周的 Blog database/sql mysql sqlx 批量插入 bulk insert query in select in order by find_in_set batch inser......

在项目中我们通常可能会使用`database/sql`连接 MySQL 数据库。本文借助使用`sqlx`实现批量插入数据的例子，介绍了`sqlx`中可能被你忽视了的`sqlx.In`和`DB.NamedExec`方法。

sqlx 介绍
-------

在项目中我们通常可能会使用`database/sql`连接 MySQL 数据库。`sqlx`可以认为是 Go 语言内置`database/sql`的超集，它在优秀的内置`database/sql`基础上提供了一组扩展。这些扩展中除了大家常用来查询的`Get(dest interface{}, ...) error`和`Select(dest interface{}, ...) error`外还有很多其他强大的功能。

安装 sqlx
-------

```
go get github.com/jmoiron/sqlx


```

基本使用
----

### 连接数据库

```
var db *sqlx.DB

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}


```

### 查询

查询单行数据示例代码如下：

```
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}


```

查询多行数据示例代码如下：

```
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}


```

### 插入、更新和删除

sqlx 中的 exec 方法与原生 sql 中的 exec 使用基本一致：

```
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() 
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}


func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() 
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}


func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() 
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}


```

### NamedExec

`DB.NamedExec`方法用来绑定 SQL 语句与结构体或 map 中的同名字段。

```
func insertUserDemo()(err error){
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "七米",
			"age": 28,
		})
	return
}


```

### NamedQuery

与`DB.NamedExec`同理，这里是支持查询。

```
func namedQuery(){
	sqlStr := "SELECT * FROM user WHERE 
	
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{
		Name: "七米",
	}
	
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}


```

### 事务操作

对于事务操作，我们可以使用`sqlx`中提供的`db.Beginx()`和`tx.Exec()`方法。示例代码如下：

```
func transactionDemo2()(err error) {
	tx, err := db.Beginx() 
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) 
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() 
		} else {
			err = tx.Commit() 
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id=?"

	rs, err := tx.Exec(sqlStr1, 1)
	if err!= nil{
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err!=nil{
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}


```

sqlx.In
-------

`sqlx.In`是`sqlx`提供的一个非常方便的函数。

### sqlx.In 的批量插入示例

#### 表结构

为了方便演示插入数据操作，这里创建一个`user`表，表结构如下：

```
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


```

#### 结构体

定义一个`user`结构体，字段通过 tag 与数据库中 user 表的列一致。

```
type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}


```

#### bindvars（绑定变量）

查询占位符`?`在内部称为 **bindvars（查询占位符）**, 它非常重要。你应该始终使用它们向数据库发送值，因为它们可以防止 SQL 注入攻击。`database/sql`不尝试对查询文本进行任何验证；它与编码的参数一起按原样发送到服务器。除非驱动程序实现一个特殊的接口，否则在执行之前，查询是在服务器上准备的。因此`bindvars`是特定于数据库的:

*   MySQL 中使用`?`
*   PostgreSQL 使用枚举的`$1`、`$2`等 bindvar 语法
*   SQLite 中`?`和`$1`的语法都支持
*   Oracle 中使用`:name`的语法

`bindvars`的一个常见误解是，它们用来在 sql 语句中插入值。它们其实仅用于参数化，不允许更改 SQL 语句的结构。例如，使用`bindvars`尝试参数化列或表名将不起作用：

```
db.Query("SELECT * FROM ?", "mytable")
 

db.Query("SELECT ?, ? FROM people", "name", "location")


```

#### 自己拼接语句实现批量插入

比较笨，但是很好理解。就是有多少个 User 就拼接多少个`(?, ?)`。

```
func BatchInsertUsers(users []*User) error {
	
	valueStrings := make([]string, 0, len(users))
	
	valueArgs := make([]interface{}, 0, len(users) * 2)
	
	for _, u := range users {
		
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := DB.Exec(stmt, valueArgs...)
	return err
}


```

#### 使用 sqlx.In 实现批量插入

前提是需要我们的结构体实现`driver.Valuer`接口：

```
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}


```

使用`sqlx.In`实现批量插入代码如下：

```
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",  // 三个(?)代表有三个(name, age)要写入
		users..., 
	)
	fmt.Println(query) 
	fmt.Println(args)  
	_, err := DB.Exec(query, args...)
	return err
}


```

#### 使用 NamedExec 实现批量插入

**注意** ：该功能需 1.3.1 版本以上，并且 1.3.1 版本目前还有点问题，sql 语句最后不能有空格和`;`，详见 [issues/690](https://github.com/jmoiron/sqlx/issues/690)。

使用`NamedExec`实现批量插入的代码如下：

```
func BatchInsertUsers3(users []*User) error {
	_, err := DB.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}


```

把上面三种方法综合起来试一下：

```
func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	u1 := User{Name: "七米", Age: 18}
	u2 := User{Name: "q1mi", Age: 28}
	u3 := User{Name: "小王子", Age: 38}

	
	users := []*User{&u1, &u2, &u3}
	err = BatchInsertUsers(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	}

	
	users2 := []interface{}{u1, u2, u3}
	err = BatchInsertUsers2(users2)
	if err != nil {
		fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	}

	
	users3 := []*User{&u1, &u2, &u3}
	err = BatchInsertUsers3(users3)
	if err != nil {
		fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	}
}


```

### sqlx.In 的查询示例

关于`sqlx.In`这里再补充一个用法，在`sqlx`查询语句中实现 In 查询和 FIND_IN_SET 函数。即实现`SELECT * FROM user WHERE id in (3, 2, 1);`和`SELECT * FROM user WHERE id in (3, 2, 1) ORDER BY FIND_IN_SET(id, '3,2,1');`。

#### in 查询

查询 id 在给定 id 集合中的数据。**返回的slice中为给定的数据会初始化为零值**，如下面返回的数据结果是[{0 30 Jadon} {0 12 haha}]

```
func QueryByIDs(ids []int)(users []User, err error){
	
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	// 相当于根据传入的ids生成相应的sql语句
	// SELECT name, age FROM user WHERE id IN (?, ?, ...)  // 其中？的数量为ids的长度
	if err != nil {
		return
	}
	
	query = DB.Rebind(query)

	err = DB.Select(&users, query, args...)
	return
}


```

#### in 查询和 FIND_IN_SET 函数

查询 id 在给定 id 集合的数据并维持给定 id 集合的顺序。

```
func QueryAndOrderByIDs(ids []int)(users []User, err error){
	
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	
	query = DB.Rebind(query)

	err = DB.Select(&users, query, args...)
	return
}


```

当然，在这个例子里面你也可以先使用`IN`查询，然后通过代码按给定的 ids 对查询结果进行排序。

参考链接：

[Illustrated guide to SQLX](http://jmoiron.github.io/sqlx/)

![](https://www.liwenzhou.com/images/wxgzh_qrcode.png)