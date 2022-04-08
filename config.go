package zdpgo_mysql

// Config MySQL配置信息
type Config struct {
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

func getDefaultConfig(config Config) Config {
	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_mysql.log"
	}

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

	return config
}
