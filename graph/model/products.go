package model

import (
	"backend/pkg/util"
	"time"
)

type Product struct {
	Base
	Name           string        `json:"name"`
	BoughtOn       time.Time     `json:"boughtOn"`
	Units          int           `json:"units"`
	CategoryID     string        `json:"categoryID"`
	Category       *Category     `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizationID string        `json:"organizationID"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (input *CreateProductInput) Validator() (*Product, *ValidationError) {
	product := &Product{}
	if len(input.Name) < 3 {
		return nil, &ValidationError{
			Message: "product name cannot be less then 3 char long",
			Code:    401,
		}
	}
	if input.BoughtOn.Unix() > time.Now().Unix() {
		return nil, &ValidationError{
			Message: "Bought Date cannot be greater then todays date",
			Code:    401,
		}
	}
	if input.Units < 1 || input.Units > 500 {
		return nil, &ValidationError{
			Message: "Units Should be between 1-500",
			Code:    401,
		}
	}
	if !util.IsValidID(input.CategoryID) {
		return nil, &ValidationError{
			Message: "invalid CategoryID provided",
			Code:    401,
		}
	}
	product.Name = input.Name
	product.BoughtOn = input.BoughtOn
	product.Units = input.Units
	product.CategoryID = input.CategoryID

	return product, nil
}
