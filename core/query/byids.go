package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_mysql/libs/sqlscan"
	"strconv"
	"strings"
)

// FindByIds 根据ID列表查询数据
func (m *Query) FindByIds(table string, columns []string, ids []int) (rows *sql.Rows, err error) {
	// 参数校验
	if ids == nil {
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
		return nil, err
	}

	// 正常返回
	return rows, nil
}

// FindByIdsToStruct 根据ID列表查询数据并映射到结构体
func (m *Query) FindByIdsToStruct(table string, columns []string, ids []int64, objects interface{}) (err error) {
	// 参数校验
	if ids == nil {
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
	err = sqlscan.Select(ctx, m.Db, objects, s)
	if err != nil {
		return err
	}

	// 正常返回
	return nil
}
