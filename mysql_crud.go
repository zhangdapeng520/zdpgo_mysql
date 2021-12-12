package zdpgo_mysql

import (
	"database/sql"
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

// 根据ID修改数据
func (mysql *Mysql) Update(sql string, args ...interface{})bool{
	ret := mysql.Execute(sql, args...)
	rows, err := ret.RowsAffected()
	if err !=nil{
		mysql.Logger.Error("获取受影响的行数失败：", err)
		return false
	}
	return rows > 0
}

// 查询单条数据
func (mysql *Mysql) Find(table string, columns []string, id int)*sql.Row{
	columnsStr := strings.Join(columns, ",")
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?;", columnsStr, table)
	row:=mysql.QueryRow(sql, id)
	return row
}

// 根据ID列表查询数据
func (mysql *Mysql) FindIds(talbe string, columns []string, ids []int) *sql.Rows{

	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(v)
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 整理字段列表
	columnsStr := strings.Join(columns, ",")

	// 执行SQL语句
	sql:= fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s);",
		columnsStr, talbe, idsStr)	

	// 执行查询
	rows, err := mysql.Query(sql)
	if err != nil{
		mysql.Logger.Error("根据ID列表查询多条数据失败：", err)
		return nil
	}
	return rows
}

// 分页查询数据
func (mysql *Mysql) FindPages(table string, columns []string, page, size int) *sql.Rows{
	// 整理字段列表
	columnsStr := strings.Join(columns, ",")

	// 计算偏移量
	offset := (page -1) *size
	sql := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d,%d;", columnsStr, table, offset, size)
	
	// 执行查询
	rows, err := mysql.Db.Query(sql)
	if err != nil{
		mysql.Logger.Error("根据ID列表查询多条数据失败：", err)
		return nil
	}
	return rows
}