package zdpgo_mysql

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
@Time : 2022/5/20 21:59
@Author : 张大鹏
@File : update
@Software: Goland2021.3.1
@Description: update更新相关方法
*/

// UpdateById 根据ID修改数据
func (m *Mysql) UpdateById(table string, columns []string, values []interface{}, id int64) (updated int64, err error) {
	// 异常情况
	if table == "" {
		err = errors.New("table不能为空")
		m.Log.Error(err.Error())
		return
	}
	if columns == nil {
		err = errors.New("columns不能为空")
		m.Log.Error(err.Error())
		return
	}
	if values == nil {
		err = errors.New("values不能为空")
		m.Log.Error(err.Error())
		return
	}
	if len(columns) != len(values) {
		err = errors.New("columns和values长度不一致，无法映射")
		m.Log.Error(err.Error())
		return
	}

	// 整理条件
	var data []string
	var args []interface{}
	for i := 0; i < len(columns); i++ {
		v := fmt.Sprintf("%s=?", columns[i])
		args = append(args, values[i])
		data = append(data, v)
	}
	args = append(args, id)
	dataStr := strings.Join(data, ",")

	// 整理SQL语句
	s := fmt.Sprintf("update %s set %s where id=?;", table, dataStr)

	// 执行更新
	updated, err = m.update(s, args...)
	if err != nil {
		m.Log.Error("更新数据失败", "error", err, "sql", s)
		return
	}

	// 正常返回
	return
}

// UpdateByIds 根据ID列表修改数据
func (m *Mysql) UpdateByIds(table string, columns []string, values []interface{}, ids []int64) (updated int64, err error) {
	// 异常情况
	if table == "" {
		err = errors.New("table不能为空")
		m.Log.Error(err.Error())
		return
	}
	if columns == nil {
		err = errors.New("columns不能为空")
		m.Log.Error(err.Error())
		return
	}
	if values == nil {
		err = errors.New("values不能为空")
		m.Log.Error(err.Error())
		return
	}
	if len(columns) != len(values) {
		err = errors.New("columns和values长度不一致，无法映射")
		m.Log.Error(err.Error())
		return
	}
	if ids == nil || len(ids) < 1 {
		err = errors.New("ids不能为空")
		m.Log.Error(err.Error())
		return
	}

	// 整理要更新的数据
	var data []string
	var args []interface{}
	for i := 0; i < len(columns); i++ {
		v := fmt.Sprintf("%s=?", columns[i])
		args = append(args, values[i])
		data = append(data, v)
	}
	dataStr := strings.Join(data, ",")

	// 整理条件
	var idsArr []string
	for i := 0; i < len(ids); i++ {
		idsArr = append(idsArr, strconv.Itoa(int(ids[i])))
	}
	idsStr := strings.Join(idsArr, ",")

	// 整理SQL语句
	s := fmt.Sprintf("update %s set %s where id in (%s);", table, dataStr, idsStr)

	// 执行更新
	updated, err = m.update(s, args...)
	if err != nil {
		m.Log.Error("批量更新数据失败", "error", err, "sql", s)
		return
	}

	// 正常返回
	return
}

func (m *Mysql) update(sql string, args ...interface{}) (updated int64, err error) {
	// 执行SQL语句
	ret, err := m.Execute(sql, args...)
	if err != nil {
		m.Log.Error("执行SQL语句失败", "error", err, "sql", sql, "args", args)
		return
	}

	// 处理执行结果
	updated, err = ret.RowsAffected()
	if err != nil {
		return 0, err
	}
	if updated <= 0 {
		err = errors.New("受影响的行数为0，更新数据失败")
		m.Log.Error(err.Error())
		return
	}

	// 正常返回
	return
}
