package router

import (
	db "github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/database"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/models"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/util"
	"github.com/araddon/dateparse"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

func SetupHotelRoutes() {
	HOTEL.Get("/", GetHotel)
	HOTEL.Group("/search", Search)
	private := HOTEL.Group("/private")
	private.Use(util.SecureAuth())
	private.Post("/", CreateHotel)

}

func GetHotel(c *fiber.Ctx) error {

	cursor := c.Query("cursor")

	h := new(models.Hotel)

	res := db.DB.Model(&models.Hotel{}).Where("uuid = ? ", cursor).First(&h)
	if res.RowsAffected <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "general": "Cannot find the Hotel"})
	}

	return c.JSON(h)
}

func CreateHotel(c *fiber.Ctx) error {
	hotel := new(models.Hotel)

	err := c.BodyParser(hotel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	_, err = govalidator.ValidateStruct(hotel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	hotelCount := db.DB.Where(&models.Hotel{Name: hotel.Name}).First(new(models.Hotel)).RowsAffected
	if hotelCount > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Hotel is already registered",
		})
	}

	err = db.DB.Create(&hotel).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later. 😕",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": hotel.UUID,
	})

}

func Search(c *fiber.Ctx) error {
	name := c.Query("name")
	nameStr := strings.ReplaceAll(name, " ", " | ")

	cursor := c.Query("cursor")
	cursorInt, _ := strconv.Atoi(cursor)

	limit := c.Query("limit")
	limitInt, _ := strconv.Atoi(limit)

	checkIn := c.Query("checkIn")
	checkInTime, err := dateparse.ParseLocal(checkIn)
	if err != nil {
		panic(err.Error())
	}

	checkOut := c.Query("checkOut")
	checkOutTime, err := dateparse.ParseLocal(checkOut)
	if err != nil {
		panic(err.Error())
	}

	h := new([]models.Hotel)

	query := db.DB.Table("hotels").Where("id > ?", cursorInt)
	if len(nameStr) != 0 {
		query = query.Where("to_tsvector(name) @@ to_tsquery(?)", nameStr).Or("to_tsvector(description) @@ to_tsquery(?)", nameStr)
	}

	err = query.Limit(limitInt).Find(&h).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later. 😕",
		})
	}

	type maxNumber struct {
		Max int `json:"max"`
	}

	for i, hotel := range *h {
		m := new(maxNumber)
		subQuery := db.DB.Select("check_in_time, SUM(amount) as mySum").Where("hotel_id = ? and check_in_time >= ? and check_out_time <= ?", hotel.UUID, checkInTime, checkOutTime.Add(24*time.Hour)).Table("tickets").Group("check_in_time")
		db.DB.Raw("SELECT MAX(mySum) FROM (?) as groupSum", subQuery).Scan(&m)
		(*h)[i].Room = hotel.Room - m.Max
	}

	return c.Status(fiber.StatusOK).JSON(h)
}
