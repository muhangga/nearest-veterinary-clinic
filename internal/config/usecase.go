package config

import (
	clinicRepository "github.com/muhangga/internal/repository/clinic"
	clinicUsecase "github.com/muhangga/internal/usecase/clinic"
)

type Usecase struct {
	ClinicUsecase clinicUsecase.ClinicUsecase
}

func InitUsecase(clinicRepository clinicRepository.ClinicRepository) Usecase {
	return Usecase{
		ClinicUsecase: clinicUsecase.NewClinicUsecase(clinicRepository),
	}
}
