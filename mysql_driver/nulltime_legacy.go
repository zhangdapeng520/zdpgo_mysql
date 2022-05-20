//go:build !go1.13
// +build !go1.13

package mysql_driver

import (
	"time"
)

// NullTime represents a time.Time that may be NULL.
// NullTime implements the Scanner interface so
// it can be used as a scan destination:
//
//  var nt NullTime
//  err := db.QueryRow("SELECT time FROM foo WHERE id=?", id).Scan(&nt)
//  ...
//  if nt.Valid {
//     // use nt.Time
//  } else {
//     // NULL value
//  }
//
// This NullTime implementation is not driver-specific
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// for internal use.
// the mysql package uses sql.NullTime if it is available.
// if not, the package uses mysql.NullTime.
type nullTime = NullTime // sql.NullTime is not available
