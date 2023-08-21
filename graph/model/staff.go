package model

import (
	"backend/pkg/util"
	"time"
)

type Staff struct {
	StaffID        string        `json:"staffID"`
	Staff          *User         `json:"staff" gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizationID string        `json:"organizationID"`
	Organization   *Organization `json:"Organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	JoinedOn       time.Time     `json:"joinedOn"`
	Post           string        `json:"post"`
	Salary         *float64      `json:"salary,omitempty"`
	IsAuthorized   *bool         `json:"isAuthorized,omitempty" gorm:"default:FALSE"`
	IsActive       *bool         `json:"isActive,omitempty" gorm:"default:TRUE"`
}

func (input *CreateStaffInput) Validator() *ValidationError {
	if len(input.ContactNumber) < 10 {
		return &ValidationError{
			Message: "invalid ContactNumber",
		}
	}
	if err := util.IsEmailValid(input.Email); err != nil {
		return &ValidationError{
			Message: "invalid email",
		}
	}
	if input.JoinedOn.Unix() > time.Now().Unix() {
		return &ValidationError{
			Message: "Join Date Cannot be greater then current date",
		}
	}
	return nil

}

func (input *UpdateStaffInput) Validator() (*Staff, *ValidationError) {
	staff := &Staff{}
	if !util.IsValidID(input.StaffID) {
		return nil, &ValidationError{
			Message: "invalid StaffID",
			Code:    401,
		}
	}
	staff.StaffID = input.StaffID
	if input.Post != nil {
		staff.Post = *input.Post
	}
	if input.IsAuthorized != nil {
		staff.IsAuthorized = input.IsAuthorized
	}
	if input.IsActive != nil {
		staff.IsActive = input.IsActive
	}
	if input.Salary != nil {
		staff.Salary = input.Salary
	}
	return staff, nil
}
