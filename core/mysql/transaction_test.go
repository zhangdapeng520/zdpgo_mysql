package mysql

import (
	"testing"
)

// 测试事务
func TestMysql_TransAction(t *testing.T) {
	m := prepareMysql()

	m.Begin()

	// 添加成功
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"王五123事务测试", 22, true}
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
		return
	}
	t.Log("删除成功", id)

	m.Commit()
}

// 测试不使用事务能否添加成功
func TestMysql_NoTransAction(t *testing.T) {
	m := prepareMysql()

	// 添加成功
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"王五123事务测试", 22, true}
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
