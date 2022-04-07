package zdpgo_mysql

import (
	"fmt"
	"testing"
)

func prepareMysql() *Mysql {
	m := New(MysqlConfig{
		Debug:    true,
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	})
	return m
}

type Student struct {
	Id     int64
	Name   string
	Age    int
	Gender bool
}

// 测试创建表格
func TestCreateTable(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	var s string // sql 语句

	// 删除表
	s = "drop table if exists student;"
	m.Db.Execute(s)

	// 创建表
	s = `
	CREATE TABLE student(
		id BIGINT PRIMARY KEY auto_increment,
		name VARCHAR(24),
		age SMALLINT,
		gender BOOL
	) ENGINE = INNODB CHARSET = utf8;
	`
	m.Db.Execute(s)
}

// 测试添加数据
func TestMysql_Add(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 准备参数
	table := "student"
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"李四", 22, true}

	// 执行添加
	add, err := m.Db.Add(table, columns, values)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(add)
}

// 测试批量添加
func TestMysql_AddMany(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 准备参数
	table := "student"
	columns := []string{"name", "age", "gender"}
	values := [][]interface{}{
		{"李四12", 22, true},
		{"李四22", 22, true},
		{"李四33", 22, true},
		{"李四44", 22, true},
		{"李四22", 22, true},
		{"李四22", 22, true},
		{"李四23", 22, true},
		{"李四24", 22, true},
		{"李四32", 22, true},
		{"李四32", 22, true},
		{"李四33", 22, true},
		{"李四34", 22, true},
		{"李四42", 22, true},
		{"李四42", 22, true},
		{"李四43", 22, true},
		{"李四44", 22, true},
	}

	// 执行添加
	add, err := m.Db.AddMany(table, columns, values)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)
}

func TestMysql_UpdateById(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	columns := []string{"name", "age"}
	values := []interface{}{"李四111", 333}

	update, err := m.Db.UpdateById("student", columns, values, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(update)
}

// 测试批量更新
func TestMysql_UpdateByIds(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	columns := []string{"name", "age"}
	values := []interface{}{"李四111 是33岁", 33}
	ids := []int64{10, 12, 13}

	update, err := m.Db.UpdateByIds("student", columns, values, ids)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(update)
}

func TestMysql_FindById(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	student := Student{}
	row, err := m.Db.FindById("student", []string{"id", "name", "age", "gender"}, 1)
	if err != nil {
		t.Error(err)
	}

	err = row.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(student.Id, student.Name, student.Age, student.Gender)
}

func TestMysql_FindByIdToStruct(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	var students []*Student
	err := m.Db.FindByIdToStruct("student", []string{"id", "name", "age", "gender"}, 1, &students)
	if err != nil {
		t.Error(err)
		return
	}

	for _, student := range students {
		t.Log(student)
		t.Log(student.Id, student.Name, student.Age, student.Gender)
	}
}

func TestMysql_FindByIdsToStruct(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	var students []*Student
	err := m.Db.FindByIdsToStruct("student", []string{"id", "name", "age", "gender"}, []int64{1, 2, 3}, &students)
	if err != nil {
		t.Error(err)
		return
	}

	for _, student := range students {
		t.Log(student)
		t.Log(student.Id, student.Name, student.Age, student.Gender)
	}
}

func TestMysql_FindByPagesToStruct(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	var students []*Student
	err := m.Db.FindByPageToStruct("student", []string{"id", "name", "age", "gender"}, 1, 20, &students)
	if err != nil {
		t.Error(err)
		return
	}

	for _, student := range students {
		t.Log(student)
		t.Log(student.Id, student.Name, student.Age, student.Gender)
	}
}

func TestFindIds(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	rows, err := m.Db.FindByIds("student", []string{"id", "name"}, []int{1, 2, 3})
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()

	// 循环读取数据
	var students []Student
	for rows.Next() {
		student := &Student{}
		err := rows.Scan(&student.Id, &student.Name)
		if err != nil {
			t.Log("根据ID列表查询数据失败：", err)
			return
		}
		fmt.Println(student.Id, student.Name)
		students = append(students, *student)
	}
	t.Log(students)
}

func TestMysql_FindByPage(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	rows, err := m.Db.FindByPage("student", []string{"id", "name"}, 1, 20)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// 循环读取数据
	var students []Student
	for rows.Next() {
		student := &Student{}
		err = rows.Scan(&student.Id, &student.Name)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(student.Id, student.Name)
		students = append(students, *student)
	}
	t.Log(students)
}

func TestMySQL_DeleteById(t *testing.T) {
	m := prepareMysql()
	rows, err := m.Db.DeleteById("student", 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rows)
}

// 测试根据ID列表删除
func TestDeleteIds(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	rows, err := m.Db.DeleteByIds("student", 1, 2, 3, 4)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rows)
}

// 测试事务
func TestMysql_TransAction(t *testing.T) {
	m := prepareMysql()

	m.Db.Begin()

	// 添加成功
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"王五123事务测试", 22, true}
	add, err := m.Db.Add("student", columns, values)
	if err != nil {
		m.Db.Rollback()
		t.Error(err)
		return
	}
	t.Log("添加成功", add)

	// 删除失败
	id, err := m.Db.DeleteById("student", 22)
	if err != nil {
		m.Db.Rollback()
		t.Log(err)
		return
	}
	t.Log(id)

	m.Db.Commit()
}

// 测试不使用事务能否添加成功
func TestMysql_NoTransAction(t *testing.T) {
	m := prepareMysql()

	// 添加成功
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"王五123事务测试", 22, true}
	add, err := m.Db.Add("student", columns, values)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(add)

	// 删除失败
	id, err := m.Db.DeleteById("student", 22)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("删除成功", id)
}

// 测试删除表格
func TestDeleteTable(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	m.Db.DeleteTable("student")
}
