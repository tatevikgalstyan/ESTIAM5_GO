package serviceReservation

import (
	"log"
	"net/http"

	"tp/db"
	"tp/model"

	//"github.com/dgkg/cmi/db"
	"github.com/gin-gonic/gin"
)

type ServiceReservation struct {
	db *db.Storage
}

func NewReservation(db *db.Storage) *ServiceReservation {
	return &ServiceReservation{
		db: db,
	}
}

// recoupere un reservation par son id
func (s *ServiceReservation) GetReservation(c *gin.Context) {
	id := c.Param("id")
	r, err := s.db.Reservation.GetReservationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"reservation": r,
	})
}

func (s *ServiceReservation) GetAllReservationByUsers(c *gin.Context) {
	id := c.Param("idUser")
	ur, err := s.db.Reservation.GetUserReservationByID(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"usersReservation": ur,
	})
}

// go to Service folder.

func (s *ServiceReservation) CreateReservation(c *gin.Context) {
	idu := c.Param("idUser")
	idh := c.Param("idHotel")
	
	var r model.Reservation
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	_, err = s.db.Reservation.CreateReservation(&r,idu,idh)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reservation": r,
	})
}

func (s *ServiceReservation) DeleteReservation(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}

	err := s.db.Reservation.DeleteReservationByID(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}
