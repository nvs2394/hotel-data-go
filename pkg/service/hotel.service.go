package service

import (
	"hotel-data-go/pkg/models"
	"hotel-data-go/pkg/repository"
)

type HotelService interface {
	GetHotels(hotelIDs []int, address string, city string) ([]models.Hotel, error)
}

type hotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) HotelService {
	return &hotelService{repo: repo}
}

func (s *hotelService) GetHotels(hotelIDs []int, address string, city string) ([]models.Hotel, error) {
	hotels, err := s.repo.GetAll(&hotelIDs, &address, &city)

	if err != nil {
		return nil, err
	}

	return hotels, nil
}
