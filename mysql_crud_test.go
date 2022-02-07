package zdpgo_mysql

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Id     int
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

	sql := `
	INSERT INTO student(name, age, gender) VALUES(?, ?, ?);
	`
	add, err := m.Add(sql, "李四", 22, true)
	if err != nil {
		return
	}
	fmt.Println("插入数据成功：", add, err)
	m.Close()
}

func TestUpdate(t *testing.T) {
	m := prepareMysql()

	sql := `
	UPDATE student SET name = ? where id = ?;
	`
	update, err := m.Update(sql, "李四111", 1)
	fmt.Println(update, err)
	m.Close()
}

func TestFind(t *testing.T) {
	m := prepareMysql()

	student := &Student{}
	row, err := m.FindById("student", []string{"id", "name", "age", "gender"}, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)
	var (
		d      []string
		id     int
		name   string
		age    int
		gender bool
	)
	//err = row.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	err = row.Scan(&d)
	fmt.Println(d)
	fmt.Println(id, name, age, gender)
	if err != nil {
		fmt.Println("row.Scan 反射数据失败：", err)
		return
	}
	fmt.Println(student.Id, student.Name, student.Age, student.Gender)
	m.Close()
}

func TestFind1(t *testing.T) {
	m := prepareMysql()
	row := m.db.QueryRow("select * from student")
	fmt.Printf("row的数据类型：%v\n", reflect.TypeOf(row))
	student := &Student{}
	err := row.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		fmt.Println("row.Scan 反射数据失败：", err)
		return
	}
	fmt.Println(student.Id, student.Name, student.Age, student.Gender)
	m.Close()
}

func TestFind2(t *testing.T) {
	m := prepareMysql()
	student := Student{}
	row, err := m.QueryRow("select * from student")
	fmt.Printf("row的数据类型：%v\n", reflect.TypeOf(row))
	err = row.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		fmt.Println("row.Scan 反射数据失败：", err)
		return
	}
	fmt.Println(student.Id, student.Name, student.Age, student.Gender)
	m.Close()
}

func TestFind3(t *testing.T) {
	m := prepareMysql()
	stmt, err := m.db.Prepare("select * from student")
	row := stmt.QueryRow()
	fmt.Printf("row的数据类型：%v\n", reflect.TypeOf(row))
	student := &Student{}
	err = row.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		fmt.Println("row.Scan 反射数据失败：", err)
		return
	}
	fmt.Println(student.Id, student.Name, student.Age, student.Gender)
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

func TestDelete(t *testing.T) {
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
