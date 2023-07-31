// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type MutationError interface {
	IsMutationError()
	GetMessage() string
	GetCode() int
}

type QueryError interface {
	IsQueryError()
	GetMessage() string
	GetCode() int
}

type AddressInput struct {
	City     *string `json:"City,omitempty"`
	District *string `json:"District,omitempty"`
	State    *string `json:"State,omitempty"`
}

type AuthMutationResponse struct {
	Data  *User         `json:"data,omitempty"`
	Error MutationError `json:"error,omitempty"`
}

type AuthQueryResponse struct {
	Data  []*User    `json:"data,omitempty"`
	Error QueryError `json:"error,omitempty"`
}

type AuthResponse struct {
	ID    *string       `json:"id,omitempty"`
	Data  *AuthToken    `json:"data,omitempty"`
	Error MutationError `json:"error,omitempty"`
}

type AuthToken struct {
	AccessToken string `json:"accessToken"`
}

type AuthorizationError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (AuthorizationError) IsQueryError()           {}
func (this AuthorizationError) GetMessage() string { return this.Message }
func (this AuthorizationError) GetCode() int       { return this.Code }

func (AuthorizationError) IsMutationError() {}

type BadRequestError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (BadRequestError) IsQueryError()           {}
func (this BadRequestError) GetMessage() string { return this.Message }
func (this BadRequestError) GetCode() int       { return this.Code }

func (BadRequestError) IsMutationError() {}

type CreateProfileInput struct {
	FirstName     string        `json:"firstName"`
	LastName      string        `json:"lastName"`
	ContactNumber string        `json:"contactNumber"`
	DateOfBirth   *time.Time    `json:"DateOfBirth,omitempty"`
	Address       *AddressInput `json:"Address,omitempty"`
}

type ForgetPasswordInput struct {
	Email    string   `json:"email"`
	UserType UserType `json:"userType"`
}

type GetByIDInput struct {
	ID string `json:"ID"`
}

type GetUserInput struct {
	ID string `json:"id"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NotFoundError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (NotFoundError) IsQueryError()           {}
func (this NotFoundError) GetMessage() string { return this.Message }
func (this NotFoundError) GetCode() int       { return this.Code }

func (NotFoundError) IsMutationError() {}

type OtpMutationResponse struct {
	Data  *bool         `json:"data,omitempty"`
	Error MutationError `json:"error,omitempty"`
}

type ProfileMutation struct {
	CreateProfile *ProfileMutationResponse `json:"createProfile"`
	UpdateProfile *ProfileMutationResponse `json:"updateProfile"`
}

type ProfileMutationResponse struct {
	Data  *Profile      `json:"data,omitempty"`
	Error MutationError `json:"error,omitempty"`
}

type ProfileQuery struct {
	GetProfile *ProfileQueryResponse `json:"getProfile"`
}

type ProfileQueryResponse struct {
	Data  *Profile      `json:"data,omitempty"`
	Error MutationError `json:"error,omitempty"`
}

type RegisterResponse struct {
	UserID *string       `json:"userID,omitempty"`
	Error  MutationError `json:"error,omitempty"`
}

type ResendOtpInput struct {
	UserID string `json:"userID"`
}

type ResendOtpMutation struct {
	Resend    *OtpMutationResponse `json:"resend"`
	VerifyOtp *OtpMutationResponse `json:"verifyOtp"`
}

type ResetPasswordInput struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type ServerError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (ServerError) IsQueryError()           {}
func (this ServerError) GetMessage() string { return this.Message }
func (this ServerError) GetCode() int       { return this.Code }

func (ServerError) IsMutationError() {}

type UpdatePasswordInput struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UpdateProfileInput struct {
	UserID        string        `json:"userID"`
	FirstName     *string       `json:"firstName,omitempty"`
	LastName      *string       `json:"lastName,omitempty"`
	ContactNumber *string       `json:"contactNumber,omitempty"`
	DateOfBirth   *time.Time    `json:"DateOfBirth,omitempty"`
	Address       *AddressInput `json:"Address,omitempty"`
}

type UserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type UserMutation struct {
	CreateUser     *AuthMutationResponse `json:"createUser"`
	LoginUser      *AuthResponse         `json:"loginUser"`
	Otp            *ResendOtpMutation    `json:"otp"`
	VerifyUser     *AuthMutationResponse `json:"verifyUser"`
	UpdatePassword *RegisterResponse     `json:"updatePassword"`
	ForgetPassword *RegisterResponse     `json:"forgetPassword"`
	ResetPassword  *RegisterResponse     `json:"resetPassword"`
}

type UserQuery struct {
	GetUserDetails *AuthQueryResponse `json:"getUserDetails"`
}

type ValidationError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (ValidationError) IsMutationError()        {}
func (this ValidationError) GetMessage() string { return this.Message }
func (this ValidationError) GetCode() int       { return this.Code }

type VerifyOtpInput struct {
	Otp   string `json:"otp"`
	Email string `json:"email"`
}

type UserVerificationInput struct {
	Otp    string `json:"otp"`
	UserID string `json:"userID"`
}

type EmailType string

const (
	EmailTypeOtpVerification EmailType = "OTP_VERIFICATION"
	EmailTypeStaffCreation   EmailType = "STAFF_CREATION"
)

var AllEmailType = []EmailType{
	EmailTypeOtpVerification,
	EmailTypeStaffCreation,
}

func (e EmailType) IsValid() bool {
	switch e {
	case EmailTypeOtpVerification, EmailTypeStaffCreation:
		return true
	}
	return false
}

func (e EmailType) String() string {
	return string(e)
}

func (e *EmailType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EmailType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EmailType", str)
	}
	return nil
}

func (e EmailType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOthers Gender = "OTHERS"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderOthers,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderOthers:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusInactive UserStatus = "INACTIVE"
	UserStatusActive   UserStatus = "ACTIVE"
)

var AllUserStatus = []UserStatus{
	UserStatusInactive,
	UserStatusActive,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusInactive, UserStatusActive:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserType string

const (
	UserTypeLogicloud UserType = "LOGICLOUD"
	UserTypeStaff     UserType = "STAFF"
	UserTypeAdmin     UserType = "ADMIN"
)

var AllUserType = []UserType{
	UserTypeLogicloud,
	UserTypeStaff,
	UserTypeAdmin,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeLogicloud, UserTypeStaff, UserTypeAdmin:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
