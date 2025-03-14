package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string `gorm:"type:varchar(50);not null"`
	Content   string `gorm:"type:text;not null"`
	Views     int64  `gorm:"default:0"`
	OwnerID   uint   `gorm:"not null"`
	CountryID uint   `gorm:"not null"`
	Tags      []*Tag `gorm:"many2many:post_tags;"`
}
