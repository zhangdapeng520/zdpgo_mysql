package zdpgo_mysql

import (
	"fmt"
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
	// 表格名称
	tableName = "student"

	// 字段列表
	columns = []string{"name", "age", "gender"}

	// 字段的值
	values = []interface{}{"李四", 22, true}

	// 用于批量新增的学生数据
	studentData = [][]interface{}{
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
)

func prepareMysql() *Mysql {
	m := New(Config{
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
func TestMysql_Table(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 查询所有表格
	tables, err := m.Table.Find()
	if err != nil {
		t.Error(err)
	}
	t.Log(tables)

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_Execute(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 添加单条数据
	add, err := m.Execute.Add(tableName, columns, values)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 添加多条数据
	many, err := m.Execute.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(many)

	// 修改单条数据
	values = []interface{}{"李四111-修改", 22, true}
	id, err := m.Execute.UpdateById(tableName, columns, values, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)

	// 修改多条数据
	var ids = []int64{1, 2, 3, 4}
	byIds, err := m.Execute.UpdateByIds(tableName, columns, values, ids)
	if err != nil {
		t.Error(err)
	}
	t.Log(byIds)

	// 删除单条
	byId, err := m.Execute.DeleteById(tableName, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(byId)

	// 删除多条
	deleteByIds, err := m.Execute.DeleteByIds(tableName, 1, 2, 3, 4)
	if err != nil {
		t.Error(err)
	}
	t.Log(deleteByIds)

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_FindById(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Execute.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 根据ID查询
	student := Student{}
	row, err := m.Query.FindById("student", []string{"id", "name", "age", "gender"}, 1)
	if err != nil {
		t.Error(err)
	}

	// 将数据映射到结构体
	err = row.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		t.Error(err)
		return
	}

	// 打印查询到的内容
	fmt.Println(student.Id, student.Name, student.Age, student.Gender)

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_FindByIdToStruct(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Execute.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 批量查询
	var students []*Student
	err = m.Query.FindByIdToStruct("student", []string{"id", "name", "age", "gender"}, 1, &students)
	if err != nil {
		t.Error(err)
		return
	}

	// 遍历查询到的数据
	for _, student := range students {
		t.Log(student)
		t.Log(student.Id, student.Name, student.Age, student.Gender)
	}

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_FindByIdsToStruct(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Execute.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 批量查询
	var students []*Student
	err = m.Query.FindByIdsToStruct("student", []string{"id", "name", "age", "gender"}, []int64{1, 2, 3}, &students)
	if err != nil {
		t.Error(err)
		return
	}

	for _, student := range students {
		t.Log(student.Id, student.Name, student.Age, student.Gender)
	}

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_FindByPagesToStruct(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Execute.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 分页查询
	var students []*Student
	err = m.Query.FindByPageToStruct("student", []string{"id", "name", "age", "gender"}, 1, 20, &students)
	if err != nil {
		t.Error(err)
	}

	for _, student := range students {
		t.Log(student.Id, student.Name, student.Age, student.Gender)
	}

	// 删除表
	m.Table.Delete("student")
}

func TestFindIds(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Execute.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 根据ID列表查询
	rows, err := m.Query.FindByIds("student", []string{"id", "name"}, []int{1, 2, 3})
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()

	// 循环读取数据
	var students []Student
	for rows.Next() {
		student := &Student{}
		err = rows.Scan(&student.Id, &student.Name)
		if err != nil {
			t.Error(err)
		}
		students = append(students, *student)
	}
	t.Log(students)

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_FindByPage(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Execute.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	rows, err := m.Query.FindByPage("student", []string{"id", "name"}, 1, 20)
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

	// 删除表
	m.Table.Delete("student")
}
