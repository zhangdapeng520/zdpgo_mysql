package zdpgo_mysql

import (
	"github.com/zhangdapeng520/zdpgo_mysql/core/execute"
	"github.com/zhangdapeng520/zdpgo_mysql/core/query"
	"github.com/zhangdapeng520/zdpgo_mysql/core/table"
)

type Mysql struct {
	Table   *table.Table     // 操作表格的核心对象
	Execute *execute.Execute // 执行增删改语句的核心对象
	Query   *query.Query     // 查询数据的核心对象
}

func New(config Config) *Mysql {
	m := Mysql{}
	config = getDefaultConfig(config)

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
