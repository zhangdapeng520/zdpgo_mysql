package zdpgo_mysql

import (
	"fmt"
	"testing"
)

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

// 测试根据ID列表删除
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


// 测试删除表格
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