package main

import (
	"fmt"
	"log"
	"tp/db"
	"tp/db/moke"
	"tp/db/sqlite"
	"tp/service/serviceCountry"
	"tp/service/serviceHotels"
	"tp/service/serviceReservation"
	"tp/service/serviceUser"
	"tp/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	EnvType    string
	ListenPort string
	SecretKey  []byte
}

var config Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	config.EnvType = viper.GetString("EnvType")
	config.SecretKey = []byte(viper.GetString("SecretKey"))
	config.ListenPort = viper.GetString("ListenPort")
}

func main() {
	r := gin.Default()
	var db *db.Storage

	//db := moke.New()
	log.Println("ENV:", config.EnvType)
	if config.EnvType == "dev" {
		log.Println("create Moke DB")
		db = moke.New()
	} else {
		log.Println("create SQLite DB")
		db = sqlite.New("storage.db")
	}

	secureJWT := util.MiddlJWT(config.SecretKey)
	s := serviceUser.New(db, config.SecretKey)

	sh := serviceHotels.NewServiceHotel(db)
	sr := serviceReservation.NewReservation(db)
	sc := serviceCountry.NewServiceCountry(db)

	r.GET("countrys", sc.GetAllCountry)
	r.GET("countrys/:id", sc.GetCountry)
	r.POST("countrys", sc.CreateCountry)

	//
	r.GET("/users/:id", s.GetUser)
	r.POST("/user", s.CreateUser)
	r.GET("/users", s.GetAllUser)
	r.DELETE("/user/:id", secureJWT, s.DeleteUser)
	r.POST("/login", s.Login)
	//
	r.GET("hotels", sh.GetAllHotels)
	r.GET("hotels/:id", sh.GetHotel)
	r.GET("users/hotels/:iduser", sh.GetHotelByIDCountry)
	//
	r.GET("reservations/:id", sr.GetReservation)
	r.GET("reservations/hotels/:idUser", sr.GetAllReservationByUsers)
	r.POST("reservations/:idUser/:idHotel", sr.CreateReservation)
	r.DELETE("reservations/:id", secureJWT, sr.DeleteReservation)

	r.Run(":" + config.ListenPort)

}
