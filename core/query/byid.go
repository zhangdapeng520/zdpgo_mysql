package query

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_mysql/libs/sqlscan"
	"strings"
)

// FindById 查询单条数据
func (m *Query) FindById(table string, columns []string, id int) (row *sql.Row, err error) {
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
		return nil, err
	}

	// 正常返回
	return row, err
}

// FindByIdToStruct 根据id查询数据，并将结果转换为结构体
func (m *Query) FindByIdToStruct(table string, columns []string, id int64, objects interface{}) (err error) {
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
		return err
	}

	// 正常返回
	return nil
}
