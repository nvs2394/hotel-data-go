// hotel_service_test.go
package service

import (
	"errors"
	"hotel-data-go/pkg/models"
	"hotel-data-go/pkg/repository"
	"reflect"
	"testing"
	"time"
)

// Mock repository to simulate database interaction
type mockHotelRepo struct {
	hotels []models.Hotel
	err    error
}

func (m *mockHotelRepo) GetAll(hotelIDs *[]int, address *string, city *string) ([]models.Hotel, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.hotels, nil
}

func TestGetHotels(t *testing.T) {

	timeNow := time.Date(2024, time.September, 29, 10, 49, 28, 0, time.UTC)

	mockHotels := []models.Hotel{
		{
			ID:          1,
			Name:        "Hotel 1",
			Description: "Hotel 1 description",
			Address:     "123 Main St",
			City:        "HCM",
			State:       "HCM State",
			Country:     "Vietnam",
			PostalCode:  "700000",
			PhoneNumber: "0123456789",
			Email:       "info@hotel1.com",
			WebsiteURL:  "https://www.hotel1.com",
			Latitude:    10.762622,
			Longitude:   106.660172,
			Rating:      4.5,
			CreatedAt:   timeNow,
			UpdatedAt:   timeNow,
		},
	}

	tests := []struct {
		name      string
		hotelIDs  []int
		city      string
		address   string
		mockRepo  repository.HotelRepository
		expected  []models.Hotel
		expectErr bool
	}{
		{
			name:      "no filtering",
			hotelIDs:  []int{},
			city:      "",
			address:   "",
			mockRepo:  &mockHotelRepo{hotels: mockHotels},
			expected:  mockHotels,
			expectErr: false,
		},
		// {
		// 	name:          "filter by hotelIDs",
		// 	hotelIDs:      []int{1, 3},
		// 	destinationID: "",
		// 	mockRepo:      &mockHotelRepo{hotels: mockHotels},
		// 	expected: []models.Hotel{
		// 		{ID: 1, Name: "Hotel Paradise", DestinationID: "ABCDE"},
		// 		{ID: 3, Name: "Mountain Retreat", DestinationID: "ABCDE"},
		// 	},
		// 	expectErr: false,
		// },
		// {
		// 	name:          "filter by destinationID",
		// 	hotelIDs:      []int{},
		// 	destinationID: "ABCDE",
		// 	mockRepo:      &mockHotelRepo{hotels: mockHotels},
		// 	expected: []models.Hotel{
		// 		{ID: 1, Name: "Hotel Paradise", DestinationID: "ABCDE"},
		// 		{ID: 3, Name: "Mountain Retreat", DestinationID: "ABCDE"},
		// 	},
		// 	expectErr: false,
		// },
		// {
		// 	name:          "filter by hotelIDs and destinationID",
		// 	hotelIDs:      []int{1},
		// 	destinationID: "ABCDE",
		// 	mockRepo:      &mockHotelRepo{hotels: mockHotels},
		// 	expected: []models.Hotel{
		// 		{ID: 1, Name: "Hotel Paradise", DestinationID: "ABCDE"},
		// 	},
		// 	expectErr: false,
		// },
		{
			name:      "repository returns error",
			hotelIDs:  []int{},
			address:   "",
			city:      "",
			mockRepo:  &mockHotelRepo{err: errors.New("database error")},
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewHotelService(tt.mockRepo)
			hotels, err := svc.GetHotels(tt.hotelIDs, tt.address, tt.city)

			if tt.expectErr && err == nil {
				t.Fatalf("expected an error but got none")
			}

			if !tt.expectErr && err != nil {
				t.Fatalf("did not expect an error but got: %v", err)
			}

			if !reflect.DeepEqual(hotels, tt.expected) {
				t.Errorf("expected hotels: %+v, got: %+v", tt.expected, hotels)
			}
		})
	}
}
