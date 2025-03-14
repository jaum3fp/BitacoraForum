package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(50);not null"`
	Posts []*Post `gorm:"many2many:post_tags;"`
}
