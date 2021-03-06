package structs

// User type definition
type User struct {
	Email    string `gorm:"type:varchar(100);unique" json:"email" binding:"required,email"`
	Password string `gorm:"type:varchar(100)" json:"password,omitempty" binding:"required"`
	UserName string `gorm:"type:varchar(100)" json:"username,omitempty" binding:"required"`
}
