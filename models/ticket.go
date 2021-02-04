package models

type Ticket struct {
	Base
	UserID      string    `json:"user_id" valid:"required"`
	HotelID     string    `json:"hotel_id" valid:"required"`
	HotelRoomID string    `json:"hotel_room_id" valid:"required"`
	Hotel       Hotel     `json:"hotel" `
	HotelRoom   HotelRoom `json:"hotel_room" `
	CheckIn     string    `json:"check_in" valid:"required"`
	CheckOut    string    `json:"check_out" valid:"required"`
	Amount      uint      `json:"amount" valid:"required,int"`
}
