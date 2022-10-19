package repository

import (
	"database/sql"

	"github.com/muhangga/internal/entities"
	"github.com/muhangga/internal/repository/clinic/mysql"
)

type ClinicRepository interface {
	GetAllClinic() ([]entities.ClinicEntity, error)
}

type clinicRepository struct {
	clinicSQL mysql.ClinicSQL
}

func NewClinicRepository(db *sql.DB) ClinicRepository {
	return &clinicRepository{
		clinicSQL: mysql.NewClinicSQL(db),
	}
}
