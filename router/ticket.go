package router

import (
	"fmt"
	db "github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/database"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/models"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/util"
	"github.com/araddon/dateparse"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func SetupTicketRoute() {

	private := TICKET.Group("/private")
	private.Use(util.SecureAuth())
	private.Post("/", CreateTicket)
	private.Get("/", GetTicket)
}

func GetTicket(c *fiber.Ctx) error {
	type TicketResponse struct {
		GroupID  uuid.UUID `json:"group_id" gorm:"autoIncrement:false"`
		HotelID  string    `json:"hotel_id" valid:"required"`
		CheckIn  time.Time `json:"check_in" valid:"required"`
		CheckOut time.Time `json:"check_out" valid:"required"`
		Amount   uint      `json:"amount" valid:"required,int"`
	}

	type HotelList struct {
		HotelID string `json:"hotel_id" valid:"required"`
	}

	ticket := new([]TicketResponse)

	id := c.Locals("id")
	db.DB.Model(&models.Ticket{}).Select("group_id, MAX(hotel_id) as hotel_id,MIN(check_in_time)  as check_in, MAX(check_out_time) as check_out,MAX(amount) as amount").Where("user_id = ? ", id).Group("group_id").Find(&ticket)
	hotelList := new([]HotelList)
	count := db.DB.Model(&models.Ticket{}).Where("user_id = ? ", id).Group("hotel_id").Find(&hotelList).RowsAffected

	headers := make([]string, count)
	for i := 0; i < int(count); i++ {
		headers[i] = (*hotelList)[i].HotelID
	}

	hotels := new([]models.Hotel)
	db.DB.Model(&models.Hotel{}).Find(&hotels, headers)
	return c.JSON(fiber.Map{
		"tickets": ticket,
		"hotels":  hotels,
	})
}

func CreateTicket(c *fiber.Ctx) error {
	ticket := new(models.Ticket)

	err := c.BodyParser(ticket)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	checkIn, err := dateparse.ParseLocal(ticket.CheckIn)
	if err != nil {
		panic(err.Error())
	}

	checkOut, err := dateparse.ParseLocal(ticket.CheckOut)
	if err != nil {
		panic(err.Error())
	}

	subDate := int((checkOut.Sub(checkIn).Hours() / 24) + 1)

	id := c.Locals("id")
	idStr := fmt.Sprintf("%v", id)
	ticket.UserID = idStr
	groupId := uuid.New()
	for i := 0; i < subDate; i++ {
		temp := checkIn.Add(time.Duration(i) * 24 * time.Hour)

		ticket.CheckIn, ticket.CheckOut = temp.String(), temp.String()
		ticket.CheckInTime, ticket.CheckOutTime = temp, temp
		ticket.GroupID = groupId
		err := db.DB.Create(&ticket).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"general": "Something went wrong, please try again later. 😕",
			})
		}
	}

	return c.JSON(fiber.Map{"checkIn": checkIn, "checkOut": checkOut})
}
