package api

import (
	"errors"
	"hotel-data-go/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockHotelService struct {
	GetHotelsFn func(hotelIDs []int, address string, city string) ([]models.Hotel, error)
}

func (m MockHotelService) GetHotels(hotelIDs []int, address string, city string) ([]models.Hotel, error) {
	return m.GetHotelsFn(hotelIDs, address, city)
}

func performRequest(router *gin.Engine, method string, url string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	return w
}

func TestGetHotelsHandler(t *testing.T) {
	mock := &MockHotelService{}
	handler := NewHotelHandler(mock)

	gin.SetMode(gin.TestMode)
	route := gin.Default()

	route.GET("/hotels", handler.GetHotels)

	timeNow := time.Date(2024, time.September, 29, 10, 49, 28, 0, time.UTC)

	tests := []struct {
		name               string
		queryParams        string
		mockGetHotelsFn    func(hotelIDs []int, address string, city string) ([]models.Hotel, error)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:        "should return 200 status code",
			queryParams: "",
			mockGetHotelsFn: func(hotelIDs []int, address string, city string) ([]models.Hotel, error) {
				return []models.Hotel{
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
				}, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[{"id": 1, "name": "Hotel 1", "description": "Hotel 1 description", "address": "123 Main St", "city": "HCM", "state": "HCM State", "country": "Vietnam", "postal_code": "700000", "phone_number": "0123456789", "email": "info@hotel1.com", "website_url": "https://www.hotel1.com", "latitude": 10.762622, "longitude": 106.660172, "rating": 4.5, "created_at": "2024-09-29T10:49:28Z", "updated_at": "2024-09-29T10:49:28Z"}]`,
		},
		{
			name:        "should return 500 status code",
			queryParams: "",
			mockGetHotelsFn: func(hotelIDs []int, address string, city string) ([]models.Hotel, error) {
				return nil, errors.New("service error")
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.GetHotelsFn = tt.mockGetHotelsFn

			w := performRequest(route, "GET", "/hotels?"+tt.queryParams)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			assert.JSONEq(t, tt.expectedResponse, w.Body.String())
		})
	}
}
