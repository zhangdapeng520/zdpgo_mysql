package zdpgo_mysql

import (
	"database/sql"
	"fmt"

	"github.com/zhangdapeng520/zdpgo_log"
)

// Mysql MySQL的核心对象
type Mysql struct {
	log    *zdpgo_log.Log // 日志对象
	DB     *sql.DB        // 操作数据库的核心对象
	config MysqlConfig    // 配置对象
}

// MysqlConfig mysql的配置对象
type MysqlConfig struct {
	Debug             bool   // 是否为调试模式
	Host              string // ip
	Port              int    // 端口
	Username          string // 用户名
	Password          string // 密码
	Database          string // 数据库
	LogFilePath       string // 日志存放路径
	MaxConnectNum     int    // 最大连接数
	MaxFreeConnectNum int    // 最大闲置连接数
}

// New 创建MySQL的实例
func New(config MysqlConfig) *Mysql {
	m := Mysql{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_mysql.log"
	}
	l := zdpgo_log.New(zdpgo_log.LogConfig{
		Debug: config.Debug,
		Path:  config.LogFilePath,
	})
	m.log = l

	// 校验参数：username password host port database
	if config.Host == "" {
		m.log.Warning("MySQL数据库的服务主机地址为空，将使用默认值：127.0.0.1")
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		m.log.Warning("MySQL数据库的服务端口号为空，将使用默认值：3306")
		config.Port = 3306
	}
	if config.Username == "" {
		m.log.Warning("MySQL数据库的服务用户名，将使用默认值：root")
		config.Username = "root"
	}
	if config.Password == "" {
		m.log.Warning("MySQL数据库的服务密码为空，将使用默认值：root")
		config.Password = "root"
	}
	if config.Database == "" {
		m.log.Panic("MySQL的数据库不能为空，必须指定一个要连接的数据库！")
	}

	// 设置最大连接数和最大闲置连接数
	if config.MaxConnectNum == 0 {
		config.MaxConnectNum = 100
	}

	if config.MaxFreeConnectNum == 0 {
		config.MaxFreeConnectNum = 10
	}

	// 初始化配置
	m.config = config

	// 初始化数据库连接
	m.init()

	return &m
}

// 建立连接
func (m *Mysql) init() {
	// 初始化Db
	var err error
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.config.Username, m.config.Password, m.config.Host, m.config.Port, m.config.Database)
	m.DB, err = sql.Open("mysql", address)
	if err != nil {
		m.log.Error("连接数据库失败：", err)
	}

	// 尝试与数据库建立连接
	err = m.DB.Ping()
	if err != nil {
		m.log.Error("Ping数据库失败：", err)
	}

	// 设置最大连接数和最大闲置连接数
	m.DB.SetMaxOpenConns(m.config.MaxConnectNum)
	m.DB.SetMaxIdleConns(m.config.MaxFreeConnectNum)
}

// Close 关闭数据库连接
func (m *Mysql) Close() {
	err := m.DB.Close()
	if err != nil {
		m.log.Error("关闭MySQL数据库连接失败：", err)
	}
}

// Execute 执行SQL语句
func (m *Mysql) Execute(sql string, args ...interface{}) sql.Result {
	m.log.Info("执行SQL语句：", sql)
	m.log.Info("参数：", args)
	ret, err := m.DB.Exec(sql, args...)
	if err != nil {
		m.log.Error("执行SQL语句失败：", err)
		return nil
	}
	return ret
}

// DeleteTable 删除表格
func (m *Mysql) DeleteTable(table string) {
	s := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
	m.Execute(s)
}

// QueryRow 查询单条数据
func (m *Mysql) QueryRow(sql string, id int) *sql.Row {
	row := m.DB.QueryRow(sql, id)
	return row
}

// Query 查询多条数据
func (m *Mysql) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	rows, err := m.DB.Query(sql, args...)
	return rows, err
}
