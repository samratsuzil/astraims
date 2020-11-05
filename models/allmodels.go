package models

import ("github.com/jinzhu/gorm"
"time")

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Address string
	UserName    string
	Password    string
}

type Category struct {
	gorm.Model
	Name   string
	Description string
	Book 		[]Book `gorm:ForeignKey"CategoryID"`
}