package mysql

import (
	"database/sql"

	"github.com/muhangga/internal/entities"
)

type ClinicSQL interface {
	GetAllClinic() ([]entities.ClinicEntity , error)
	GetNearbyClinic(lat, long float64) ([]entities.ClinicEntity, error)
}

type clinicConnection struct {
	db *sql.DB
}

func NewClinicSQL(db *sql.DB) ClinicSQL {
	return &clinicConnection{db: db}
}
