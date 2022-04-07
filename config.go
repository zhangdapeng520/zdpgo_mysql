package zdpgo_mysql

// MysqlConfig MySQL配置信息
type MysqlConfig struct {
	Debug             bool   // 是否为debug模式
	Host              string // ip
	Port              int    // 端口
	Username          string // 用户名
	Password          string // 密码
	Database          string // 数据库
	LogFilePath       string // 日志路径
	MaxConnectNum     int    // 最大连接数
	MaxFreeConnectNum int    // 最大闲置连接数
}
