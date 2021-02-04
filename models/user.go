package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Base
	Email     string `json:"email" gorm:"unique" valid:"email,required"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
	Tickets   []Ticket
}

type UserResponse struct {
	Base
	Email     string `json:"email" gorm:"unique" valid:"email,required"`
	Username  string `json:"username" gorm:"unique"`
	BirthDate string `json:"birth_date"`
	Tickets   []Ticket
}

type UserErrors struct {
	Err       bool   `json:"error"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}

type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}
