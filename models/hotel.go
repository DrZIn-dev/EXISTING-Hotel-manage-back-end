package models

type Hotel struct {
	Base
	Name        string `json:"name" gorm:"unique" valid:"required"`
	Description string `json:"description" valid:"required"`
	Rooms       []HotelRoom
}

type HotelErrors struct {
	Err         bool   `json:"error"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type HotelRoom struct {
	Base
	HotelID     string `json"hotel_id" valid:"required"`
	Name        string `json:"name"  valid:"required"`
	Description string `json:"description" valid:"required"`
	Amount      uint   `json:"amount" valid:"required,int"`
	Price       uint   `json:"price" valid:"required,int"`
}
