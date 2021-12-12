# zdpgo_mysql
Golang操作MySQL的快捷工具库

项目地址：https://github.com/zhangdapeng520/zdpgo_mysql

## 一、基本增删改查

### 1.1 创建表格
```go

// 测试创建表格
func TestCreateTable(t *testing.T){
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	sql := `
	CREATE TABLE student(
		id BIGINT PRIMARY KEY auto_increment,
		name VARCHAR(24)
	) ENGINE = INNODB CHARSET = utf8;
	`
	db.Execute(sql)
	defer db.Close()
}
```

### 1.2 添加数据
```go

func TestAdd(t *testing.T) {
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	sql := `
	INSERT INTO student(name) VALUES(?);
	`
	uid:=db.Add(sql, "李四")
	fmt.Println("插入数据成功：", uid)
	defer db.Close()
}
```

### 1.3 修改数据
```go
func TestUpdate(t *testing.T){
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	sql := `
	UPDATE student SET name = ? where id = ?;
	`
	uid:=db.Update(sql, "李四111", 1)
	fmt.Println("更新数据成功：", uid)
	defer db.Close()
}
```

### 1.4 根据ID查询数据
```go
func TestFind(t *testing.T){
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	student := &Student{}
	row:=db.Find("student", []string{"id", "name"}, 1)
	err := row.Scan(&student.Id, &student.Name)
	if err != nil{
		fmt.Println("根据ID查询数据失败：", err)
		return
	}
	fmt.Println(student.Id, student.Name)
	defer db.Close()
}
```

### 1.5 根据ID列表查询数据
```go
func TestFindIds(t *testing.T){
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()
	defer db.Close()

	rows:=db.FindIds("student", []string{"id", "name"},[]int{1,2,3})
	defer rows.Close()

	// 循环读取数据
	var students []Student
	for rows.Next(){
		student := &Student{}
		err := rows.Scan(&student.Id, &student.Name)
		if err != nil{
			fmt.Println("根据ID列表查询数据失败：", err)
			return
		}
		fmt.Println(student.Id, student.Name)
		students = append(students, *student)
	}
	fmt.Println(students)
}
```

### 1.6 分页查询数据
```go
func TestFindPages(t *testing.T){
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()
	defer db.Close()

	rows:=db.FindPages("student", []string{"id", "name"}, 1, 20)
	defer rows.Close()

	// 循环读取数据
	var students []Student
	for rows.Next(){
		student := &Student{}
		err := rows.Scan(&student.Id, &student.Name)
		if err != nil{
			fmt.Println("分页查询数据失败：", err)
			return
		}
		fmt.Println(student.Id, student.Name)
		students = append(students, *student)
	}
	fmt.Println(students)
}
```

### 1.7 根据ID删除数据
```go
func TestDelete(t *testing.T) {
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	flag:=db.Delete("student", 1)
	if flag{
		fmt.Println("删除数据成功：", flag)
	}
	defer db.Close()
}
```

### 1.8 根据ID列表删除数据
```go
func TestDeleteIds(t *testing.T) {
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	flag:=db.DeleteIds("student", 1, 2, 3, 4)
	if flag{
		fmt.Println("根据ID列表删除数据成功：", flag)
	}
	defer db.Close()
}
```

### 1.9 删除表格
```go
func TestDeleteTable(t *testing.T){
	db := Mysql{
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "test",
	}
	db.Init()

	db.DeleteTable("student")
	defer db.Close()
}
```
