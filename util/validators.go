package util

import (
	"regexp"

	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/models"

	valid "github.com/asaskevich/govalidator"
)

func IsEmpty(str string) (bool, string) {
	if valid.HasWhitespaceOnly(str) || str == "" {
		return true, "Must not empty"
	}
	return false, ""
}

func ValidateRegister(u *models.User) *models.UserErrors {
	e := &models.UserErrors{}
	e.Err, e.Username = IsEmpty(u.Username)
	e.Err, e.BirthDate = IsEmpty(u.BirthDate)

	if !valid.IsEmail(u.Email) {
		e.Err, e.Email = true, "Must be a valid email"
	}

	re := regexp.MustCompile("\\d") // regex check for at least one integer in string
	if !(len(u.Password) >= 8 && valid.HasLowerCase(u.Password) && valid.HasUpperCase(u.Password) && re.MatchString(u.Password)) {
		e.Err, e.Password = true, "Length of password should be atleast 8 and it must be a combination of uppercase letters, lowercase letters and numbers"
	}

	return e
}