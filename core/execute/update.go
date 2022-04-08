package execute

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// UpdateById 根据ID修改数据
func (m *Execute) UpdateById(table string, columns []string, values []interface{}, id int64) (updated int64, err error) {
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

	// 执行更新
	update, err := m.update(s, args...)
	if err != nil {
		return 0, err
	}

	// 正常返回
	return update, nil
}

// UpdateByIds 根据ID列表修改数据
func (m *Execute) UpdateByIds(table string, columns []string, values []interface{}, ids []int64) (updated int64, err error) {
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

	// 执行更新
	update, err := m.update(s, args...)
	if err != nil {
		return 0, err
	}

	// 正常返回
	return update, nil
}

func (m *Execute) update(sql string, args ...interface{}) (updated int64, err error) {
	// 执行SQL语句
	ret, err := m.Execute(sql, args...)
	if err != nil {
		return 0, err
	}

	// 处理执行结果
	rows, err := ret.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rows <= 0 {
		err = errors.New("受影响的行数为0，更新数据失败")
		return 0, err
	}

	// 正常返回
	return rows, nil
}
