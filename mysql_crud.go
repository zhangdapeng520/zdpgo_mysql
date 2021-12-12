package zdpgo_mysql

import (
	"fmt"
	"strconv"
	"strings"
)

// 添加单条数据
func (mysql *Mysql) Add(sql string, args ...interface{}) int64 {
	ret := mysql.Execute(sql, args...)
	uid, err := ret.LastInsertId() // 获取新插入的数据的ID
	if err != nil {
		mysql.Logger.Error("获取新插入的数据ID失败：", err)
		return -1
	}
	return uid
}

// 根据ID删除数据
func (mysql *Mysql) Delete(table string, id int) bool{
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", table, id)
	ret := mysql.Execute(sql)
	rows, err := ret.RowsAffected()
	if err!=nil{
		mysql.Logger.Error("获取受影响的行数失败：", err)
	}
	return rows > 0
}

// 根据ID列表删除
func (mysql *Mysql)DeleteIds(table string, ids ...int) bool {
	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(v)
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 执行SQL语句
	sql:= fmt.Sprintf("DELETE FROM %s WHERE id IN (%s);",
		table, idsStr)	
	mysql.Logger.Info("执行批量删除的SQL语句：", sql)
	ret := mysql.Execute(sql)

	rows, err := ret.RowsAffected()
	if err!=nil{
		mysql.Logger.Error("获取受影响的行数失败：", err)
	}
	return rows > 0
}