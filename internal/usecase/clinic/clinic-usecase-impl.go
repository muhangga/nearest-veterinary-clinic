package usecase

import "github.com/muhangga/internal/entities"

func (usecase *clinicUsecase) GetAllClinic() ([]entities.ClinicEntity, error) {
	return usecase.clinicRepository.GetAllClinic()
}

func (usecase *clinicUsecase) GetNearbyClinic(lat, long float64) ([]entities.ClinicEntity, error) {
	return usecase.clinicRepository.GetNearbyClinic(lat, long)
}
