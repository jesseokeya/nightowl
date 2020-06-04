package models

import (
	"github.com/jesseokeya/nightowl/database"
	"github.com/jesseokeya/nightowl/services"
	"github.com/jesseokeya/nightowl/structs"
)

// User type definition
type User struct {
	Model
	structs.User
}

func init() {
	database.DBCon.AutoMigrate(&User{})
}

// BeforeCreate hashes the password before user is craeted
func (u *User) BeforeCreate() *User {
	if u.Password != "" {
		u.Password = services.HashAndSalt([]byte(u.Password))
	}
	return u
}
