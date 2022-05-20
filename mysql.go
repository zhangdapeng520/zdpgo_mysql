package zdpgo_mysql

import (
	"database/sql"
	"github.com/zhangdapeng520/zdpgo_log"
	_ "github.com/zhangdapeng520/zdpgo_mysql/mysql_driver"
)

type Mysql struct {
	Config  *Config
	Log     *zdpgo_log.Log // 日志对象
	Db      *sql.DB        // db核心对象
	Tx      *sql.Tx        // 事务对象
	Address string         // Mysql连接地址
}

func New(config *Config) *Mysql {
	m := Mysql{}

	// 基本信息
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 3306
	}
	if config.Username == "" {
		config.Username = "root"
	}
	if config.Password == "" {
		config.Password = "root"
	}
	if config.Database == "" {
		config.Database = "test"
	}

	// 设置最大连接数
	if config.MaxConnectNum == 0 {
		config.MaxConnectNum = 100
	}

	// 设置最大闲置连接数
	if config.MaxFreeConnectNum == 0 {
		config.MaxFreeConnectNum = 10
	}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_mysql.log"
	}
	m.Log = zdpgo_log.NewWithDebug(config.Debug, config.LogFilePath)

	// 配置
	m.Config = config

	// 连接
	if !m.IsHealth() {
		m.Log.Error("无法连接MySQL服务")
	}

	return &m
}
