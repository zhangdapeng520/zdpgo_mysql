package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/zhangdapeng520/zdpgo_mysql/libs/mysql_driver"
)

// Mysql 操作MySQL核心对象
type Mysql struct {
	db *sql.DB // db核心对象
	tx *sql.Tx // 事务对象
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
		config.LogFilePath = "logs/zdpgo/zdpgo_mysql.log"
	}
	// 初始化连接
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", address)
	if err != nil {
		panic(err)
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
	return err
}

// Execute 执行SQL语句
func (m *Mysql) Execute(sql string, args ...interface{}) sql.Result {
	// 预处理SQL
	stmt, err := m.db.Prepare(sql)

	// 如果存在事务，则使用事务
	if m.tx != nil {
		stmt, err = m.tx.Prepare(sql)
	}

	if err != nil {
		return nil
	}
	defer stmt.Close()

	// 执行SQL语句
	ret, err := stmt.Exec(args...)
	if err != nil {
		// 如果事务不为空，需要回滚事务
		if m.tx != nil {
			_ = m.Rollback()
		}
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
	// 预处理SQL语句
	stmt, err := m.db.Prepare(sql)

	// 处理错误
	if err != nil {
		return nil, err
	}
	if stmt == nil {
		err = errors.New("stml为nil，预处理语句失败")
		return nil, err
	}

	// 执行查询
	if args != nil {
		row = stmt.QueryRow(args...)
	} else {
		row = stmt.QueryRow()
	}

	// 关闭预处理对象
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	// 正常返回
	return row, nil
}

// Query 查询多条数据
func (m *Mysql) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := m.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	return rows, err
}

// TransRowsToMaps 将rows转换为map
func (m *Mysql) TransRowsToMaps(rows *sql.Rows) (list []map[string]interface{}, err error) {
	// 获取字段列表
	columns, _ := rows.Columns()
	columnLength := len(columns)

	// 临时存储每行数据
	cache := make([]interface{}, columnLength)
	for index, _ := range cache { //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}

	// 转换数据
	for rows.Next() {
		err = rows.Scan(cache...)
		if err != nil {
			return nil, err
		}
		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}

	// 关闭rows
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}
