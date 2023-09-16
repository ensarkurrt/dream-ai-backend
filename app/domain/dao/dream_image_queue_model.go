package dao

import "gorm.io/gorm"

type DreamImageQueue struct {
	gorm.Model
	QueueID string `json:"queue_id"`
	Version string `json:"version"`
	Output  string `json:"output"`
	GetUrl  string `json:"get_url"`
	Status  string `json:"status"`
	DreamID uint   `json:"dream_id"`
	Dream   Dream  `json:"dream" gorm:"constraint:OnDelete:CASCADE;"`
}
