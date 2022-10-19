package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	helper "github.com/muhangga/internal/helper"
)

func (handler *clinicHandler) GetAllClinic(c *gin.Context) {
	clinics, err := handler.clinicUsecase.GetAllClinic()
	if err != nil {
		helper.ResponseErrorJSON(500, "Internal Server Error", "Somethining went wrong")
		return
	}

	res := helper.ResponseSuccessJSON("Success", clinics)
	c.JSON(200, res)
}

func (handler *clinicHandler) GetNearbyClinic(c *gin.Context) {
	lat := c.Request.URL.Query().Get("lat")
	long := c.Request.URL.Query().Get("long")

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		helper.ResponseErrorJSON(400, "Bad Request", "Latitude is required")
		return
	}

	longitude, err := strconv.ParseFloat(long, 64)
	if err != nil {
		helper.ResponseErrorJSON(400, "Bad Request", "Longitude is required")
		return
	}
	clinic, err := handler.clinicUsecase.GetNearbyClinic(latitude, longitude)
	if err != nil {
		helper.ResponseErrorJSON(500, "Internal Server Error", "Somethining went wrong")
		return
	}

	res := helper.ResponseSuccessJSON("Success", clinic)
	c.JSON(200, res)
}
