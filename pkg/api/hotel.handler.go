package api

import (
	"hotel-data-go/pkg/models"
	"hotel-data-go/pkg/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	service service.HotelService
}

func NewHotelHandler(service service.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

func getHotelIds(hotelIDsParam string) ([]int, error) {

	var hotelIDs []int
	if hotelIDsParam != "" {
		for _, idStr := range strings.Split(hotelIDsParam, ",") {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return nil, err
			}

			hotelIDs = append(hotelIDs, id)
		}
	}
	return hotelIDs, nil
}

// GetHotels godoc
// @Summary Get list of hotels
// @Description Get a list of hotels filtered by hotel_ids or destination_id
// @Tags hotels
// @Accept  json
// @Produce  json
// @Param hotel_ids query []int false "List of hotel IDs to filter by"
// @Param address query string false "Address to filter by"
// @Param city query string false "City to filter by"
// @Success 200 {array} models.Hotel
// @Failure 500 500 {object} models.ErrorResponse
// @Router /hotels [get]
func (h *HotelHandler) GetHotels(c *gin.Context) {
	hotelIDsParam := c.Query("hotel_ids")

	address := c.Query("address")

	city := c.Query("city")

	hotelIDs, _ := getHotelIds(hotelIDsParam)

	hotels, err := h.service.GetHotels(hotelIDs, address, city)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, hotels)
}
