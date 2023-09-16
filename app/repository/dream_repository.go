package repository

import (
	"fmt"
	"github.com/yazilimcigenclik/dream-ai-backend/models"
	"gorm.io/gorm"
)

type DreamRepoInterface interface {
	DreamById(id int) (*models.Dream, error)
	Dreams() ([]*models.Dream, error)
	CreateDream(u *models.Dream) error
}

type DreamRepository struct {
	DB *gorm.DB
}

func NewDreamRepository(db *gorm.DB) *DreamRepository {
	return &DreamRepository{DB: db}
}

func (r *DreamRepository) DreamById(id int) (*models.Dream, error) {
	var dream models.Dream

	err := r.DB.Where("id = ?", id).First(&dream).Error

	if err != nil {
		fmt.Println("Error occurred while fetching dream on model", err)
		return nil, err
	}
	return &dream, nil
}

func (r *DreamRepository) Dreams() ([]*models.Dream, error) {
	var dreams []*models.Dream

	err := r.DB.Find(&dreams).Error

	if err != nil {
		fmt.Println("Error occurred while fetching dreams on model", err)
		return nil, err
	}

	return dreams, nil
}

func (r *DreamRepository) CreateDream(dream *models.Dream) error {
	err := r.DB.Create(&dream).Error

	if err != nil {
		fmt.Println("Error occurred while creating dream on model", err)
		return err
	}

	return nil
}
