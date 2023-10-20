package dao

import "time"

type DreamQueue struct {
	ID        uint        `gorm:"primarykey" `
	DreamID   uint        `json:"dream_id" gorm:"not null"`
	Dream     Dream       `gorm:"foreignKey:DreamID"`
	Status    DreamStatus `json:"status" gorm:"not null;default:pending"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
