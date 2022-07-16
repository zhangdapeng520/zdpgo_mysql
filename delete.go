package zdpgo_mysql

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
@Time : 2022/5/20 20:32
@Author : 张大鹏
@File : delete
@Software: Goland2021.3.1
@Description: delete 删除相关方法
*/

// DeleteTable 删除表格
func (m *Mysql) DeleteTable(tableName string) error {
	s := fmt.Sprintf("drop table if exists %s;", tableName)
	_, err := m.Execute(s)
	return err
}

// DeleteById 根据ID删除数据
func (m *Mysql) DeleteById(table string, id int64) (deleted int64, err error) {
	// 整理SQL语句
	s := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", table, id)

	// 执行SQL语句
	ret, err := m.Execute(s)
	if err != nil {
		return
	}

	// 处理执行结果
	deleted, err = ret.RowsAffected()
	if err != nil {
		return
	}
	if deleted <= 0 {
		err = errors.New("受影响的行数为0，删除数据失败")
		return
	}

	// 正常结果
	return
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
	s := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s);", table, idsStr)

	// 执行SQL语句
	ret, err := m.Execute(s)
	if err != nil {
		return
	}

	// 处理执行结果
	deleted, err = ret.RowsAffected()
	if err != nil {
		return
	}
	if deleted <= 0 {
		err = errors.New("受影响的行数为0，根据ID列表删除失败")
		return
	}

	// 正常返回
	return
}
