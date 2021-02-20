package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code  string `gorm:"column:code" json:"code"`
	Price int    `gorm:"column:price" json:"price"`
}

func (Product) TableName() string {
	return "product"
}
