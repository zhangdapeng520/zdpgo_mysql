package zdpgo_mysql

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// 测试建立连接
func TestMysql_New(t *testing.T) {
	m := New(MysqlConfig{
		Debug:    true,
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "user_service",
	})
	fmt.Println(m)
	defer m.Close()
}
