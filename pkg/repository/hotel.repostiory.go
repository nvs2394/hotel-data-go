package repository

import (
	"hotel-data-go/pkg/models"

	"gorm.io/gorm"
)

type HotelRepository interface {
	GetAll(hotelIDs *[]int, address *string, city *string) ([]models.Hotel, error)
}

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *hotelRepository {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) GetAll(hotelIDs *[]int, address *string, city *string) ([]models.Hotel, error) {
	var hotels []models.Hotel
	query := r.db

	if len(*hotelIDs) > 0 {
		query = query.Where("id IN ?", *hotelIDs)
	}

	if address != nil && *address != "" {
		query = query.Where("address ILIKE ?", "%"+*address+"%")
	}
	if city != nil && *city != "" {
		query = query.Where("city ILIKE ?", "%"+*city+"%")
	}

	if err := query.Find(&hotels).Error; err != nil {
		return nil, err
	}

	return hotels, nil
}
