package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name"`
	SecCode    string `gorm:"type:varchar(10);uniqueIndex;not null" json:"secCode"`
	EDINETCode string `gorm:"type:varchar(10)" json:"edinetCode"`
	Industry   string `gorm:"type:varchar(100)" json:"industry"`
}