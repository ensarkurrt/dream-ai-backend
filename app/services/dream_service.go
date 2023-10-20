package services

import (
	log "github.com/sirupsen/logrus"
	"github.com/yazilimcigenclik/dream-ai-backend/app/constants"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dto"
	"github.com/yazilimcigenclik/dream-ai-backend/app/pkg"
	"github.com/yazilimcigenclik/dream-ai-backend/app/repository"
)

type DreamService interface {
	GetAllDream() []dto.DreamDTO
	GetDreamById(id int) dto.DreamDTO
	CreateDream(request dto.CreateDreamRequest) dto.DreamDTO
}

type DreamServiceImpl struct {
	dreamRepository repository.DreamRepository
}

func (u DreamServiceImpl) GetDreamById(id int) dto.DreamDTO {
	data, err := u.dreamRepository.FindDreamById(id)

	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}

	var dreamDTO dto.DreamDTO
	dreamDTO.FromDream(data)

	return dreamDTO
}

func (u DreamServiceImpl) GetAllDream() []dto.DreamDTO {
	dreams, err := u.dreamRepository.FindAllDream()

	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}
	var dreamDTOs []dto.DreamDTO

	for _, dream := range dreams {
		var dreamDTO dto.DreamDTO
		dreamDTO.FromDream(dream)
		dreamDTOs = append(dreamDTOs, dreamDTO)
	}

	return dreamDTOs
}

func (u DreamServiceImpl) CreateDream(request dto.CreateDreamRequest) dto.DreamDTO {

	dream := dao.Dream{
		Content:       request.Content,
		GenerateImage: request.GenerateImage,
	}

	data, err := u.dreamRepository.CreateDream(dream)

	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}

	var dreamDTO dto.DreamDTO
	dreamDTO.FromDream(data)

	return dreamDTO
}

func DreamServiceInit(dreamRepository repository.DreamRepository) *DreamServiceImpl {
	return &DreamServiceImpl{
		dreamRepository: dreamRepository,
	}
}
