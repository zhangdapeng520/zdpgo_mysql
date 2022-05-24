package zdpgo_mysql

import (
	"database/sql"
	"errors"
)

/*
@Time : 2022/5/20 8:56
@Author : 张大鹏
@File : common
@Software: Goland2021.3.1
@Description: common 通用方法
*/

// QueryRow 查询单条数据
func (m *Mysql) QueryRow(sql string, args ...interface{}) (row *sql.Row, err error) {
	// 预处理SQL语句
	stmt, err := m.Db.Prepare(sql)

	// 处理错误
	if err != nil {
		return nil, err
	}
	if stmt == nil {
		err = errors.New("stml为nil，预处理语句失败")
		return nil, err
	}

	// 执行查询
	if args != nil {
		row = stmt.QueryRow(args...)
	} else {
		row = stmt.QueryRow()
	}

	// 关闭预处理对象
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	// 正常返回
	return row, nil
}

// Query 查询多条数据
func (m *Mysql) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := m.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	return rows, err
}
