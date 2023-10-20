package dao

import (
	"time"
)

type Dream struct {
	ID            uint   `gorm:"primarykey" `
	Title         string `json:"title" gorm:"not null"`
	Content       string `json:"content"`
	Explanation   string `json:"explanation"`
	ImageUrl      string `json:"image_url"`
	ImagePrompt   string `json:"image_prompt"`
	GenerateImage bool   `json:"generate_image"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
