package util

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func sendVerificationEmail(userEmail string, message string) {

	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	// toList is list of email address that email is to be sent.
	toList := []string{userEmail}
	host := "smtp.gmail.com"
	// Its the default port of smtp server
	port := "587"

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte(message)

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

func SendStaffAccountCreationEmail(staffEmail string, staffName string) {
	defaultPassword := os.Getenv("DEFAULT_STAFF_PASSWORD")
	emailTemplate := fmt.Sprintf(`Subject: New Staff Member Account Created

Dear %s,

We hope this message finds you well. We would like to inform you that a new staff member account has been successfully created for your organization. Here are the details:

Staff Member's Name: %s
Staff Member's Email: %s
Default Password: %s

Please ensure that you share these credentials securely with the staff member. We recommend advising them to change their password upon their first login for security purposes.

If you have any questions or concerns regarding this new account setup, please feel free to reach out to our support team or admin. We're here to assist you.

Thank you for using our services to manage your organization's staff accounts. We appreciate your trust in us.

Best regards,
Your Name
Your Title
Your Contact Information`, staffName, staffName, staffEmail, defaultPassword)
	sendVerificationEmail(staffEmail, emailTemplate)

}

//send Email when user creates an account
func SendOtpEmail(userName string, userEmail string, otp string) {
	emailTemplate := fmt.Sprintf(`Subject: Account Registration - OTP Verification

Dear %s,

Thank you for creating an account with our app! To complete your registration, please use the following OTP code:

OTP Code: %s

This OTP code is valid for a single use and will expire after a certain period of time. Please enter it on the app's verification screen.

If you did not create this account or have any questions, please contact our support team.

Welcome to LogiCloud!

Best regards,
Your App Team`, userName, otp)
	sendVerificationEmail(userEmail, emailTemplate)

}

func SendResetPasswordOtpEmail(userName string, userEmail string, otp string) {
	emailTemplate := fmt.Sprintf(`Subject: Reset Account Password - OTP Verification

Dear %s,

You have requested to reset your account's password as you might have forgotten the old one.! To complete the process, please use the following OTP code:

OTP Code: %s

This OTP code is valid for a single use and will expire after a certain period of time. Please enter it on the app's verification screen.

If you did not request for this otp. You can ignore this email.

Welcome to LogiCloud!

Best regards,
LogiCloud Team`, userName, otp)
	sendVerificationEmail(userEmail, emailTemplate)
}
