package database

import (
	"github.com/jinzhu/gorm"
	// configuration for sqlite dialect database
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is gorm database object
var (
	DBCon *gorm.DB
)

// Connect helps connect to the data base using gorm db object
func init() {
	var err error
	DBCon, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	DBCon.LogMode(true)
}
