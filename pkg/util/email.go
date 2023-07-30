package util

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendVerificationEmail(otp string, userEmail string) {

	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	// toList is list of email address that email is to be sent.
	toList := []string{userEmail}
	host := "smtp.gmail.com"
	// Its the default port of smtp server
	port := "587"

	// This is the message to send in the mail
	// msg := GetEmailMessage()
	msg := fmt.Sprintf("Your Otp is ", otp)

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte(msg)

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		log.Println("Error While Sending Email", err.Error())
	}
	fmt.Println("Successfully sent mail to all user", userEmail)
}

func GetEmailMessage() string {
	return "Hello"
}
