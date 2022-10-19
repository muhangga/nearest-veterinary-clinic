package repository

import "github.com/muhangga/internal/entities"

func (repository *clinicRepository) GetAllClinic() ([]entities.ClinicEntity, error) {
	return repository.clinicSQL.GetAllClinic()
}

func (repository *clinicRepository) GetNearbyClinic(lat, long float64) ([]entities.ClinicEntity, error){
	return repository.clinicSQL.GetNearbyClinic(lat, long)
}