package common

import (
	"github.com/zhangdapeng520/zdpgo_mysql"
)

/*
@Time : 2022/5/20 20:00
@Author : 张大鹏
@File : common
@Software: Goland2021.3.1
@Description: 通用方法测试
*/

func IsHealth(m *zdpgo_mysql.Mysql) {
	if !m.IsHealth() {
		panic("无法正常连接MySQL服务")
	}
}
