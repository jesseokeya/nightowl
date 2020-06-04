package models

import (
	"github.com/jesseokeya/nightowl/database"
	"github.com/jesseokeya/nightowl/structs"
)

// Image type definition
type Image struct {
	Model
	structs.Image
}

func init() {
	database.DBCon.AutoMigrate(&Image{})
}
