package zdpgo_mysql

import (
	"database/sql"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_mysql/core/execute"
	"github.com/zhangdapeng520/zdpgo_mysql/core/query"
	"github.com/zhangdapeng520/zdpgo_mysql/core/table"
)

type Mysql struct {
	Table   *table.Table     // 操作表格的核心对象
	Execute *execute.Execute // 执行增删改语句的核心对象
	Query   *query.Query     // 查询数据的核心对象
	Config  *Config
	Log     *zdpgo_log.Log // 日志对象
	Db      *sql.DB        // db核心对象
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

	// 初始化核心对象
	m.Table = table.NewTable(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	m.Execute = execute.NewExecute(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	m.Query = query.NewQuery(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	return &m
}
