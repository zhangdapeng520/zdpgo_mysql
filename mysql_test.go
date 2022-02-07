package zdpgo_mysql

import (
	"fmt"
	"testing"
)

func prepareMysql() *Mysql {
	m := New(MysqlConfig{
		Debug:    true,
		Host:     "192.168.33.101",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	})
	return m
}

func TestMysql_New(t *testing.T) {
	m := prepareMysql()
	fmt.Println(m)
}
