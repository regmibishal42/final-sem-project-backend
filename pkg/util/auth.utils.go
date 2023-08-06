package util

import (
	"errors"
	"log"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

//hash password
func HashPassword(password string) (string, error) {
	costForHashing := 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costForHashing)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Check if email is valid
func IsEmailValid(email string) error {
	emailRegex := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email address")
	}
	return nil
}

//Check if password is strong
func PasswordValidator(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

func GetDefaultStaffPassword() string {
	defaultPassword := "Staff@123"
	hash, err := HashPassword(defaultPassword)
	if err != nil {
		log.Println("Error Hashing Staff Password", err.Error())
	}
	return hash
}
