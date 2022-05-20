package zdpgo_mysql

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/20 8:28
@Author : 张大鹏
@File : is_bool_test
@Software: Goland2021.3.1
@Description: is类型的判断方法的相关测试
*/

// 测试能否正常连接MySQL数据库
func TestMysql_IsHealth(t *testing.T) {
	m := getMysql()
	fmt.Println(m.IsHealth())
}
