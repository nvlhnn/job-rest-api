package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"varchar(255); not null" json:"name"`
	Password  string   `gorm:"->;<-;not null" json:"-"`
}