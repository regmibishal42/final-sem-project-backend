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
	CostPrice      float64       `json:"costPrice"`
	SellingPrice   float64       `json:"sellingPrice"`
	Category       *Category     `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizationID string        `json:"organizationID"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

//Create Products Input Validation
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

//Update Inputs Validation
func (input UpdateProductInput) Validator() (*Product, *ValidationError) {
	product := Product{}
	if !util.IsValidID(input.ProductID) {
		return nil, &ValidationError{
			Message: "invalid ProductID",
			Code:    401,
		}
	}
	product.ID = input.ProductID
	//validate CategoryID
	if input.CategoryID != nil {
		if !util.IsValidID(*input.CategoryID) {
			return nil, &ValidationError{
				Message: "invalid ProductID",
				Code:    401,
			}
		}
		product.CategoryID = *input.CategoryID
	}
	//validate name
	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, &ValidationError{
				Message: "invalid ProductID",
				Code:    401,
			}
		}
		product.Name = *input.Name
	}
	//validate units
	if input.Units != nil {
		if *input.Units < 1 || *input.Units > 500 {
			return nil, &ValidationError{
				Message: "Units Should be between 1-500",
				Code:    401,
			}
		}
		product.Units = *input.Units
	}
	//validate boughtDate
	if input.BoughtOn != nil {
		if input.BoughtOn.Unix() > time.Now().Unix() {
			return nil, &ValidationError{
				Message: "Bought Date cannot be greater then todays date",
				Code:    401,
			}
		}
		product.BoughtOn = *input.BoughtOn
	}
	return &product, nil
}
