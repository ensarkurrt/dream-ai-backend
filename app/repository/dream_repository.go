package repository

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"gorm.io/gorm"
)

type DreamRepository interface {
	FindAllDream() ([]dao.Dream, error)
	FindDreamById(id int) (dao.Dream, error)
	CreateDream(dream dao.Dream) (dao.Dream, error)
}

type DreamRepositoryImpl struct {
	db *gorm.DB
}

func (u *DreamRepositoryImpl) FindAllDream() ([]dao.Dream, error) {
	var dreams []dao.Dream

	err := u.db.Find(&dreams).Error

	if err != nil {
		fmt.Println("Error occurred while fetching dreams on model", err)
		return nil, err
	}

	return dreams, nil
}

func (u *DreamRepositoryImpl) FindDreamById(id int) (dao.Dream, error) {
	var dream dao.Dream

	err := u.db.Where("id = ?", id).First(&dream).Error

	if err != nil {
		log.Error("Got an error when find user by id. Error: ", err)
		return dao.Dream{}, err
	}

	return dream, nil
}

func (u *DreamRepositoryImpl) CreateDream(dream dao.Dream) (dao.Dream, error) {
	err := u.db.Create(&dream).Error

	if err != nil {
		fmt.Println("Error occurred while creating dream on model", err)
		return dao.Dream{}, err
	}

	return dream, nil
}

func DreamRepositoryInit(db *gorm.DB) *DreamRepositoryImpl {
	log.Info("Dream repository initialized")
	err := db.AutoMigrate(&dao.Dream{})

	if err != nil {
		log.Error("Got an error when migrate user. Error: ", err)
		return nil
	}

	return &DreamRepositoryImpl{
		db: db,
	}
}
