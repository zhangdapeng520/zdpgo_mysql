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
