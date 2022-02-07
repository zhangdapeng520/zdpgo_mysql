package zdpgo_mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Add 添加数据
func (m *Mysql) Add(sql string, args ...interface{}) (id int64, err error) {
	m.log.Info("Add 添加数据", "sql", sql, "args", args)

	// 执行SQL语句
	ret := m.Execute(sql, args...)

	// 执行结果处理
	affected, err := ret.RowsAffected()
	if affected <= 0 {
		m.log.Error("受影响的行数为0，插入数据失败")
		err = errors.New("受影响的行数为0，插入数据失败")
		return 0, err
	}
	if err != nil {
		m.log.Error("Add 添加数据失败", "error", err.Error())
		return 0, err
	}

	// 获取新插入的数据的ID
	uid, err := ret.LastInsertId()

	// 处理错误
	if err != nil {
		m.log.Error("获取新插入的数据ID失败", "error", err)
		return 0, err
	}

	// 正常返回
	return uid, nil
}

// DeleteById 根据ID删除数据
func (m *Mysql) DeleteById(table string, id int) (deleted int64, err error) {
	// 整理SQL语句
	s := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", table, id)
	m.log.Info("Delete 根据ID删除数据", "table", table, "id", id, "sql", s)

	// 执行SQL语句
	ret := m.Execute(s)

	// 处理执行结果
	rows, err := ret.RowsAffected()
	if err != nil {
		m.log.Error("获取受影响的行数失败", "error", err)
		return 0, err
	}
	if rows <= 0 {
		err = errors.New("受影响的行数为0，删除数据失败")
		return 0, err
	}

	// 正常结果
	return rows, nil
}

// DeleteByIds 根据ID列表删除
func (m *Mysql) DeleteByIds(table string, ids ...int) (deleted int64, err error) {
	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(v)
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 整理SQL语句
	s := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s);",
		table, idsStr)
	m.log.Info("根据ID列表删除", "sql", s)

	// 执行SQL语句
	ret := m.Execute(s)

	// 处理执行结果
	rows, err := ret.RowsAffected()
	if err != nil {
		m.log.Error("获取受影响的行数失败", "error", err)
		return 0, err
	}
	if rows <= 0 {
		m.log.Error("受影响的行数为0，根据ID列表删除失败")
		err = errors.New("受影响的行数为0，根据ID列表删除失败")
		return 0, err
	}

	// 正常返回
	return rows, nil
}

// Update 根据ID修改数据
func (m *Mysql) Update(sql string, args ...interface{}) (updated int64, err error) {
	m.log.Info("Update 根据ID修改数据", "sql", sql, "args", args)

	// 执行SQL语句
	ret := m.Execute(sql, args...)

	// 处理执行结果
	rows, err := ret.RowsAffected()
	if err != nil {
		m.log.Error("获取受影响的行数失败", "error", err)
		return 0, err
	}
	if rows <= 0 {
		m.log.Error("受影响的行数为0，更新数据失败")
		err = errors.New("受影响的行数为0，更新数据失败")
		return 0, err
	}

	// 正常返回
	return rows, nil
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
	m.log.Info("FindById 查询单条数据", "sql", s)

	// 执行SQL语句
	row, err = m.QueryRow(s, id)
	if err != nil {
		m.log.Error("QueryRow 执行查询失败", "error", err.Error())
		return nil, err
	}

	// 正常返回
	return row, err
}

// FindByIds 根据ID列表查询数据
func (m *Mysql) FindByIds(table string, columns []string, ids []int) (rows *sql.Rows, err error) {
	// 参数校验
	if ids == nil {
		m.log.Error("ids不能为空")
		return nil, errors.New("ids不能为空")
	}

	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(v)
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 整理字段列表
	columnsStr := "*"
	if columns != nil {
		columnsStr = strings.Join(columns, ",")
	}

	// 执行SQL语句
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s);",
		columnsStr, table, idsStr)

	// 执行查询
	rows, err = m.Query(s)

	// 处理查询错误
	if err != nil {
		m.log.Error("Query 查询数据失败", "error", err)
		return nil, err
	}

	// 正常返回
	return rows, nil
}

// FindPages 分页查询数据
func (m *Mysql) FindPages(table string, columns []string, page, size int) (rows *sql.Rows, err error) {
	// 整理字段列表
	columnsStr := "*"
	if columns != nil {
		columnsStr = strings.Join(columns, ",")
	}

	// 计算偏移量
	offset := (page - 1) * size
	s := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d,%d;", columnsStr, table, offset, size)

	// 执行查询
	rows, err = m.db.Query(s)

	// 处理查询错误
	if err != nil {
		m.log.Error("Query 查询多条数据失败", "error", err)
		return nil, err
	}

	// 正常返回
	return rows, nil
}
