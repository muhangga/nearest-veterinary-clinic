package usecase

import (
	"github.com/muhangga/internal/entities"
	repository "github.com/muhangga/internal/repository/clinic"
)

type ClinicUsecase interface {
	GetAllClinic() ([]entities.ClinicEntity, error)
	GetNearbyClinic(lat, long float64) ([]entities.ClinicEntity, error)
}

type clinicUsecase struct {
	clinicRepository repository.ClinicRepository
}

func NewClinicUsecase(clinicRepository repository.ClinicRepository) ClinicUsecase {
	return &clinicUsecase{clinicRepository: clinicRepository}
}
