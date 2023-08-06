package util

import (
	"math/rand"
	"strings"
	"time"
)

func OtpGenerator() string {
	lengthOFOtp := 6
	rand.Seed(time.Now().UnixNano())

	// List of characters for generating the OTP
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	otp := make([]byte, lengthOFOtp)
	for i := range otp {
		otp[i] = characters[rand.Intn(len(characters))]
	}

	return strings.ToUpper(string(otp))
}
