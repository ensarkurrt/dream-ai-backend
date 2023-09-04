package models

import "gorm.io/gorm"

/* declare gorm Dream Model */

type Dream struct {
	gorm.Model
	Title       string  `json:"title"`
	Content     string  `json:"content"`
	Explanation string  `json:"explanation"`
	ImageUrl    *string `json:"image_url"`
}

type DreamCreateInput struct {
	Content string `json:"content" validate:"required"`
}
