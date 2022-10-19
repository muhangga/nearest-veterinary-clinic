package usecase

import "github.com/muhangga/internal/entities"

func (usecase *clinicUsecase) GetAllClinic() ([]entities.ClinicEntity, error) {
	return usecase.clinicRepository.GetAllClinic()
}
