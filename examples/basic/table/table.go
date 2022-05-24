package table

import "github.com/zhangdapeng520/zdpgo_mysql"

/*
@Time : 2022/5/20 20:19
@Author : 张大鹏
@File : table
@Software: Goland2021.3.1
@Description: table表格相关
*/

func FindAllTable(m *zdpgo_mysql.Mysql) {
	tables, err := m.FindAllTable()
	if err != nil {
		m.Log.Panic("FindAll查找所有表格失败", "error", err)
	}
	m.Log.Debug("查找所有表格成功", "tables", tables)
}

func AddTable(m *zdpgo_mysql.Mysql) {
	sql := `
create table IF NOT EXISTS students
(
    id     int primary key auto_increment not null,
    name   varchar(255)                   not null,
    age    smallint                       not null,
    gender varchar(6) default '男'
) charset = utf8mb4;
`
	err := m.AddTable(sql)
	if err != nil {
		m.Log.Panic("创建表格失败", "error", err)
		return
	}
	m.Log.Debug("添加表格成功")
}

func DeleteTable(m *zdpgo_mysql.Mysql) {
	err := m.DeleteTable("students")
	if err != nil {
		m.Log.Panic("删除表格失败", "error", err)
	}
	m.Log.Debug("删除表格成功")
}
