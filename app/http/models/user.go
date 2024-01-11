package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
	Mobile   string `gorm:"column:valid" json:"valid"`
}

// TableName 删除表名带s
func (u *User) TableName() string {
	return "user"
}
