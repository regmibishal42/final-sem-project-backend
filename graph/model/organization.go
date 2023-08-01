package model

import "backend/pkg/util"

type Organization struct {
	Base
	Name               string              `json:"name"`
	Email              string              `json:"email,omitempty"`
	Contact            string              `json:"contact,omitempty"`
	Address            Address             `json:"Address,omitempty" gorm:"serializer:json"`
	PanNumber          *string             `json:"panNumber"`
	CreatedByID        string              `json:"createdByID"`
	CreatedBy          *User               `json:"createdBy,omitempty" gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VerificationStatus *VerificationStatus `json:"verificationStatus,omitempty" gorm:"default:NOT_VERIFIED"`
}

func (input *CreateOrganizationInput) Validator() (Organization, *ValidationError) {
	organization := Organization{}
	//validate name
	if len(input.Name) < 5 {
		return organization, &ValidationError{
			Message: "organization name should be al-least 5 char long",
			Code:    401,
		}
	}
	organization.Name = input.Name
	//validate email
	if err := util.IsEmailValid(input.Email); err != nil {
		return organization, &ValidationError{
			Message: "invalid email provided",
			Code:    401,
		}
	}
	organization.Email = input.Email
	//validate contact
	if len(input.Contact) < 10 {
		return organization, &ValidationError{
			Message: "invalid contact provided",
			Code:    401,
		}
	}
	organization.Contact = input.Contact
	//TODO: Validate Pan
	if input.PanNumber != nil {
		organization.PanNumber = input.PanNumber
	}
	//validate address
	if input.Address != nil {
		//city := util.Ref()
		organization.Address = Address{
			City:     input.Address.City,
			District: input.Address.District,
			State:    input.Address.State,
		}
	}
	return organization, nil

}
