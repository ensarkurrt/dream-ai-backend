package dao

import (
	"time"
)

type Dream struct {
	ID          uint        `gorm:"primarykey" `
	Title       string      `json:"title"`
	Content     string      `json:"content" gorm:"not null"`
	Explanation string      `json:"explanation"`
	Status      DreamStatus `json:"status" gorm:"not null;default:pending"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DreamStatus string

const (
	Pending    DreamStatus = "pending"
	Processing DreamStatus = "processing"
	Completed  DreamStatus = "completed"
	Failed     DreamStatus = "failed"
)
