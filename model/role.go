package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"type:varchar(20);not null"`
}

type RoleList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
