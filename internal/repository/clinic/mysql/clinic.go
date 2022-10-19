package mysql

import (
	"context"
	"time"

	"github.com/muhangga/internal/entities"
)

func (repository *clinicConnection) GetAllClinic() ([]entities.ClinicEntity, error) {
	queryStmt := `SELECT id, clinic_name, address, phone_number, latitude, longitude FROM clinic`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := repository.db.QueryContext(ctx, queryStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clinics []entities.ClinicEntity
	for rows.Next() {
		var clinic entities.ClinicEntity
		err := rows.Scan(&clinic.ID, &clinic.ClinicName, &clinic.Address, &clinic.PhoneNumber, &clinic.Latitude, &clinic.Longitude)
		if err != nil {
			return nil, err
		}
		clinics = append(clinics, clinic)
	}

	return clinics, nil
}

func (repository *clinicConnection) GetNearbyClinic(lat, long float64) ([]entities.ClinicEntity, error) {
	queryStmt := `SELECT id, clinic_name, address, phone_number, latitude, longitude, ( 6371 * acos( cos( radians(?) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians(?) ) + sin( radians(?) ) * sin( radians( latitude ) ) ) ) AS distance FROM clinic HAVING distance < 30 ORDER BY distance`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var clinic entities.ClinicEntity
	rows, err := repository.db.QueryContext(ctx, queryStmt, lat, long, lat)
	if err != nil {
		return nil, err
	}

	var clinics []entities.ClinicEntity
	for rows.Next() {
		err := rows.Scan(&clinic.ID, &clinic.ClinicName, &clinic.Address, &clinic.PhoneNumber, &clinic.Latitude, &clinic.Longitude, &clinic.Distance)
		if err != nil {
			return nil, err
		}
		clinics = append(clinics, clinic)
	}

	return clinics, nil
}
