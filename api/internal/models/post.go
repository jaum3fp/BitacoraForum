package models

import (
	"gorm.io/gorm"
)


type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(50);not null"` 
	Content string `json:"content" gorm:"type:text;not null"`
	OwnerID uint   `json:"owner_id" gorm:"not null"`
}