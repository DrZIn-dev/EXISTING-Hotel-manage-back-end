package router

import (
	db "github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/database"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/models"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/util"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

func SetupUserRoutes() {
	USER.Post("/signup", CreateUser)
	USER.Post("/signin", LoginUser)
}

func CreateUser(c *fiber.Ctx) error {
	u := new(models.User)

	err := c.BodyParser(u)

	if err != nil {
		return c.JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	errors := util.ValidateRegister(u)
	if errors.Err {
		return c.JSON(errors)
	}

	emailCount := db.DB.Where(&models.User{Email: u.Email}).First(new(models.User)).RowsAffected
	if emailCount > 0 {
		errors.Err, errors.Username = true, "Email is already registered"
	}

	usernameCount := db.DB.Where(&models.User{Username: u.Username}).First(new(models.User)).RowsAffected
	if usernameCount > 0 {
		errors.Err, errors.Username = true, "Username is already registered"
	}
	if errors.Err {
		return c.JSON(errors)
	}

	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, rand.Intn(bcrypt.MaxCost-bcrypt.MinCost)+bcrypt.MinCost)

	if err != nil {
		panic(err)
	}

	u.Password = string(hashedPassword)
	err = db.DB.Create(&u).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later. 😕",
		})
	}

	accessToken, refreshToken := util.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)

	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func LoginUser(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity" valid:"required"`
		Password string `json:"password" valid:"required"`
	}

	input := new(LoginInput)
	err := c.BodyParser(input)
	if err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}
	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	u := new(models.User)
	res := db.DB.Where(&models.User{Email: input.Identity}).Or(&models.User{Username: input.Identity}).First(&u)
	if res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))
	if err != nil {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	accessToken, refreshToken := util.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
