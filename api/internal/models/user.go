package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"type:binary(64);not null"`
	Email    string `gorm:"type:varchar(50);not null"`
	Name     string `gorm:"type:varchar(50)"`
	Surnames string `gorm:"type:varchar(126)"`
	Posts    []Post `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
}
