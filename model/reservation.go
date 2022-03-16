package model

type Reservation struct {
	ID          string `json:"id"`
	IDUser      string `json:"id_user"`
	IDHotel     string `json:"id_hotel"`
	NameHotel   string `json:"hotel_name"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
