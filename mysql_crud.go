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

// AddMany 批量添加数据
func (m *Mysql) AddMany(table string, columns []string, values [][]interface{}) (affected int64, err error) {
	// 处理异常
	if columns == nil {
		return 0, errors.New("columns字段列表不能为空")
	}
	if values == nil {
		return 0, errors.New("values值列表不能为空")
	}
	if len(columns) != len(values[0]) {
		return 0, errors.New("columns的长度和values子数组长度不相等，字段无法映射")
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
	m.log.Info("批量插入数据", "sql", s, "args", args)

	// 执行SQL语句
	ret := m.Execute(s, args...)

	// 执行结果处理
	affected, err = ret.RowsAffected()
	if affected <= 0 {
		m.log.Error("受影响的行数为0，批量插入数据失败")
		err = errors.New("受影响的行数为0，批量插入数据失败")
		return 0, err
	}
	if err != nil {
		m.log.Error("AddMany 批量添加数据失败", "error", err.Error())
		return 0, err
	}

	// 返回受影响的行数
	return affected, nil
}

// Add 添加数据
func (m *Mysql) Add(table string, columns []string, values []interface{}) (id int64, err error) {
	// 处理异常
	if columns == nil {
		return 0, errors.New("columns字段列表不能为空")
	}
	if values == nil {
		return 0, errors.New("values值列表不能为空")
	}
	if len(columns) != len(values) {
		return 0, errors.New("columns的长度和values长度不相等，字段无法映射")
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
		return 0, err
	}
	m.log.Info("插入数据成功", "addId", addId)

	// 返回添加的id
	return addId, nil
}

func (m *Mysql) add(sql string, args ...interface{}) (id int64, err error) {
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
func (m *Mysql) DeleteById(table string, id int64) (deleted int64, err error) {
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
func (m *Mysql) DeleteByIds(table string, ids ...int64) (deleted int64, err error) {
	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(int(v))
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

// UpdateById 根据ID修改数据
func (m *Mysql) UpdateById(table string, columns []string, values []interface{}, id int64) (updated int64, err error) {
	// 异常情况
	if table == "" {
		return 0, errors.New("table不能为空")
	}
	if columns == nil {
		return 0, errors.New("columns不能为空")
	}
	if values == nil {
		return 0, errors.New("values不能为空")
	}
	if len(columns) != len(values) {
		return 0, errors.New("columns和values长度不一致，无法映射")
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
	m.log.Info("UpdateById 根据ID修改数据", "sql", s, "args", args)

	// 执行更新
	update, err := m.update(s, args...)
	if err != nil {
		m.log.Error("执行更新失败", "error", err.Error())
		return 0, err
	}

	// 正常返回
	return update, nil
}

// UpdateByIds 根据ID列表修改数据
func (m *Mysql) UpdateByIds(table string, columns []string, values []interface{}, ids []int64) (updated int64, err error) {
	// 异常情况
	if table == "" {
		return 0, errors.New("table不能为空")
	}
	if columns == nil {
		return 0, errors.New("columns不能为空")
	}
	if values == nil {
		return 0, errors.New("values不能为空")
	}
	if len(columns) != len(values) {
		return 0, errors.New("columns和values长度不一致，无法映射")
	}
	if ids == nil || len(ids) < 1 {
		return 0, errors.New("ids不能为空")
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
	m.log.Info("UpdateByIds 根据ID列表修改数据", "sql", s, "args", args)

	// 执行更新
	update, err := m.update(s, args...)
	if err != nil {
		m.log.Error("执行更新失败", "error", err.Error())
		return 0, err
	}

	// 正常返回
	return update, nil
}

func (m *Mysql) update(sql string, args ...interface{}) (updated int64, err error) {
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

// FindByIdToStruct 根据id查询数据，并将结果转换为结构体
func (m *Mysql) FindByIdToStruct(table string, columns []string, id int64, objects interface{}) (err error) {
	// 处理字段
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 整理SQL语句
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?;", columnsStr, table)
	m.log.Info("FindById 查询单条数据", "sql", s)

	// 结构体映射
	ctx := context.Background()
	err = sqlscan.Select(ctx, m.db, objects, s, id)
	if err != nil {
		m.log.Error("执行查询并映射数据失败", "error", err.Error())
		return err
	}

	// 正常返回
	return nil
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
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s) ORDER BY ID DESC;",
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

// FindByIdsToStruct 根据ID列表查询数据并映射到结构体
func (m *Mysql) FindByIdsToStruct(table string, columns []string, ids []int64, objects interface{}) (err error) {
	// 参数校验
	if ids == nil {
		m.log.Error("ids不能为空")
		return errors.New("ids不能为空")
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
	err = sqlscan.Select(ctx, m.db, objects, s)
	if err != nil {
		m.log.Error("执行查询并映射数据失败", "error", err.Error())
		return err
	}

	// 正常返回
	return nil
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

// FindByPageToStruct 执行分页查询并映射到结构体
func (m *Mysql) FindByPageToStruct(table string, columns []string, page, size int, objects interface{}) (err error) {
	// 整理字段列表
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 计算偏移量
	offset := (page - 1) * size
	s := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d,%d;", columnsStr, table, offset, size)

	// 结构体映射
	ctx := context.Background()
	err = sqlscan.Select(ctx, m.db, objects, s)
	if err != nil {
		m.log.Error("执行查询并映射数据失败", "error", err.Error())
		return err
	}

	// 正常返回
	return nil
}
