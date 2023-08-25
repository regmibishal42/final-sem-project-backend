package model

import (
	"time"
)

type Profile struct {
	UserID        string     `json:"userID" gorm:"primaryKey"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	ContactNumber string     `json:"contactNumber"`
	DateOfBirth   *time.Time `json:"DateOfBirth,omitempty"`
	Address       *Address   `json:"Address,omitempty" gorm:"serializer:json"`
}

func (input *UpdateProfileInput) Validator() (*Profile, *ValidationError) {
	profile := Profile{}
	if input.DateOfBirth != nil {
		if input.DateOfBirth.Unix() >= time.Now().Unix() {
			return nil, &ValidationError{
				Message: "invalid date of birth",
			}
		}
		profile.DateOfBirth = input.DateOfBirth
	}
	if input.Address != nil {
		profile.Address = &Address{
			City:     input.Address.City,
			District: input.Address.District,
			State:    input.Address.District,
		}
	}
	if input.FirstName != nil {
		profile.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		profile.LastName = *input.LastName
	}
	if input.ContactNumber != nil {
		profile.ContactNumber = *input.ContactNumber
	}
	return &profile, nil
}
