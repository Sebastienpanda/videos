package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title   string
	Content string
}
