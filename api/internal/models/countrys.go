package models

import  "gorm.io/gorm"

type Country struct {
	gorm.Model
	Flag 	string	`json:"flag" gorm:"varchar(60);not null"`
	Name 	string	`json:"name" gorm:"varchar(20);not null"`
	Posts   []Post	`json:"posts" gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE"`
}