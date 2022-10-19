package repository

import "github.com/muhangga/internal/entities"

func (repository *clinicRepository) GetAllClinic() ([]entities.ClinicEntity, error) {
	return repository.clinicSQL.GetAllClinic()
}
