package config

import (
	"database/sql"

	clinicRepository "github.com/muhangga/internal/repository/clinic"
)

type Repository struct {
	ClinicRepository clinicRepository.ClinicRepository
}

func InitRepository(db *sql.DB) Repository {
	return Repository{
		ClinicRepository: clinicRepository.NewClinicRepository(db),
	}
}
