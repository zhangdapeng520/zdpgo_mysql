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

// 测试添加数据
func TestMysql_Add(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.Add(tableName, columns, values)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(add)

	// 删除表
	m.Table.Delete("student")
}

// 测试批量添加
func TestMysql_AddMany(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_UpdateById(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.Add(tableName, columns, values)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(add)

	// 修改数据
	columns = []string{"name", "age"}
	values = []interface{}{"李四111", 333}
	update, err := m.Db.UpdateById("student", columns, values, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(update)

	// 删除表
	m.Table.Delete("student")
}

// 测试批量更新
func TestMysql_UpdateByIds(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 批量更新
	columns = []string{"name", "age"}
	values1 := []interface{}{"李四111 是33岁", 33}
	ids := []int64{1, 2, 3, 4}
	update, err := m.Db.UpdateByIds("student", columns, values1, ids)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(update)

	// 删除表
	m.Table.Delete("student")
}

func TestMysql_FindById(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 根据ID查询
	student := Student{}
	row, err := m.Db.FindById("student", []string{"id", "name", "age", "gender"}, 1)
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
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 批量查询
	var students []*Student
	err = m.Db.FindByIdToStruct("student", []string{"id", "name", "age", "gender"}, 1, &students)
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
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 批量查询
	var students []*Student
	err = m.Db.FindByIdsToStruct("student", []string{"id", "name", "age", "gender"}, []int64{1, 2, 3}, &students)
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
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 分页查询
	var students []*Student
	err = m.Db.FindByPageToStruct("student", []string{"id", "name", "age", "gender"}, 1, 20, &students)
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
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 根据ID列表查询
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
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

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

	// 删除表
	m.Table.Delete("student")
}

func TestMySQL_DeleteById(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	rows, err := m.Db.DeleteById("student", 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rows)

	// 删除表
	m.Table.Delete("student")
}

// 测试根据ID列表删除
func TestDeleteIds(t *testing.T) {
	m := prepareMysql()
	defer m.Db.Close()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	rows, err := m.Db.DeleteByIds("student", 1, 2, 3, 4)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rows)

	// 删除表
	m.Table.Delete("student")
}

// 测试事务
func TestMysql_TransAction(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	m.Db.Begin()

	// 添加成功
	columns = []string{"name", "age", "gender"}
	values = []interface{}{"王五123事务测试", 22, true}
	add, err = m.Db.Add("student", columns, values)
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

	// 删除表
	m.Table.Delete("student")
}

// 测试不使用事务能否添加成功
func TestMysql_NoTransAction(t *testing.T) {
	m := prepareMysql()

	// 创建表
	m.Table.Add(studentSql)

	// 执行添加
	add, err := m.Db.AddMany(tableName, columns, studentData)
	if err != nil {
		t.Error(err)
	}
	t.Log(add)

	// 添加成功
	columns = []string{"name", "age", "gender"}
	values = []interface{}{"王五123事务测试", 22, true}
	add, err = m.Db.Add("student", columns, values)
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

	// 删除表
	m.Table.Delete("student")
}
