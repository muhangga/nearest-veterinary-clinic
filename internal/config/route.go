package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRoute() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := OpenDB(os.Getenv("MYSQL_URL"), true)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	repository := InitRepository(db)
	usecase := InitUsecase(repository.ClinicRepository)
	handler := InitHandler(usecase.ClinicUsecase)

	r := gin.Default()

	clinicsRoute := r.Group("/api/v1/clinics")
	{
		clinicsRoute.GET("/", handler.ClinicHandler.GetAllClinic)
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"message": "Not Found",
		})
	})

	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(405, gin.H{
			"message": "Method Not Allowed",
		})
	})

	r.Run(":8080")
}
