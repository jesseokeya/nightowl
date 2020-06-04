package database

import (
	"github.com/jesseokeya/nightowl/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is gorm database object
var (
	DBCon *gorm.DB
)

type Database struct 

// Connect helps connect to the data base using gorm db object
func Connect(d structs.Database) (*gorm.DB, error) {
	return gorm.Open(d.sqlType, d.name)
}
