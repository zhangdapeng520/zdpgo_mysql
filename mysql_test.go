package zdpgo_mysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangdapeng520/zdpgo_log"
)

// 测试建立连接
func TestConect(t *testing.T){
	log:= zdpgo_log.NewLogger("zdpgo_mysql.log")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err!= nil{
		log.Error("建立数据库连接失败：", err)
	}
	log.Info(db)
	defer db.Close()
}

// 测试建立连接
func TestConect1(t *testing.T){
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	fmt.Println(db)
	defer db.Close()
}
