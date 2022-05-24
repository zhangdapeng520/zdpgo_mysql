package zdpgo_mysql

import (
	"database/sql"
	"errors"
)

/*
@Time : 2022/5/20 19:52
@Author : 张大鹏
@File : execute
@Software: Goland2021.3.1
@Description: execute 执行SQL语句相关的代码
*/

// Execute 执行SQL语句
func (m *Mysql) Execute(sqlStr string, args ...interface{}) (result sql.Result, err error) {
	var (
		stmt *sql.Stmt // 预处理对象
	)

	// 预处理SQL
	stmt, err = m.Db.Prepare(sqlStr)
	if err != nil {
		m.Log.Error("预处理SQL失败", "error", err)
		return
	}
	defer stmt.Close()

	// 执行SQL语句
	result, err = stmt.Exec(args...)
	if err != nil {
		m.Log.Error("执行SQL语句失败", "error", err, "sql", sqlStr, "args", args)
		return
	}

	// 返回执行结果
	return
}

// Begin 开始事务
func (m *Mysql) Begin() error {
	tx, err := m.Db.Begin()
	m.Tx = tx
	return err
}

// Rollback 回滚事务
func (m *Mysql) Rollback() error {
	// 事务不存在
	var msg string
	if m.Tx == nil {
		msg = "事务为空，回滚事务失败"
		return errors.New(msg)
	}

	// 事务存在，回滚事务
	err := m.Tx.Rollback()
	return err
}

// Commit 提交事务
func (m *Mysql) Commit() error {
	// 事务不存在
	var msg string
	if m.Tx == nil {
		msg = "事务为空，提交事务失败"
		return errors.New(msg)
	}

	// 事务存在，提交事务
	err := m.Tx.Commit()
	return err
}
