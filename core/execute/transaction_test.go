package execute

import (
	"testing"
)

const (
	// 创建学生表的SQL语句
	studentSql = `
	CREATE TABLE student(
		id BIGINT PRIMARY KEY auto_increment,
		name VARCHAR(24),
		age SMALLINT,
		gender BOOL
	) ENGINE = INNODB CHARSET = utf8;
	`
)

var (
	// 字段列表
	columns = []string{"name", "age", "gender"}

	// 字段的值
	values = []interface{}{"李四", 22, true}
)

func getExecute() *Execute {
	return NewExecute("root", "root", "127.0.0.1", 3306, "test")
}

// 测试事务
func TestMysql_TransAction(t *testing.T) {
	m := getExecute()

	m.Begin()

	// 创建表
	m.Execute(studentSql)

	// 添加成功
	add, err := m.Add("student", columns, values)
	if err != nil {
		m.Rollback()
		t.Error(err)
		return
	}
	t.Log("添加成功", add)

	// 删除失败
	id, err := m.DeleteById("student", 22)
	if err != nil {
		m.Rollback()
		t.Log(err)
	}
	t.Log("删除成功", id)

	m.Commit()
}

// 测试不使用事务能否添加成功
func TestMysql_NoTransAction(t *testing.T) {
	m := getExecute()

	// 创建表
	m.Execute(studentSql)

	// 添加成功
	add, err := m.Add("student", columns, values)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("添加成功", add)

	// 删除失败
	id, err := m.DeleteById("student", 22)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("删除成功", id)
}
