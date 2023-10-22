package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"github.com/yazilimcigenclik/dream-ai-backend/app/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (dao.User, error)
	FindById(id uint) (dao.User, error)
	Create(user dao.User) (dao.User, error)
}

func (repo *UserRepositoryImpl) FindById(id uint) (dao.User, error) {
	var user dao.User
	result := repo.db.First(&user, id)

	if result.Error != nil {
		return dao.User{}, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) FindByUsername(username string) (dao.User, error) {
	var user dao.User
	result := repo.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return dao.User{}, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Create(user dao.User) (dao.User, error) {

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return dao.User{}, err
	}

	user.Password = hash
	result := repo.db.Create(&user)

	if result.Error != nil {
		return dao.User{}, result.Error
	}

	return user, nil

}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	log.Info("User repository initialized")
	err := db.AutoMigrate(&dao.User{})

	if err != nil {
		log.Error("Got an error when migrate user. Error: ", err)
		return nil
	}

	return &UserRepositoryImpl{
		db,
	}
}
