package zdpgo_mysql

import (
	"database/sql"
	"fmt"

	"github.com/zhangdapeng520/zdpgo_log"
)

type Mysql struct {
	Host string // ip
	Port int // 端口
	Username string // 用户名
	Password string // 密码
	Database string // 数据库
	LogFile string // 日志名称
	MaxConnectNum int // 最大连接数
	MaxFreeConnectNum int // 最大闲置连接数
	Logger *zdpgo_log.Logger
	Db *sql.DB
}

// 建立连接
func (mysql *Mysql) Init() {
	// 初始化日志
	if mysql.LogFile == ""{
		mysql.LogFile = "zdpgo_mysql.log"
	}
	mysql.Logger = zdpgo_log.NewLogger(mysql.LogFile)

	// 初始化Db
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",mysql.Username,mysql.Password,mysql.Host,mysql.Port,mysql.Database)
	db, err:= sql.Open("mysql", address)
	if err!=nil{
		mysql.Logger.Error("连接数据库失败：", err)
	}

	// 尝试与数据库建立连接
	err = db.Ping()
	if err!= nil{
		mysql.Logger.Error("Ping数据库失败：", err)
	}

	// 初始化
	mysql.Db = db

	// 设置最大连接数和最大闲置连接数
	if mysql.MaxConnectNum == 0{
		mysql.MaxConnectNum = 100
	}
	mysql.Db.SetMaxOpenConns(mysql.MaxConnectNum)

	if mysql.MaxFreeConnectNum==0{
		mysql.MaxFreeConnectNum = 10
	}
	mysql.Db.SetMaxIdleConns(mysql.MaxFreeConnectNum)
}

// 关闭数据库连接
func (mysql *Mysql) Close(){
	mysql.Db.Close()
}