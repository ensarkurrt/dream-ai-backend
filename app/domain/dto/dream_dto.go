package dto

import (
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"time"
)

type GetDreamByIDRequest struct {
	ID uint `json:"id" binding:"required"`
}

type CreateDreamRequest struct {
	Content string `json:"content" binding:"required"`
}

type DreamDTO struct {
	ID          uint            `gorm:"primarykey" json:"id"`
	Title       string          `json:"title"`
	Content     string          `json:"content" gorm:"not null"`
	Explanation string          `json:"explanation"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Status      dao.DreamStatus `json:"status"`
}

func (dto *DreamDTO) FromDream(dream dao.Dream) {
	dto.ID = dream.ID
	dto.Title = dream.Title
	dto.Content = dream.Content
	dto.Explanation = dream.Explanation
	dto.CreatedAt = dream.CreatedAt
	dto.UpdatedAt = dream.UpdatedAt
	dto.Status = dream.Status
}
