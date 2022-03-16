package serviceHotels

import (
	"log"
	"net/http"

	"tp/db"

	//"github.com/dgkg/cmi/db"
	"github.com/gin-gonic/gin"
)

type ServiceHotels struct {
	db *db.Storage
}

func NewServiceHotel(db *db.Storage) *ServiceHotels {
	return &ServiceHotels{
		db: db,
	}
}

// go to Service folder.
func (s *ServiceHotels) GetHotel(c *gin.Context) {
	id := c.Param("id")
	h, err := s.db.Hotel.GetHotelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": h,
	})
}

func (s *ServiceHotels) GetAllHotels(c *gin.Context) {
	hs, err := s.db.Hotel.GetAllHotels()

	if err != nil {

		log.Println("service:", hs)
		//c.JSON(http.StatusInternalServerError, gin.H{
		//		"error": "error internal",
		//	})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"country": hs,
	})
}

// go to Service folder.
func (s *ServiceHotels) GetHotelByIDCountry(c *gin.Context) {
	id := c.Param("idCountry")
	h, err := s.db.Hotel.GetAllHotelsByIDPays(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hotels": h,
	})
}
