package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"column:id; primary_key" json:"id"`
	Name     string    `gorm:"column:name; type:varchar" json:"name"`
	Email    string    `gorm:"column:email; type:varchar; unique" json:"email"`
	Password string    `gorm:"column:password" json:"-"`
}

type UserRequest struct {
	Name     string `gorm:"column:name; type:varchar" json:"name"`
	Email    string `gorm:"column:email; type:varchar; unique" json:"email"`
	Password string `gorm:"column:password; type:varchar" json:"password"`
}
