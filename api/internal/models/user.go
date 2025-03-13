package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string 	`json:"username" gorm:"type:varchar(50);not null;unique"` 
	Password string 	`json:"password" gorm:"type:binary(64);not null"` 
	Email    string 	`json:"email" gorm:"type:varchar(50);not null"` 
	Name     string 	`json:"name" gorm:"type:varchar(50)"` 
	Surnames string 	`json:"surnames" gorm:"type:varchar(126)"` 
	Posts    []Post 	`json:"posts" gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
}
