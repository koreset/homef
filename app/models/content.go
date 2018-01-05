package models

import (
	"github.com/jinzhu/gorm"
)

type Content struct {
	gorm.Model
	Type string
	Title string `gorm:not null`
	Body string
	Summary string
	Category Category
	CategoryID int
	Author Author
	AuthorID uint `gorm:not null`
	Images []Photo
}
