package zdpgo_mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangdapeng520/zdpgo_zap"
)

type Mysql struct {
	log *zdpgo_zap.Zap
	db  *sql.DB
}

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

// New 创建mysql实例
func New(config MysqlConfig) *Mysql {
	m := Mysql{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_mysql.log"
	}
	m.log = zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        config.Debug,
		OpenGlobal:   false,
		OpenFileName: false,
		LogFilePath:  config.LogFilePath,
	})

	// 初始化连接
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", address)
	if err != nil {
		m.log.Error("连接MySQL数据库失败", "error", err.Error())
	}
	m.db = db

	// 设置最大连接数和最大闲置连接数
	if config.MaxConnectNum == 0 {
		config.MaxConnectNum = 100
	}
	m.db.SetMaxOpenConns(config.MaxConnectNum)

	if config.MaxFreeConnectNum == 0 {
		config.MaxFreeConnectNum = 10
	}
	m.db.SetMaxIdleConns(config.MaxFreeConnectNum)

	return &m
}

// Close 关闭数据库连接
func (m *Mysql) Close() error {
	err := m.db.Close()
	if err != nil {
		m.log.Error("关闭MySQL数据库连接失败", "error", err.Error())
	}
	return err
}

// Execute 执行SQL语句
func (m *Mysql) Execute(sql string, args ...interface{}) sql.Result {
	m.log.Info("执行SQL语句", "sql", sql, "args", args)

	// 预处理SQL
	stmt, err := m.db.Prepare(sql)
	if err != nil {
		m.log.Error("Prepare SQL语句失败", "error", err.Error())
		return nil
	}
	defer stmt.Close()

	// 执行SQL语句
	ret, err := stmt.Exec(args...)
	if err != nil {
		m.log.Error("执行SQL语句失败", "error", err)
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
func (m *Mysql) QueryRow(sql string, args ...interface{}) (row *sql.Row, err error) {
	m.log.Info("QueryRow 查询单条数据", "sql", sql)

	// 预处理SQL语句
	stmt, err := m.db.Prepare(sql)

	// 处理错误
	if err != nil {
		m.log.Error("Prepare SQL语句失败", "error", err)
		return nil, err
	}
	if stmt == nil {
		m.log.Error("stml为nil，预处理语句失败")
		err = errors.New("stml为nil，预处理语句失败")
		return nil, err
	}

	// 执行查询
	if args != nil {
		row = stmt.QueryRow(args)
	} else {
		row = stmt.QueryRow()
	}

	// 关闭预处理对象
	err = stmt.Close()
	if err != nil {
		m.log.Error("关闭预处理对象失败", "error", err.Error())
		return nil, err
	}

	// 正常返回
	return row, nil
}

// Query 查询多条数据
func (m *Mysql) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	m.log.Info("Query 查询多条数据", "sql", sql, "args", args)
	stmt, err := m.db.Prepare(sql)
	if err != nil {
		m.log.Error("Prepare SQL语句失败", "error", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	return rows, err
}
