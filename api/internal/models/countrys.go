package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Flag  string `gorm:"type:char(2);not null"`
	Name  string `gorm:"type:varchar(50);not null"`
	Posts []Post `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE"`
}
