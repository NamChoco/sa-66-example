package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Profile   string `gorm:"type:longtext"`

	// GenderID ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id"`
}

type Gender struct {
	gorm.Model
	Name string
	User []User `gorm:"foreignKey:GenderID"`
}

type Member struct {
	gorm.Model
	Username	string  `gorm:"uniqueIndex" valid:"required~Username is required"`
	Password	string
	Email       string  `gorm:"uniqueIndex" valid:"required~Email is required, email~Email is invalid"`
}