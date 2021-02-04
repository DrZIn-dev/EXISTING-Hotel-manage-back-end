package models

type Hotel struct {
	ID int `json:"id" gorm:"autoIncrement:true;unique"`
	Base
	Name        string  `json:"name" gorm:"unique" valid:"required"`
	Description string  `json:"description" valid:"required"`
	Lat         float64 `json:"lat" gorm:"type:decimal(10,8)"`
	Lng         float64 `json:"lng" gorm:"type:decimal(11,8)"`
	Room        int     `json:"room" valid:"required"`
	Price       int     `json:"price" valid:"required"`
}

type HotelErrors struct {
	Err  bool   `json:"error"`
	Name string `json:"name"`
}
