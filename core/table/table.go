package table

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/zhangdapeng520/zdpgo_mysql/libs/mysql_driver"
)

type Table struct {
	Db *sql.DB // db核心对象
}

func NewTable(username, password, host string, port int, database string) *Table {
	t := Table{}

	// 初始化DB
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	db, err := sql.Open("mysql", address)
	if err != nil {
		panic(err)
	}
	t.Db = db

	// 返回表格处理对象
	return &t
}

// Execute 执行SQL语句
func (m *Table) Execute(sql string, args ...interface{}) (sql.Result, error) {
	// 预处理SQL
	stmt, err := m.Db.Prepare(sql)
	if err != nil {
		return nil, errors.New("预处理SQL失败")
	}
	defer stmt.Close()

	// 执行SQL语句
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, errors.New("执行SQL语句失败")
	}

	// 返回执行结果
	return result, nil
}

// QueryRow 查询单条数据
func (m *Table) QueryRow(sql string, args ...interface{}) (row *sql.Row, err error) {
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
func (m *Table) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := m.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	return rows, err
}

// Add 增加表格
func (t *Table) Add(sql string) error {
	_, err := t.Execute(sql)
	return err
}

// Delete 删除表格
func (t *Table) Delete(table string) error {
	s := fmt.Sprintf("drop table if exists %s", table)
	_, err := t.Execute(s)
	return err
}

// Find 查询所有表格
func (t *Table) Find() ([]string, error) {
	s := "show tables;"
	rows, err := t.Query(s)
	if err != nil {
		return nil, err
	}

	// 循环读取数据
	var tables []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	// 返回表格数据
	return tables, nil
}
