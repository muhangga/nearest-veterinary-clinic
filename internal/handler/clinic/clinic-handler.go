package handler

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/muhangga/internal/usecase/clinic"
)

type ClinicHandler interface {
	GetAllClinic(c *gin.Context)
	GetNearbyClinic(c *gin.Context)
}

type clinicHandler struct {
	clinicUsecase usecase.ClinicUsecase
}

func NewClinicHandler(clinicUsecase usecase.ClinicUsecase) ClinicHandler {
	return &clinicHandler{clinicUsecase: clinicUsecase}
}
