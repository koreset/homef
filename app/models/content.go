package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct{
	gorm.Model
	Name string
}


type Content struct {
	gorm.Model
	Type string
	Title string
	Body string
	Summary string
	Category Category
	Author Author
	Image []string
}

type Author struct {
	gorm.Model
	Email string
	Name string
}
