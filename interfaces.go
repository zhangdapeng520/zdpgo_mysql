package zdpgo_mysql

import "database/sql"

// Crud 增删改查的SQL接口
type Crud interface {
	Add(sql string, args ...interface{}) int64                          // 添加数据
	Delete(table string, id int) bool                                   // 根据ID删除数据
	DeleteIds(table string, ids ...int) bool                            // 根据ID列表删除数据
	Update(sql string, args ...interface{}) bool                        // 根据ID修改数据
	Find(table string, columns []string, id int) *sql.Row               // 根据ID查询数据
	FindIds(table string, columns []string, ids []int) *sql.Rows        // 根据ID列表查询数据
	FindPages(table string, columns []string, page, size int) *sql.Rows // 根据分页查询数据
}
