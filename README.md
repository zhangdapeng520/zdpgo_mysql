# zdpgo_mysql

Golang操作MySQL的快捷工具库

项目地址：https://github.com/zhangdapeng520/zdpgo_mysql

## 版本历史

- v1.4.0 2022/02/7
- v1.5.0 2022/02/9
- v1.6.0 2022/02/18 相关bug修复，日志库升级
- v1.7.0 2022/02/18 增加事务相关方法
- v1.7.1 2022/04/7 移除第三方依赖
- v1.7.2 2022/04/8 代码结构优化
- v1.7.3 2022/05/20 优化：代码架构优化
- v1.7.4 2022/05/24 优化：代码架构优化
- v1.7.5 2022/07/16 优化：代码优化

## 使用示例

## 添加表和删除表

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_mysql"
)

/*
@Time : 2022/7/16 17:02
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 添加表
*/

func main() {
	m := zdpgo_mysql.NewWithConfig(&zdpgo_mysql.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	})

	// 检查是否能够连接
	fmt.Println(m.IsHealth())

	// 删除表
	err := m.DeleteTable("students")
	if err != nil {
		panic(err)
	}

	// 添加表
	sql := `
create table IF NOT EXISTS students
(
    id     int primary key auto_increment not null,
    name   varchar(255)                   not null,
    age    smallint                       not null,
    gender varchar(6) default '男'
) charset = utf8mb4;
`
	err = m.AddTable(sql)
	if err != nil {
		panic(err)
	}

	// 查找所有表
	tables, err := m.FindAllTable()
	if err != nil {
		panic(err)
	}
	fmt.Println(tables)

	// 删除表
	err = m.DeleteTable("students")
	if err != nil {
		panic(err)
	}
}
```

### 添加数据

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_mysql"
)

/*
@Time : 2022/7/16 17:02
@Author : 张大鹏
@File : main.gop
@Software: Goland2021.3.1
@Description: 添加数据
*/

func main() {
	m := zdpgo_mysql.NewWithConfig(&zdpgo_mysql.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	})

	// 删除表
	err := m.DeleteTable("students")
	if err != nil {
		panic(err)
	}

	// 添加表
	sql := `
create table IF NOT EXISTS students
(
    id     int primary key auto_increment not null,
    name   varchar(255)                   not null,
    age    smallint                       not null,
    gender varchar(6) default '男'
) charset = utf8mb4;
`
	err = m.AddTable(sql)
	if err != nil {
		panic(err)
	}

	// 添加一条数据
	table := "students"
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"张大鹏", 22, "男"}
	_, err = m.Add(table, columns, values)
	if err != nil {
		panic(err)
	}

	// 添加一条数据
	table = "students"
	columns = []string{"name", "age", "gender"}
	values = []interface{}{"孙悟空", 22, "男"}
	_, err = m.Add(table, columns, values)
	if err != nil {
		panic(err)
	}

	// 添加一条数据
	table = "students"
	columns = []string{"name", "age", "gender"}
	values = []interface{}{"白骨精", 22, "女"}
	_, err = m.Add(table, columns, values)
	if err != nil {
		panic(err)
	}

	// 查询所有数据
	table = "students"
	columns = []string{"name", "age", "gender"}
	var (
		name   string
		age    int
		gender string
	)

	// 查询一条数据
	rows, err := m.FindByPage(table, columns, 1, 20)
	if err != nil {
		panic(err)
	}

	// 迭代取出数据
	for rows.Next() {
		err = rows.Scan(&name, &age, &gender)
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("姓名：", name)
		fmt.Println("年龄：", age)
		fmt.Println("性别：", gender)
		fmt.Println("=======================")
	}

	// 删除表
	err = m.DeleteTable("students")
	if err != nil {
		panic(err)
	}
}
```