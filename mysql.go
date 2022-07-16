package zdpgo_mysql

import (
	"database/sql"
	_ "github.com/zhangdapeng520/zdpgo_mysql/mysql_driver"
)

type Mysql struct {
	Config  *Config
	Db      *sql.DB // db核心对象
	Tx      *sql.Tx // 事务对象
	Address string  // Mysql连接地址
}

func New() *Mysql {
	return NewWithConfig(&Config{})
}

// NewWithConfig 根据配置新增MySQL对象
func NewWithConfig(config *Config) *Mysql {
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

	// 配置
	m.Config = config

	// 连接
	m.IsHealth()

	return &m
}
