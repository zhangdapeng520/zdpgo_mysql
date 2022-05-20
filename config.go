package zdpgo_mysql

// Config MySQL配置信息
type Config struct {
	Debug             bool   `yaml:"debug" json:"debug"`                               // 是否为Debug模式
	LogFilePath       string `yaml:"log_file_path" json:"log_file_path"`               // 日志路径
	Host              string `yaml:"host" json:"host"`                                 // ip
	Port              int    `yaml:"port" json:"port"`                                 // 端口
	Username          string `yaml:"username" json:"username"`                         // 用户名
	Password          string `yaml:"password" json:"password"`                         // 密码
	Database          string `yaml:"database" json:"database"`                         // 数据库
	MaxConnectNum     int    `yaml:"max_connect_num" json:"max_connect_num"`           // 最大连接数
	MaxFreeConnectNum int    `yaml:"max_free_connect_num" json:"max_free_connect_num"` // 最大闲置连接数
}
