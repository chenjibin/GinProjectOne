package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required" gorm:"type:varchar(20);not null"`
	Password string `json:"password" binding:"required,min=6" gorm:"size:256;not null"`
	Phone    string `json:"phone" binding:"required,len=11" gorm:"type:varchar(11);unique;not null"`
}

type UserLogin struct {
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"required,len=11"`
}

type UserList struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
