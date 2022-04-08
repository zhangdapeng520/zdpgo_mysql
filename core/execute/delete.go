package execute

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DeleteById 根据ID删除数据
func (m *Execute) DeleteById(table string, id int64) (deleted int64, err error) {
	// 整理SQL语句
	s := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", table, id)

	// 执行SQL语句
	ret, err := m.Execute(s)
	if err != nil {
		return 0, err
	}

	// 处理执行结果
	rows, err := ret.RowsAffected()
	if err != nil {
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
func (m *Execute) DeleteByIds(table string, ids ...int64) (deleted int64, err error) {
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

	// 执行SQL语句
	ret, err := m.Execute(s)
	if err != nil {
		return 0, err
	}

	// 处理执行结果
	rows, err := ret.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rows <= 0 {
		err = errors.New("受影响的行数为0，根据ID列表删除失败")
		return 0, err
	}

	// 正常返回
	return rows, nil
}
