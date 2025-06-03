package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     		string 	`gorm:"type:varchar(50);not null"`
	Description		string	`gorm:"type:text;not null"`
	Content   		string 	`gorm:"type:text;not null"`
	Views     		int64  	`gorm:"default:0"`
	OwnerID   		uint   	`gorm:"not null"`
	SuperID			*uint	`gorm:"index"`
	CountryAlpha 	string  `gorm:"type:char(2);not null"`
	Tags      		[]*Tag 	`gorm:"many2many:post_tags;"`
	Comments		[]*Post	`gorm:"foreignKey:SuperID"`
}
