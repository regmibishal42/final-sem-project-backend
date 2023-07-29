package model

import "backend/pkg/util"

type User struct {
	Base
	Email      string     `json:"email" gorm:"uniqueIndex"`
	UserType   UserType   `json:"userType"`
	IsVerified bool       `json:"isVerified" gorm:"default:false"`
	Status     UserStatus `json:"status"`
	Password   string     `json:"password"`
	Profile    *Profile   `json:"profile" gorm:"foreignKey:UserID"`
}

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
