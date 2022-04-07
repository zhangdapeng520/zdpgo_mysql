package mysql

import (
	"testing"
)

func prepareMysql() *Mysql {
	m := New(MysqlConfig{
		Debug:    true,
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	})
	return m
}

func TestMysql_New(t *testing.T) {
	m := prepareMysql()
	t.Log(m)
}
