package models

import "github.com/jinzhu/gorm"

type Movie struct {
	gorm.Model

	Title     string `gorm:"column:title" json:"title"`
	Year      int    `gorm:"column:year" json:"year"`
	Overview  string `gorm:"column:overview" json:"overview"`
	Directors string `gorm:"column:directors" json:"directors"`
	Budget    uint64 `gorm:"column:budget" json:"budget"`
	Gross     uint64 `gorm:"column:gross" json:"gross"`
}

func (Movie) TableName() string {
	return "test"
}

type User struct {
	gorm.Model
	FirstName string `gorm:"column:firstname" json:"first_name"`
	LastName  string `gorm:"column:lastname" json:"last_name"`
	Email     string `gorm:"column:email" json:"email,omitempty"`
	Password  string `json:"-"`
	LoggedIn  bool   `gorm:"column:loggedin" json:"loggedin"`
}

func (User) TableName() string {
	return "users"
}
