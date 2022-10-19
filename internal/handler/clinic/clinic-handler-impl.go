package handler

import "github.com/gin-gonic/gin"

func (handler *clinicHandler) GetAllClinic(c *gin.Context) {
	clinics, err := handler.clinicUsecase.GetAllClinic()
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data":    clinics,
	})
}
