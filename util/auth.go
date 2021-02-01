package util

import (
	db "github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/database"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("PRIVATE_KEY"))

func GenerateTokens(uuid string) (string, string) {
	claims, accessToken := GenerateAccessClaims(uuid)
	refreshToken := GenerateRefreshClaims(claims)

	return accessToken, refreshToken
}

func GenerateAccessClaims(uuid string) (*models.Claims, string) {
	t := time.Now()
	accessClaim := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(15 * time.Minute).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return accessClaim, accessTokenString
}

func GenerateRefreshClaims(cl *models.Claims) string {
	result := db.DB.Where(&models.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer: cl.Issuer,
		},
	}).Find(&models.Claims{})

	if result.RowsAffected > 3 {
		db.DB.Where(&models.Claims{
			StandardClaims: jwt.StandardClaims{
				Issuer: cl.Issuer,
			},
		}).Delete(&models.Claims{})
	}

	t := time.Now()

	refreshClaim := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    cl.Issuer,
			ExpiresAt: t.Add(30 * 24 * time.Hour).Unix(),
			Subject:   "refresh_token",
			IssuedAt:  t.Unix(),
		},
	}

	db.DB.Create(&refreshClaim)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return refreshTokenString
}

func GetAuthCookies(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {
	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   true,
	}

	refreshCookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	}

	return accessCookie, refreshCookie
}
