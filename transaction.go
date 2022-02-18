package zdpgo_mysql

import (
	"github.com/pkg/errors"
)

// Begin 开始事务
func (m *Mysql) Begin() error {
	tx, err := m.db.Begin()
	if err != nil {
		m.log.Error("开启事务失败", "error", err.Error())
	}
	m.tx = tx
	return err
}

// Rollback 回滚事务
func (m *Mysql) Rollback() error {
	// 事务不存在
	var msg string
	if m.tx == nil {
		msg = "事务为空，回滚事务失败"
		m.log.Error(msg)
		return errors.New(msg)
	}

	// 事务存在，回滚事务
	err := m.tx.Rollback()
	if err != nil {
		m.log.Error("回滚事务失败", "error", err.Error())
	}
	return err
}

// Commit 提交事务
func (m *Mysql) Commit() error {
	// 事务不存在
	var msg string
	if m.tx == nil {
		msg = "事务为空，提交事务失败"
		m.log.Error(msg)
		return errors.New(msg)
	}

	// 事务存在，提交事务
	err := m.tx.Commit()
	if err != nil {
		m.log.Error("回滚事务失败", "error", err.Error())
	}
	return err
}
