package zdpgo_mysql

import "github.com/zhangdapeng520/zdpgo_mysql/core/mysql"

type Mysql struct {
	Db *mysql.Mysql
}

func New(config MysqlConfig) *Mysql {
	m := Mysql{}

	m.Db = mysql.New(mysql.MysqlConfig{
		Debug:             config.Debug,
		Host:              config.Host,
		Port:              config.Port,
		Username:          config.Username,
		Password:          config.Password,
		Database:          config.Database,
		LogFilePath:       config.LogFilePath,
		MaxConnectNum:     config.MaxConnectNum,
		MaxFreeConnectNum: config.MaxFreeConnectNum,
	})

	return &m
}
