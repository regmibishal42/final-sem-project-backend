package model

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

// func (input *CreateSaleInput) Validator() (*Sales, *ValidationError) {
// 	sales := Sales{}
// 	if !util.IsValidID(input.ProductID) {
// 		return nil, &ValidationError{
// 			Message: "invalid productID",
// 			Code:    401,
// 		}
// 	}
// 	sales.ProductID = input.ProductID
// 	sales.SoldAt = input.SoldAt
// 	input.Units = in
// 	return nil, nil
// }
