package config

import (
	clinicHandler "github.com/muhangga/internal/handler/clinic"
	clinicUsecase "github.com/muhangga/internal/usecase/clinic"
)

type Handler struct {
	ClinicHandler clinicHandler.ClinicHandler
}

func InitHandler(clinicUsecase clinicUsecase.ClinicUsecase) Handler {
	return Handler{
		ClinicHandler: clinicHandler.NewClinicHandler(clinicUsecase),
	}
}
