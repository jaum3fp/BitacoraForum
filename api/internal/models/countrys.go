package models

import  "gorm.io/gorm"

type Country struct {
	gorm.Model
	Flag 	string	`json:"flag" gorm:"type:char(2);not null"`
	Name 	string	`json:"name" gorm:"type:varchar(50);not null"`
	Posts   []Post	`json:"posts" gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE"`
}