package main

import (
	"basic/common"
	"basic/data"
	"basic/table"
	"github.com/zhangdapeng520/zdpgo_mysql"
)

/*
@Time : 2022/5/20 19:56
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description:
*/

func getMysql() *zdpgo_mysql.Mysql {
	m := zdpgo_mysql.New(&zdpgo_mysql.Config{
		Debug:    true,
		Host:     "ubuntu20_server",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	})
	return m
}

func main() {
	m := getMysql()
	common.IsHealth(m)    // 测试健康状态
	table.AddTable(m)     // 添加表格
	table.FindAllTable(m) // 查找所有表格

	data.Add(m)                // 添加数据
	data.FindById(m)           // 查找数据
	data.FindByIdToStruct(m)   // 查找数据
	data.UpdateById(m)         // 更新数据
	data.FindById(m)           // 查找数据
	data.FindByIdToStruct(m)   // 查找数据
	data.UpdateByIds(m)        // 批量更新数据
	data.FindByIds(m)          // 查找数据
	data.FindByPage(m)         // 查找数据
	data.FindByIdsToStruct(m)  // 查找数据
	data.FindByPageToStruct(m) // 查找数据
	data.DeleteById(m)         // 根据ID删除
	data.DeleteByIds(m)        // 根据ID列表删除
	data.FindByIds(m)          // 查找数据
	data.FindByPage(m)         // 查找数据
	data.FindByIdsToStruct(m)  // 查找数据
	data.FindByPageToStruct(m) // 查找数据
	data.AddMany(m)            // 批量添加
	data.FindByPage(m)         // 查找数据
	data.FindByPageToStruct(m) // 查找数据

	table.DeleteTable(m)  // 删除表格
	table.FindAllTable(m) // 查找所有表格
}
