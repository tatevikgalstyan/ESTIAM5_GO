package sqlite

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"tp/db"
	"tp/model"
)

type SQLite struct {
	Conn *gorm.DB
}

func New(dbName string) *db.Storage {
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = conn.AutoMigrate(&model.Reservation{})
	if err != nil {
		panic(err)
	}
	err = conn.AutoMigrate(&model.Hotels{})
	if err != nil {
		panic(err)
	}
	err = conn.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	err = conn.AutoMigrate(&model.Country{})
	if err != nil {
		panic(err)
	}

	return &db.Storage{
		User:        &SQLite{Conn: conn},
		Reservation: &SQLite{Conn: conn},
		Hotel:       &SQLite{Conn: conn},
		Country:     &SQLite{Conn: conn},
	}
}

func (c *SQLite) GetUserByID(id string) (*model.User, error) {
	var u model.User
	err := c.Conn.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *SQLite) GetByEmail(email string) (*model.User, error) {
	var u model.User
	err := c.Conn.First(&u, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *SQLite) GetAllUser() ([]model.User, error) {
	var us []model.User
	err := c.Conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (c *SQLite) DeleteUserByID(id string) error {
	return c.Conn.Where("id = ?", id).Delete(&model.User{}).Error
}

func (c *SQLite) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.NewString()
	err := c.Conn.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *SQLite) UpdateUser(id string, data map[string]interface{}) (*model.User, error) {
	u := model.User{ID: id}
	err := c.Conn.Model(&u).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return c.GetUserByID(id)
}

//Reservation
func (c *SQLite) GetReservationByID(id string) (*model.Reservation, error) {
	var r model.Reservation
	err := c.Conn.First(&r, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &r, nil
}

//Trouver les reservation par client
func (c *SQLite) GetUserReservationByID(idUser string) ([]model.Reservation, error) {

	var urs []model.Reservation
	err := c.Conn.First(&urs, "id = ?", idUser).Error
	if err != nil {
		return nil, err
	}
	return urs, nil
}
func (c *SQLite) DeleteReservationByID(id string) error {
	return c.Conn.Where("id = ?", id).Delete(&model.Reservation{}).Error
}

//le client crée reservation en choisisant l'hotel donc on recoupere idhotel et idutilisation pour remplire la table reservation
func (c *SQLite) CreateReservation(r *model.Reservation, idUser, idHotel string) (*model.Reservation, error) {
	r.ID = uuid.NewString()
	r.IDHotel = idHotel
	r.IDUser = idUser
	err := c.Conn.Create(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

//Hotel permette chercher tous les hotels depuis bdd
func (c *SQLite) GetHotelByID(id string) (*model.Hotels, error) {
	var h model.Hotels
	err := c.Conn.First(&h, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (c *SQLite) GetAllHotels() ([]model.Hotels, error) {
	var hs []model.Hotels
	err := c.Conn.Find(&hs).Error
	if err != nil {
		return nil, err
	}
	return hs, nil
}

//l'utilisateur choisi le peys destination on cherche les hotel depuis bdd dans cette pays
func (c *SQLite) GetAllHotelsByIDPays(idCountry string) ([]model.Hotels, error) {
	var hts []model.Hotels
	err := c.Conn.First(&hts, "id = ?", idCountry).Error
	if err != nil {
		return nil, err
	}
	return hts, nil
}

//Pays destination
//Recoupere tous les destination depuis bdd
func (c *SQLite) GetAllCountry() ([]model.Country, error) {
	var cs []model.Country
	err := c.Conn.Find(&cs).Error
	if err != nil {
		return nil, err
	}
	return cs, nil
}

//Récouperer destination selectionée
func (c *SQLite) GetCountryByID(id string) (*model.Country, error) {
	var ct model.Country
	err := c.Conn.First(&ct, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &ct, nil
}

//Cree pays
func (c *SQLite) CreateCountry(ct *model.Country) (*model.Country, error) {
	ct.ID = uuid.NewString()
	err := c.Conn.Create(&ct).Error
	if err != nil {
		return nil, err
	}
	return ct, nil
}
