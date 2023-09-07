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

type DreamImageQueue struct {
	gorm.Model
	DreamId uint   `json:"dream_id"`
	QueueId string `json:"queue_id"`
	Version string `json:"version"`
	Output  string `json:"output"`
	GetUrl  string `json:"get_url"`
	Status  string `json:"status"`
}

type DreamCreateInput struct {
	Content       string `json:"content" validate:"required"`
	GenerateImage bool   `json:"generate_image" validate:"required"`
}
