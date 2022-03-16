package db

import ("tp/model"
      
)

type Storage struct {
	User        StorageUser
	Reservation StorageReservation
	Hotel       StorageHotel
	Country     StorageCountry
}


type StorageUser interface {
	GetUserByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAllUser() ([]model.User, error)
	DeleteUserByID(id string) error
	CreateUser(u *model.User) (*model.User, error)
	UpdateUser(id string, data map[string]interface{}) (*model.User, error)
	//
}
type StorageReservation interface {
	GetReservationByID(id string) (*model.Reservation, error)
	GetUserReservationByID(idUser string) ([]model.Reservation, error)
	DeleteReservationByID(id string) error
	CreateReservation(u *model.Reservation, idUser,idHotel string) (*model.Reservation, error)
}

//
type StorageCountry interface {
	GetAllCountry() ([]model.Country, error)
	GetCountryByID(id string) (*model.Country, error)
	CreateCountry(u *model.Country) (*model.Country, error)
	//
}
type StorageHotel interface {
	GetHotelByID(id string) (*model.Hotels, error)
	GetAllHotels() ([]model.Hotels, error)
	GetAllHotelsByIDPays(idPays string) ([]model.Hotels, error)
}

//UpdateUserAccount(id string, data map[string]interface{}) (*model.UserAccount, error)
