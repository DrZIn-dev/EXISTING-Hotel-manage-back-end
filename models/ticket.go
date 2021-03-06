package models

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	Base
	GroupID      uuid.UUID `json:"group_id" gorm:"autoIncrement:false"`
	UserID       string    `json:"user_id" valid:"required"`
	HotelID      string    `json:"hotel_id" valid:"required"`
	Hotel        Hotel     `json:"hotel" `
	CheckIn      string    `json:"check_in" valid:"required"`
	CheckOut     string    `json:"check_out" valid:"required"`
	CheckInTime  time.Time `json:"check_in_time" valid:"required"`
	CheckOutTime time.Time `json:"check_out_time" valid:"required"`
	Amount       uint      `json:"amount" valid:"required,int"`
}
