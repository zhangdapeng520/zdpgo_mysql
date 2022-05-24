package zdpgo_mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_mysql/sqlscan"
	"strconv"
	"strings"
)

/*
@Time : 2022/5/20 8:55
@Author : 张大鹏
@File : find
@Software: Goland2021.3.1
@Description: find 查找相关
*/

// FindAllTable 查询所有表格
func (m *Mysql) FindAllTable() (tables []string, err error) {
	var (
		rows *sql.Rows
	)

	s := "show tables;"
	rows, err = m.Query(s)
	if err != nil {
		m.Log.Error("查询所有表格失败", "error", err)
		return
	}

	// 循环读取数据
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			m.Log.Error("读取表格名称失败", "error", err)
			return
		}
		tables = append(tables, tableName)
	}

	// 返回表格数据
	return
}

// TransRowsToMaps 将rows转换为map
func (m *Mysql) TransRowsToMaps(rows *sql.Rows) (list []map[string]interface{}, err error) {
	// 获取字段列表
	columns, _ := rows.Columns()
	columnLength := len(columns)

	// 临时存储每行数据
	cache := make([]interface{}, columnLength)
	for index, _ := range cache { //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}

	// 转换数据
	for rows.Next() {
		err = rows.Scan(cache...)
		if err != nil {
			return nil, err
		}
		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}

	// 关闭rows
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// FindByPage 分页查询数据
func (m *Mysql) FindByPage(table string, columns []string, page, size int) (rows *sql.Rows, err error) {
	// 整理字段列表
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 计算偏移量
	offset := (page - 1) * size
	s := fmt.Sprintf("SELECT %s FROM %s LIMIT %d,%d;", columnsStr, table, offset, size)

	// 执行查询
	rows, err = m.Db.Query(s)

	// 处理查询错误
	if err != nil {
		m.Log.Error("执行查询失败", "error", err, "s", s)
		return
	}

	// 正常返回
	return
}

// FindByPageToStruct 执行分页查询并映射到结构体
func (m *Mysql) FindByPageToStruct(table string, columns []string, page, size int, objects interface{}) (err error) {
	// 整理字段列表
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 计算偏移量
	offset := (page - 1) * size
	s := fmt.Sprintf("SELECT %s FROM %s LIMIT %d,%d;", columnsStr, table, offset, size)

	// 结构体映射
	ctx := context.Background()
	err = sqlscan.Select(ctx, m.Db, objects, s)
	if err != nil {
		return err
	}

	// 正常返回
	return nil
}

// FindByIds 根据ID列表查询数据
func (m *Mysql) FindByIds(table string, columns []string, ids []int64) (rows *sql.Rows, err error) {
	// 参数校验
	if ids == nil {
		err = errors.New("ids不能为空")
		m.Log.Error(err.Error())
		return
	}

	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.FormatInt(v, 10)
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 整理字段列表
	columnsStr := "*"
	if columns != nil {
		columnsStr = strings.Join(columns, ",")
	}

	// 执行SQL语句
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s) ORDER BY ID DESC;",
		columnsStr, table, idsStr)

	// 执行查询
	rows, err = m.Query(s)

	// 处理查询错误
	if err != nil {
		m.Log.Error("查询数据失败", "error", err, "sql", s)
		return
	}

	// 正常返回
	return
}

// FindByIdsToStruct 根据ID列表查询数据并映射到结构体
func (m *Mysql) FindByIdsToStruct(table string, columns []string, ids []int64, objects interface{}) (err error) {
	// 参数校验
	if ids == nil {
		err = errors.New("ids不能为空")
		m.Log.Error(err.Error())
		return
	}

	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(int(v))
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 整理字段列表
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 执行SQL语句
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s) ORDER BY ID DESC;",
		columnsStr, table, idsStr)

	// 结构体映射
	ctx := context.Background()
	err = sqlscan.Select(ctx, m.Db, objects, s)
	if err != nil {
		m.Log.Error("结构体映射失败", "error", err, "sql", s)
		return
	}

	// 正常返回
	return
}

// FindById 查询单条数据
func (m *Mysql) FindById(table string, columns []string, id int) (row *sql.Row, err error) {
	// 处理字段
	columnsStr := "*"
	if columns != nil {
		columnsStr = strings.Join(columns, ",")
	}

	// 整理SQL语句
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?;", columnsStr, table)

	// 执行SQL语句
	row, err = m.QueryRow(s, id)
	if err != nil {
		m.Log.Error("查询单条数据失败", "error", err, "sql", s, "id", id)
		return
	}

	// 正常返回
	return
}

// FindByIdToStruct 根据id查询数据，并将结果转换为结构体
func (m *Mysql) FindByIdToStruct(table string, columns []string, id int64, objects interface{}) (err error) {
	// 处理字段
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 整理SQL语句
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?;", columnsStr, table)

	// 结构体映射
	ctx := context.Background()
	err = sqlscan.Select(ctx, m.Db, objects, s, id)
	if err != nil {
		m.Log.Error("结构体映射失败", "error", err, "sql", s, "id", id)
		return
	}

	// 正常返回
	return
}
