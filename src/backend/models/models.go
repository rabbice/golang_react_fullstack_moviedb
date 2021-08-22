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
