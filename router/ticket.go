package router

import (
	"fmt"
	db "github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/database"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/models"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/util"
	"github.com/araddon/dateparse"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SetupTicketRoute() {

	private := TICKET.Group("/private")
	private.Use(util.SecureAuth())
	private.Post("/", CreateTicket)
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

	for i := 0; i < subDate; i++ {
		temp := checkIn.Add(time.Duration(i) * 24 * time.Hour)

		ticket.CheckIn, ticket.CheckOut = temp.String(), temp.String()
		ticket.CheckInTime, ticket.CheckOutTime = temp, temp

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
