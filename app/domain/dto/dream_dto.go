package dto

import (
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"time"
)

type GetDreamByIDRequest struct {
	ID uint `json:"id" binding:"required"`
}

type CreateDreamRequest struct {
	Content       string `json:"content" binding:"required"`
	GenerateImage bool   `json:"generate_image"`
}

type DreamDTO struct {
	ID            uint      `gorm:"primarykey" json:"id"`
	Title         string    `json:"title" gorm:"not null"`
	Content       string    `json:"content"`
	Explanation   string    `json:"explanation"`
	ImageUrl      *string   `json:"image_url"`
	GenerateImage bool      `json:"generate_image"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (dto *DreamDTO) FromDream(dream dao.Dream) {
	dto.ID = dream.ID
	dto.Title = dream.Title
	dto.Content = dream.Content
	dto.Explanation = dream.Explanation
	dto.ImageUrl = &dream.ImageUrl
	dto.GenerateImage = dream.GenerateImage
	dto.CreatedAt = dream.CreatedAt
	dto.UpdatedAt = dream.UpdatedAt
}
