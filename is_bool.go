package zdpgo_mysql

import (
	"database/sql"
	"fmt"
	"time"
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
	address := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		m.Config.Username,
		m.Config.Password,
		m.Config.Host,
		m.Config.Port,
		m.Config.Database)

	// 连接
	var err error
	for i := 0; i < 3; i++ {
		m.Db, err = sql.Open("mysql", address)
		if err == nil {
			break // 连接成功
		}
		time.Sleep(time.Second * 3)
	}
	if err != nil {
		m.Log.Error("连接数据库失败", "error", err)
		return false
	}

	return true
}
