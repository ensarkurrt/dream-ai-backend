package dao

import "gorm.io/gorm"

type Dream struct {
	gorm.Model
	Title         string  `json:"title"`
	Content       string  `json:"content"`
	Explanation   string  `json:"explanation"`
	ImageUrl      *string `json:"image_url"`
	GenerateImage bool    `json:"generate_image"`
}
