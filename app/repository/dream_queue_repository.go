package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"gorm.io/gorm"
)

type DreamQueueRepository interface {
	FindById(id uint) (dao.DreamQueue, error)
	FindByDreamId(dreamId uint) ([]dao.DreamQueue, error)
	Create(dreamQueue dao.DreamQueue) (dao.DreamQueue, error)
	Update(dreamQueue dao.DreamQueue) (dao.DreamQueue, error)
}

type DreamQueueRepositoryImpl struct {
	db *gorm.DB
}

func (d *DreamQueueRepositoryImpl) FindById(id uint) (dao.DreamQueue, error) {
	var dreamQueue dao.DreamQueue
	err := d.db.First(&dreamQueue, id).Error
	if err != nil {
		return dao.DreamQueue{}, err
	}

	return dreamQueue, nil
}

func (d *DreamQueueRepositoryImpl) FindByDreamId(dreamId uint) ([]dao.DreamQueue, error) {
	var dreamQueues []dao.DreamQueue
	err := d.db.Where("dream_id = ?", dreamId).Find(&dreamQueues).Error
	if err != nil {
		return nil, err
	}

	return dreamQueues, nil
}

func (d *DreamQueueRepositoryImpl) Create(dreamQueue dao.DreamQueue) (dao.DreamQueue, error) {
	err := d.db.Create(&dreamQueue).Error
	if err != nil {
		log.Println("Error occurred while creating dream queue on model", err)
		return dao.DreamQueue{}, err
	}

	return dreamQueue, nil
}

func (d *DreamQueueRepositoryImpl) Update(dreamQueue dao.DreamQueue) (dao.DreamQueue, error) {
	err := d.db.Save(&dreamQueue).Error
	if err != nil {
		return dao.DreamQueue{}, err
	}

	return dreamQueue, nil
}

func NewDreamQueueRepository(db *gorm.DB) *DreamQueueRepositoryImpl {
	log.Info("Dream Queue repository initialized")
	err := db.AutoMigrate(&dao.DreamQueue{})

	if err != nil {
		log.Error("Got an error when migrate dream. Error: ", err)
		return nil
	}

	return &DreamQueueRepositoryImpl{
		db: db,
	}
}
