package models

import (
	"github.com/jinzhu/gorm"
)

type Content struct {
	gorm.Model
	Type string
	Title string
	Body string
	Summary string
	Category Category
	CategoryID int
	Author Author
	AuthorID int
}
