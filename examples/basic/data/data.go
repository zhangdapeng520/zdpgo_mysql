package data

import (
	"database/sql"
	"github.com/zhangdapeng520/zdpgo_mysql"
)

/*
@Time : 2022/5/20 20:51
@Author : 张大鹏
@File : data
@Software: Goland2021.3.1
@Description: data表格数据相关
*/

// Add 测试添加
func Add(m *zdpgo_mysql.Mysql) {
	var (
		table   string
		columns []string
		values  []interface{}
		uid     int64
		err     error
	)

	// 添加一条数据
	table = "students"
	columns = []string{"name", "age", "gender"}
	values = []interface{}{"张大鹏", 22, "男"}
	uid, err = m.Add(table, columns, values)
	if err != nil {
		m.Log.Panic("添加数据失败", "error", err)
	}
	m.Log.Debug("添加数据成功", "uid", uid)

	// 添加一条数据
	table = "students"
	columns = []string{"name", "age", "gender"}
	values = []interface{}{"孙悟空", 22, "男"}
	uid, err = m.Add(table, columns, values)
	if err != nil {
		m.Log.Panic("添加数据失败", "error", err)
	}
	m.Log.Debug("添加数据成功", "uid", uid)

	// 添加一条数据
	table = "students"
	columns = []string{"name", "age", "gender"}
	values = []interface{}{"白骨精", 22, "女"}
	uid, err = m.Add(table, columns, values)
	if err != nil {
		m.Log.Panic("添加数据失败", "error", err)
	}
	m.Log.Debug("添加数据成功", "uid", uid)

}

// AddMany 批量添加
func AddMany(m *zdpgo_mysql.Mysql) {
	var (
		table    string
		columns  []string
		values   [][]interface{}
		affected int64
		err      error
	)

	// 添加一条数据
	table = "students"
	columns = []string{"name", "age", "gender"}
	values = [][]interface{}{
		{"张三", 22, "男"},
		{"李四", 22, "男"},
		{"王五", 22, "男"},
	}
	affected, err = m.AddMany(table, columns, values)
	if err != nil {
		m.Log.Panic("批量添加数据失败", "error", err)
	}
	m.Log.Debug("批量添加数据成功", "affected", affected)
}

func FindById(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"name", "age", "gender"}
	var (
		name   string
		age    int
		gender string
		row    *sql.Row
		err    error
	)

	// 查询一条数据
	row, err = m.FindById(table, columns, 1)
	if err != nil {
		m.Log.Panic("根据ID查询数据失败", "error", err)
	}
	err = row.Scan(&name, &age, &gender)
	if err != nil {
		if err != nil {
			m.Log.Panic("提取数据失败", "error", err)
		}
	}
	m.Log.Debug("查询数据成功", "name", name, "age", age, "gender", gender)
}

func FindByIdToStruct(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"id", "name", "age", "gender"}
	var (
		err      error
		students []Students
	)

	// 查询一条数据
	err = m.FindByIdToStruct(table, columns, 1, &students)
	if err != nil {
		m.Log.Panic("根据ID查询数据失败", "error", err)
	}
	m.Log.Debug("查询数据成功", "student", students)
}

func FindByIds(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"name", "age", "gender"}
	ids := []int64{1, 2, 3}
	var (
		name   string
		age    int
		gender string
		rows   *sql.Rows
		err    error
	)

	// 查询一条数据
	rows, err = m.FindByIds(table, columns, ids)
	if err != nil {
		m.Log.Panic("根据ID查询数据失败", "error", err)
	}

	// 迭代取出数据
	for rows.Next() {
		err = rows.Scan(&name, &age, &gender)
		if err != nil {
			if err != nil {
				m.Log.Panic("提取数据失败", "error", err)
			}
		}
		m.Log.Debug("查询数据成功", "name", name, "age", age, "gender", gender)
	}
}

func FindByPage(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"name", "age", "gender"}
	var (
		name   string
		age    int
		gender string
		rows   *sql.Rows
		err    error
	)

	// 查询一条数据
	rows, err = m.FindByPage(table, columns, 1, 20)
	if err != nil {
		m.Log.Panic("根据分页查询数据失败", "error", err)
	}

	// 迭代取出数据
	for rows.Next() {
		err = rows.Scan(&name, &age, &gender)
		if err != nil {
			if err != nil {
				m.Log.Panic("提取数据失败", "error", err)
			}
		}
		m.Log.Debug("查询数据成功", "name", name, "age", age, "gender", gender)
	}
}

func FindByIdsToStruct(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"id", "name", "age", "gender"}
	ids := []int64{1, 2, 3}
	var (
		err      error
		students []Students
	)

	// 查询数据
	err = m.FindByIdsToStruct(table, columns, ids, &students)
	if err != nil {
		m.Log.Panic("根据ID查询数据失败", "error", err)
	}

	// 迭代取出数据
	for _, student := range students {
		m.Log.Debug("查询数据成功", "student", student)
	}
}

func FindByPageToStruct(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"id", "name", "age", "gender"}
	var (
		err      error
		students []Students
	)

	// 查询数据
	err = m.FindByPageToStruct(table, columns, 1, 100, &students)
	if err != nil {
		m.Log.Panic("根据ID查询数据失败", "error", err)
	}

	// 迭代取出数据
	for _, student := range students {
		m.Log.Debug("查询数据成功", "student", student)
	}
}

func UpdateById(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"孙悟空111", 22, "男"}
	var err error

	// 查询一条数据
	updated, err := m.UpdateById(table, columns, values, 1)
	if err != nil {
		m.Log.Panic("根据ID更新数据失败", "error", err)
	}
	m.Log.Debug("更新数据成功", "updated", updated)
}

func UpdateByIds(m *zdpgo_mysql.Mysql) {
	table := "students"
	columns := []string{"age"}
	values := []interface{}{33}
	ids := []int64{1, 2, 3}
	var err error

	// 更新
	updated, err := m.UpdateByIds(table, columns, values, ids)
	if err != nil {
		m.Log.Panic("根据ID更新数据失败", "error", err)
	}
	m.Log.Debug("更新数据成功", "updated", updated)
}

func DeleteById(m *zdpgo_mysql.Mysql) {
	table := "students"
	var err error

	// 删除
	deleted, err := m.DeleteById(table, 1)
	if err != nil {
		m.Log.Panic("根据ID删除数据失败", "error", err)
	}
	m.Log.Debug("删除数据成功", "deleted", deleted)
}

func DeleteByIds(m *zdpgo_mysql.Mysql) {
	table := "students"
	var err error

	// 删除
	deleted, err := m.DeleteByIds(table, 2, 3)
	if err != nil {
		m.Log.Panic("根据ID删除数据失败", "error", err)
	}
	m.Log.Debug("删除数据成功", "deleted", deleted)
}
