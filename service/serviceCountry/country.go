package serviceCountry

import (
	"log"
	"net/http"

	"tp/model"

	"tp/db"
	//"github.com/dgkg/cmi/db"
	"github.com/gin-gonic/gin"
)

type ServiceCountry struct {
	db *db.Storage
}

func NewServiceCountry(db *db.Storage) *ServiceCountry {
	return &ServiceCountry{
		db: db,
	}
}

// go to Service folder.
func (s *ServiceCountry) GetCountry(c *gin.Context) {
	id := c.Param("id")
	cy, err := s.db.Country.GetCountryByID(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"country": cy,
	})
}

func (s *ServiceCountry) GetAllCountry(c *gin.Context) {
	ct, err := s.db.Country.GetAllCountry()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"country": ct,
	})
}

func (s *ServiceCountry) CreateCountry(c *gin.Context) {
	var p model.Country
	err := c.BindJSON(&p)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	_, err = s.db.Country.CreateCountry(&p)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": p,
	})
}
