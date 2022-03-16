package moke

import (
	"errors"
	"tp/db"

	//"tp/db"

	"tp/model"

	"github.com/google/uuid"
)

var _ db.StorageUser = &Moke{}
var _ db.StorageCountry = &Moke{}
var _ db.StorageHotel = &Moke{}
var _ db.StorageReservation = &Moke{}

type Moke struct {
	listUser            map[string]*model.User
	listUserReservation map[string]*model.Reservation
	listCountry         map[string]*model.Country
	listHotels          map[string]*model.Hotels
}

func New() *db.Storage {
	return &db.Storage{
		User: &Moke{

			listUser: make(map[string]*model.User),
		},
		Country:     &Moke{listCountry: make(map[string]*model.Country)},
		Hotel:       &Moke{listHotels: make(map[string]*model.Hotels)},
		Reservation: &Moke{listUserReservation: make(map[string]*model.Reservation)},
	}
}

func (m *Moke) GetUserByID(id string) (*model.User, error) {
	u, ok := m.listUser[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	return u, nil
}
func (m *Moke) GetByEmail(email string) (*model.User, error) {
	for k := range m.listUser {
		if m.listUser[k].Email == email {
			return m.listUser[k], nil
		}
	}

	return nil, errors.New("db user: not found")
}

func (m *Moke) DeleteUserByID(id string) error {
	_, ok := m.listUser[id]
	if !ok {
		return errors.New("db user: not found")
	}
	delete(m.listUser, id)
	return nil
}
func (m *Moke) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.New().String()
	m.listUser[u.ID] = u
	return u, nil

}
func (m *Moke) UpdateUser(id string, data map[string]interface{}) (*model.User, error) {
	u, ok := m.listUser[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	if value, ok := data["first_name"]; ok {
		u.FirstName = value.(string)
	}
	if value, ok := data["last_name"]; ok {
		u.FirstName = value.(string)
	}
	return nil, nil
}

func (m *Moke) GetAllUser() ([]model.User, error) {
	us := make([]model.User, len(m.listUser))
	var i int
	for k := range m.listUser {
		if m.listUser[k] != nil {
			us[i] = *m.listUser[k]
		}
		i++
	}
	return us, nil
}

//reservation 

func (m *Moke) GetReservationByID(id string) (*model.Reservation, error) {
	ur, ok := m.listUserReservation[id]
	if !ok {
		return nil, errors.New("db Hotels: not found")
	}
	return ur, nil
}

func (m *Moke) GetUserReservationByID(idUser string) ([]model.Reservation, error) {
	ur := make([]model.Reservation, len(m.listUserReservation))
	var i int
	for k := range m.listUserReservation {
		if m.listUserReservation[k].IDUser == idUser {
			ur[i] = *m.listUserReservation[k]
		}
		i++
	}
	return ur, nil
}

func (m *Moke) DeleteReservationByID(id string) error {
	_, ok := m.listUserReservation[id]
	if !ok {
		return errors.New("dbreservation: not found")
	}
	delete(m.listUserReservation, id)
	return nil
}

func (m *Moke) CreateReservation(u *model.Reservation, idUser, idHotel string) (*model.Reservation, error) {
	u.ID = uuid.New().String()
	u.IDHotel=idHotel
	u.IDUser=idUser
	m.listUserReservation[u.ID] = u

	return u, nil

}

//Pays
func (m *Moke) GetAllCountry() ([]model.Country, error) {
	us := make([]model.Country, len(m.listCountry))
	var i int
	for k := range m.listCountry {
		if m.listCountry[k] != nil {
			us[i] = *m.listCountry[k]
		}
		i++
	}
	return us, nil
}
func (m *Moke) GetCountryByID(id string) (*model.Country, error) {
	u, ok := m.listCountry[id]
	if !ok {
		return nil, errors.New("db country: not found")
	}
	return u, nil
}
func (m *Moke) CreateCountry(ct *model.Country) (*model.Country, error) {
	ct.ID = uuid.New().String()
	m.listCountry[ct.ID] = ct
	return ct, nil

}

//hotels
func (m *Moke) GetHotelByID(id string) (*model.Hotels, error) {
	u, ok := m.listHotels[id]
	if !ok {
		return nil, errors.New("db Hotels: not found")
	}
	return u, nil
}

func (m *Moke) GetAllHotels() ([]model.Hotels, error) {
	us := make([]model.Hotels, len(m.listHotels))
	var i int
	for k := range m.listHotels {
		if m.listHotels[k] != nil {
			us[i] = *m.listHotels[k]
		}
		i++
	}
	return us, nil
}
func (m *Moke) GetAllHotelsByIDPays(idCountry string) ([]model.Hotels, error) {
	h := make([]model.Hotels, len(m.listHotels))
	var i int
	for k := range m.listHotels {
		if m.listHotels[k] != nil && m.listHotels[k].ID_pays == idCountry {
			h[i] = *m.listHotels[k]
		}
		i++
	}
	return h, nil

}
