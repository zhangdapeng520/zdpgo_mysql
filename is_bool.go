package zdpgo_mysql

import (
	"database/sql"
	"fmt"
)

/*
@Time : 2022/5/20 0:04
@Author : 张大鹏
@File : is_bool
@Software: Goland2021.3.1
@Description: is类型的判断方法
*/

// IsHealth 判断MySQL是否可用
func (m *Mysql) IsHealth() bool {
	// 准备连接地址
	if m.Address == "" {
		m.Address = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			m.Config.Username,
			m.Config.Password,
			m.Config.Host,
			m.Config.Port,
			m.Config.Database)
	}

	// 连接
	var err error
	m.Log.Debug("正常尝试连接MySQL服务器", "address", m.Address)
	m.Db, err = sql.Open("mysql", m.Address)
	if err != nil {
		m.Log.Error("连接MySQL数据库服务器失败", "error", err, "address", m.Address)
		return false
	}
	m.Log.Debug("连接MySQL服务器成功")
	return true
}
