package data

/*
@Time : 2022/5/20 22:34
@Author : 张大鹏
@File : form
@Software: Goland2021.3.1
@Description: form 表单相关
*/

type Students struct {
	Id     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Age    uint   `db:"age" json:"age"`
	Gender string `db:"gender" json:"gender"`
}
