package execute

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/zhangdapeng520/zdpgo_mysql/libs/mysql_driver"
)

type Execute struct {
	Db *sql.DB // db核心对象
}

func NewExecute(username, password, host string, port int, database string) *Execute {
	t := Execute{}

	// 初始化DB
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	db, err := sql.Open("mysql", address)
	if err != nil {
		panic(err)
	}
	t.Db = db

	// 返回表格处理对象
	return &t
}

// Execute 执行SQL语句
func (m *Execute) Execute(sql string, args ...interface{}) (sql.Result, error) {
	// 预处理SQL
	stmt, err := m.Db.Prepare(sql)
	if err != nil {
		return nil, errors.New("预处理SQL失败")
	}
	defer stmt.Close()

	// 执行SQL语句
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, errors.New("执行SQL语句失败")
	}

	// 返回执行结果
	return result, nil
}
