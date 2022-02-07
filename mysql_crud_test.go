package zdpgo_mysql

import (
	"fmt"
	"testing"
)

type Student struct {
	Id     int64
	Name   string
	Age    int
	Gender bool
}

// 测试创建表格
func TestCreateTable(t *testing.T) {
	m := prepareMysql()

	s := `
	CREATE TABLE student(
		id BIGINT PRIMARY KEY auto_increment,
		name VARCHAR(24),
		age SMALLINT,
		gender BOOL
	) ENGINE = INNODB CHARSET = utf8;
	`
	m.Execute(s)
	m.Close()
}

func TestAdd(t *testing.T) {
	m := prepareMysql()
	defer m.Close()

	// 准备参数
	table := "student"
	columns := []string{"name", "age", "gender"}
	values := []interface{}{"李四", 22, true}

	// 执行添加
	add, err := m.Add(table, columns, values)
	if err != nil {
		fmt.Println("执行添加失败：", err)
		return
	}
	fmt.Println("插入数据成功：", add, err)

}

func TestMysql_UpdateById(t *testing.T) {
	m := prepareMysql()

	columns := []string{"name", "age"}
	values := []interface{}{"李四111", 333}

	update, err := m.UpdateById("student", columns, values, 1)
	fmt.Println(update, err)
	m.Close()
}

func TestMysql_FindById(t *testing.T) {
	m := prepareMysql()

	student := Student{}
	row, err := m.FindById("student", []string{"id", "name", "age", "gender"}, 1)
	if err != nil {
		fmt.Println(err)
	}
	err = row.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(student.Id, student.Name, student.Age, student.Gender)
	m.Close()
}

func TestMysql_FindByIdToStruct(t *testing.T) {
	m := prepareMysql()

	var students []*Student
	err := m.FindByIdToStruct("student", []string{"id", "name", "age", "gender"}, 1, &students)
	fmt.Println(err)

	for _, student := range students {
		fmt.Println(student)
		fmt.Println(student.Id, student.Name, student.Age, student.Gender)
	}
	m.Close()
}

func TestMysql_FindByIdsToStruct(t *testing.T) {
	m := prepareMysql()

	var students []*Student
	err := m.FindByIdsToStruct("student", []string{"id", "name", "age", "gender"}, []int{1, 2, 3}, &students)
	fmt.Println(err)

	for _, student := range students {
		fmt.Println(student)
		fmt.Println(student.Id, student.Name, student.Age, student.Gender)
	}
	m.Close()
}

func TestMysql_FindByPagesToStruct(t *testing.T) {
	m := prepareMysql()

	var students []*Student
	err := m.FindByPagesToStruct("student", []string{"id", "name", "age", "gender"}, 1, 20, &students)
	fmt.Println(err)

	for _, student := range students {
		fmt.Println(student)
		fmt.Println(student.Id, student.Name, student.Age, student.Gender)
	}
	m.Close()
}

func TestFindIds(t *testing.T) {
	m := prepareMysql()
	defer m.Close()

	rows, err := m.FindByIds("student", []string{"id", "name"}, []int{1, 2, 3})
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// 循环读取数据
	var students []Student
	for rows.Next() {
		student := &Student{}
		err := rows.Scan(&student.Id, &student.Name)
		if err != nil {
			fmt.Println("根据ID列表查询数据失败：", err)
			return
		}
		fmt.Println(student.Id, student.Name)
		students = append(students, *student)
	}
	fmt.Println(students)
}

func TestFindPages(t *testing.T) {
	m := prepareMysql()
	defer m.Close()

	rows, err := m.FindPages("student", []string{"id", "name"}, 1, 20)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// 循环读取数据
	var students []Student
	for rows.Next() {
		student := &Student{}
		err := rows.Scan(&student.Id, &student.Name)
		if err != nil {
			fmt.Println("分页查询数据失败：", err)
			return
		}
		fmt.Println(student.Id, student.Name)
		students = append(students, *student)
	}
	fmt.Println(students)
}

func TestMySQL_DeleteById(t *testing.T) {
	m := prepareMysql()
	rows, err := m.DeleteById("student", 1)
	fmt.Println(rows, err)
}

// 测试根据ID列表删除
func TestDeleteIds(t *testing.T) {
	m := prepareMysql()
	rows, err := m.DeleteByIds("student", 1, 2, 3, 4)
	fmt.Println(rows, err)
	m.Close()
}

// 测试删除表格
func TestDeleteTable(t *testing.T) {
	m := prepareMysql()
	m.DeleteTable("student")
	m.Close()
}
