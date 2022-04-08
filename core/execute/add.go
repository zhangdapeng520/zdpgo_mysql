package execute

import (
	"errors"
	"fmt"
	"strings"
)

// AddMany 批量添加数据
func (m *Execute) AddMany(table string, columns []string, values [][]interface{}) (affected int64, err error) {
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

	// 执行SQL语句
	ret, err := m.Execute(s, args...)
	if err != nil {
		return 0, err
	}

	// 执行结果处理
	affected, err = ret.RowsAffected()
	if affected <= 0 {
		err = errors.New("受影响的行数为0，批量插入数据失败")
		return 0, err
	}
	if err != nil {
		return 0, err
	}

	// 返回受影响的行数
	return affected, nil
}

// Add 添加数据
func (m *Execute) Add(table string, columns []string, values []interface{}) (id int64, err error) {
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

	// 返回添加的id
	return addId, nil
}

func (m *Execute) add(sql string, args ...interface{}) (id int64, err error) {
	// 执行SQL语句
	ret, err := m.Execute(sql, args...)
	if err != nil {
		return 0, err
	}

	// 执行结果处理
	affected, err := ret.RowsAffected()
	if affected <= 0 {
		err = errors.New("受影响的行数为0，插入数据失败")
		return 0, err
	}
	if err != nil {
		return 0, err
	}

	// 获取新插入的数据的ID
	uid, err := ret.LastInsertId()

	// 处理错误
	if err != nil {
		return 0, err
	}

	// 正常返回
	return uid, nil
}
