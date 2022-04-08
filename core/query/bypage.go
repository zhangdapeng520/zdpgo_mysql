package query

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_mysql/libs/sqlscan"
	"strings"
)

// FindByPage 分页查询数据
func (m *Query) FindByPage(table string, columns []string, page, size int) (rows *sql.Rows, err error) {
	// 整理字段列表
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 计算偏移量
	offset := (page - 1) * size
	s := fmt.Sprintf("SELECT %s FROM %s LIMIT %d,%d;", columnsStr, table, offset, size)

	// 执行查询
	rows, err = m.Db.Query(s)

	// 处理查询错误
	if err != nil {
		return nil, err
	}

	// 正常返回
	return rows, nil
}

// FindByPageToStruct 执行分页查询并映射到结构体
func (m *Query) FindByPageToStruct(table string, columns []string, page, size int, objects interface{}) (err error) {
	// 整理字段列表
	columnsStr := "*"
	if columns != nil && len(columns) > 0 {
		columnsStr = strings.Join(columns, ",")
	}

	// 计算偏移量
	offset := (page - 1) * size
	s := fmt.Sprintf("SELECT %s FROM %s LIMIT %d,%d;", columnsStr, table, offset, size)

	// 结构体映射
	ctx := context.Background()
	err = sqlscan.Select(ctx, m.Db, objects, s)
	if err != nil {
		return err
	}

	// 正常返回
	return nil
}
