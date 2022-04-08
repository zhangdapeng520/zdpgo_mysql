package table

import "testing"

func getTable() *Table {
	return NewTable("root", "root", "127.0.0.1", 3306, "test")
}

// 测试查询所有表格
func TestTable_Find(t *testing.T) {
	tb := getTable()
	tables, err := tb.Find()
	if err != nil {
		t.Error(err)
	}
	t.Log(tables)
}
