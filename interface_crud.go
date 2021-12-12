package zdpgo_mysql

// 增删改查的SQL接口
type CrudSql interface{
	Add(sql string, args ...interface{}) int64// 添加数据
	Delete(table string, id int) bool // 根据ID数据
	DeleteIds(table string, ids ...int) bool // 根据ID列表数据
	Update(sql string, id int, args ...interface{})// 根据ID修改数据
	Find(sql string, id int)// 根据ID查询数据
	FindIds(sql string, id ...int)// 根据ID列表查询数据
	FindPages(sql string, page, size int)// 根据分页查询数据
}