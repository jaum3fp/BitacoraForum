package models

import (
	"gorm.io/gorm"
)


type Post struct {
	gorm.Model
	Title   	string 	`json:"title" gorm:"type:varchar(50);not null"` 
	Content 	string 	`json:"content" gorm:"type:text;not null"`
	Views   	int64 	`json:"views" gorm:"default:0"`
	OwnerID 	uint  	`json:"owner_id" gorm:"not null"`
	CountryID 	uint 	`json:"country_id" gorm:"not null"`
	Tags    	[]*Tag 	`json:"tags" gorm:"many2many:post_tags;"`
}