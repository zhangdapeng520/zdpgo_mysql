package zdpgo_mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

/*
@Time : 2022/5/20 20:22
@Author : 张大鹏
@File : add
@Software: Goland2021.3.1
@Description: add添加相关的方法
*/

// AddTable 增加表格
func (m *Mysql) AddTable(sql string) error {
	_, err := m.Execute(sql)
	if err != nil {
		m.Log.Error("添加新的表格失败", "error", err)
	}
	return err
}

// Add 添加数据
func (m *Mysql) Add(table string, columns []string, values []interface{}) (id int64, err error) {
	// 处理异常
	if columns == nil {
		err = errors.New("columns字段列表不能为空")
		return
	}
	if values == nil {
		err = errors.New("values值列表不能为空")
		return
	}
	if len(columns) != len(values) {
		err = errors.New("columns的长度和values长度不相等，字段无法映射")
		return
	}

	// 整理字段
	columnStr := strings.Join(columns, ",")
	var valueArr []string
	for i := 0; i < len(values); i++ {
		valueArr = append(valueArr, "?")
	}
	valueStr := strings.Join(valueArr, ",")

	// 整理SQL语句
	s := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s);", table, columnStr, valueStr)

	// 执行添加
	addId, err := m.add(s, values...)
	if err != nil {
		m.Log.Error("添加数据失败", "error", err)
		return
	}

	// 返回添加的id
	return addId, nil
}

// 添加数据
func (m *Mysql) add(sqlStr string, args ...interface{}) (uid int64, err error) {
	var (
		ret sql.Result
	)
	// 执行SQL语句
	ret, err = m.Execute(sqlStr, args...)
	if err != nil {
		m.Log.Error("执行SQL语句失败", "error", err, "sql", sqlStr, "args", args)
		return
	}

	// 执行结果处理
	affected, err := ret.RowsAffected()
	if affected <= 0 {
		msg := "受影响的行数为0，插入数据失败"
		m.Log.Error(msg)
		err = errors.New(msg)
		return
	}
	if err != nil {
		m.Log.Error("获取受影响的行数失败", "error", err)
		return
	}

	// 获取新插入的数据的ID
	uid, err = ret.LastInsertId()

	// 处理错误
	if err != nil {
		m.Log.Error("获取新插入数据的ID失败", "error", err)
		return
	}

	// 正常返回
	return
}

// AddMany 批量添加数据
func (m *Mysql) AddMany(table string, columns []string, values [][]interface{}) (affected int64, err error) {
	// 处理异常
	if columns == nil {
		err = errors.New("columns字段列表不能为空")
		m.Log.Error(err.Error())
		return
	}
	if values == nil {
		err = errors.New("values值列表不能为空")
		m.Log.Error(err.Error())
		return
	}
	if len(columns) != len(values[0]) {
		err = errors.New("columns的长度和values子数组长度不相等，字段无法映射")
		m.Log.Error(err.Error())
		return
	}

	// 整理字段
	columnStr := strings.Join(columns, ",")

	// 整理占位符
	var argsArr []string
	var args []interface{}
	for i := 0; i < len(values); i++ {
		var valueArr []string
		for j := 0; j < len(values[0]); j++ {
			valueArr = append(valueArr, "?")
			args = append(args, values[i][j])
		}
		valueStr := strings.Join(valueArr, ",")
		tempStr := fmt.Sprintf("(%s)", valueStr)
		argsArr = append(argsArr, tempStr)
	}
	argsStr := strings.Join(argsArr, ",")

	// 整理SQL语句
	s := fmt.Sprintf("INSERT INTO %s(%s) VALUES %s;", table, columnStr, argsStr)

	// 执行SQL语句
	ret, err := m.Execute(s, args...)
	if err != nil {
		m.Log.Error("执行SQL语句失败", "error", err, "sql", s)
		return
	}

	// 执行结果处理
	affected, err = ret.RowsAffected()
	if affected <= 0 {
		err = errors.New("受影响的行数为0，批量插入数据失败")
		m.Log.Error(err.Error())
		return
	}
	if err != nil {
		m.Log.Error("获取受影响的行数失败", "error", err)
		return
	}

	// 返回受影响的行数
	return
}
