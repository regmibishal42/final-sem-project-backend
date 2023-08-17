package model

import "backend/pkg/util"

type Sales struct {
	Base
	ProductID      string        `json:"productID"`
	Product        *Product      `json:"product" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizationID string        `json:"organizationID"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UnitsSold      int           `json:"unitsSold"`
	SoldAt         float64       `json:"soldAt"`
	SoldByID       string        `json:"soldByID"`
	SoldBy         *User         `json:"soldBy" gorm:"foreignKey:SoldByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (input *CreateSaleInput) Validator() (*Sales, *ValidationError) {
	sales := Sales{}
	if !util.IsValidID(input.ProductID) {
		return nil, &ValidationError{
			Message: "invalid productID",
			Code:    401,
		}
	}
	if input.SoldAt < 1 {
		return nil, &ValidationError{
			Message: "invalid SoldAt",
			Code:    401,
		}
	}
	if input.Units < 1 {
		return nil, &ValidationError{
			Message: "invalid Units",
			Code:    401,
		}
	}
	sales.ProductID = input.ProductID
	sales.SoldAt = input.SoldAt
	sales.UnitsSold = input.Units
	return &sales, nil
}

func (input *UpdateSalesInput) Validator() (*Sales, *ValidationError) {
	sales := Sales{}
	if !util.IsValidID(input.SalesID) {
		return nil, &ValidationError{
			Message: "invalid SalesID",
			Code:    401,
		}
	}
	sales.ID = input.SalesID
	if input.SoldAt != nil {
		if *input.SoldAt < 1 {
			return nil, &ValidationError{
				Message: "invalid SoldAt",
				Code:    401,
			}

		}
		sales.SoldAt = *input.SoldAt
	}
	return &sales, nil
}

func (input *FilterSalesInput) Validator() *ValidationError {
	if input.CategoryID != nil && !util.IsValidID(*input.CategoryID) {
		return &ValidationError{
			Message: "invalid CategoryID",
			Code:    401,
		}
	}
	if input.ProductID != nil && !util.IsValidID(*input.ProductID) {
		return &ValidationError{
			Message: "invalid CategoryID",
			Code:    401,
		}
	}
	return nil

}
