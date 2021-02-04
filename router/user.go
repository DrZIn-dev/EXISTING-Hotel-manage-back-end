package router

import (
	db "github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/database"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/models"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/util"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("PRIVATE_KEY"))

func SetupUserRoutes() {
	USER.Post("/signup", CreateUser)
	USER.Post("/signin", LoginUser)
	USER.Get("/get-access-token", GetAccessToken)

	private := USER.Group("/private")
	private.Use(util.SecureAuth())
	private.Get("/user", GetUserData)
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
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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
	accessCookie.HTTPOnly = false
	accessCookie.SameSite = ""
	accessCookie.Secure = false
	refreshCookie.HTTPOnly = false
	refreshCookie.SameSite = ""
	refreshCookie.Secure = false

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
	accessCookie.HTTPOnly = false
	accessCookie.SameSite = ""
	accessCookie.Secure = false
	refreshCookie.HTTPOnly = false
	refreshCookie.SameSite = ""
	refreshCookie.Secure = false
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func GetAccessToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	refreshClaims := new(models.Claims)
	token, _ := jwt.ParseWithClaims(refreshToken, refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if res := db.DB.Where(
		"expires_at = ? AND issued_at = ? AND issuer = ?",
		refreshClaims.ExpiresAt, refreshClaims.IssuedAt, refreshClaims.Issuer,
	).First(&models.Claims{}); res.RowsAffected <= 0 {
		// no such refresh token exist in the database
		c.ClearCookie("access_token", "refresh_token")
		return c.SendStatus(fiber.StatusForbidden)
	}

	if token.Valid {
		if refreshClaims.ExpiresAt < time.Now().Unix() {
			// refresh token is expired
			c.ClearCookie("access_token", "refresh_token")
			return c.SendStatus(fiber.StatusForbidden)
		}
	} else {
		// malformed refresh token
		c.ClearCookie("access_token", "refresh_token")
		return c.SendStatus(fiber.StatusForbidden)
	}

	_, accessToken := util.GenerateAccessClaims(refreshClaims.Issuer)

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: false,
		SameSite: "",
		Secure:   false,
	})

	return c.JSON(fiber.Map{"access_token": accessToken})

}

func GetUserData(c *fiber.Ctx) error {
	id := c.Locals("id")
	u := new(models.UserResponse)
	res := db.DB.Model(&models.User{}).Where("uuid = ? ", id).First(&u)
	if res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Cannot find the User"})
	}

	return c.JSON(u)
}
