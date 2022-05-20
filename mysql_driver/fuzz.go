//go:build gofuzz
// +build gofuzz

package mysql_driver

import (
	"database/sql"
)

func Fuzz(data []byte) int {
	db, err := sql.Open("mysql", string(data))
	if err != nil {
		return 0
	}
	db.Close()
	return 1
}
