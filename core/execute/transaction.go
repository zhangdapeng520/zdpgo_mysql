package execute

import "errors"

// Begin 开始事务
func (m *Execute) Begin() error {
	tx, err := m.Db.Begin()
	m.tx = tx
	return err
}

// Rollback 回滚事务
func (m *Execute) Rollback() error {
	// 事务不存在
	var msg string
	if m.tx == nil {
		msg = "事务为空，回滚事务失败"
		return errors.New(msg)
	}

	// 事务存在，回滚事务
	err := m.tx.Rollback()
	return err
}

// Commit 提交事务
func (m *Execute) Commit() error {
	// 事务不存在
	var msg string
	if m.tx == nil {
		msg = "事务为空，提交事务失败"
		return errors.New(msg)
	}

	// 事务存在，提交事务
	err := m.tx.Commit()
	return err
}
