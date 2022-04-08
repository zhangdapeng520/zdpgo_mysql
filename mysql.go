package zdpgo_mysql

import (
	"github.com/zhangdapeng520/zdpgo_mysql/core/mysql"
	"github.com/zhangdapeng520/zdpgo_mysql/core/table"
)

type Mysql struct {
	Db    *mysql.Mysql // 操作MySQL数据库的核心对象
	Table *table.Table // 操作表格的核心对象
}

func New(config Config) *Mysql {
	m := Mysql{}
	config = getDefaultConfig(config)

	// 初始化核心对象
	m.Db = mysql.New(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.MaxConnectNum,
		config.MaxFreeConnectNum,
	)
	m.Table = table.NewTable(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	return &m
}
