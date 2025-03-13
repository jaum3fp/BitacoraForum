package models

import (
	"gorm.io/gorm"
)


type Tag struct {
	gorm.Model
	Name 	string 		`json:"name" gorm:"type:varchar(50);not null"`
	Posts 	[]*Post 	`json:"posts" gorm:"many2many:post_tags;"`
}