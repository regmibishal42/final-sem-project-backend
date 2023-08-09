package model

import (
	"backend/pkg/util"
)

type User struct {
	Base
	Email      string     `json:"email" gorm:"uniqueIndex"`
	UserType   UserType   `json:"userType"`
	IsVerified bool       `json:"isVerified" gorm:"default:false"`
	Status     UserStatus `json:"status"`
	Password   string     `json:"password"`
	Profile    *Profile   `json:"profile" gorm:"foreignKey:UserID"`
}

//validate create USER Input
func (input *UserInput) Validator() (*User, *ValidationError) {
	user := User{}
	if err := util.IsEmailValid(input.Email); err != nil {
		return nil, &ValidationError{Message: "invalid email", Code: 401}
	}
	user.Email = input.Email

	if !util.PasswordValidator(input.Password) {
		return nil, &ValidationError{Message: "enter a strong password", Code: 401}
	}
	user.UserType = UserTypeAdmin
	return &user, nil
}

//validate login input
func (input *LoginInput) Validator() *ValidationError {
	if err := util.IsEmailValid(input.Email); err != nil {
		return &ValidationError{Message: "invalid email", Code: 401}
	}

	if !util.PasswordValidator(input.Password) {
		return &ValidationError{Message: "enter a strong password", Code: 401}
	}
	return nil
}

func (input *UpdatePasswordInput) Validator() *ValidationError {
	if !util.PasswordValidator(input.OldPassword) {
		return &ValidationError{
			Message: "invalid old password",
			Code:    401,
		}
	}
	if !util.PasswordValidator(input.NewPassword) {
		return &ValidationError{Message: "enter a strong new password", Code: 401}
	}
	return nil
}
