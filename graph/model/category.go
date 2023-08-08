package model

type Category struct {
	Base
	Name           string        `json:"name" gorm:"uniqueIndex:Idx;"`
	OrganizationID string        `json:"organizationID" gorm:"uniqueIndex:Idx;"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (input *CreateCategoryInput) Validator() *ValidationError {
	if len(input.Name) < 3 {
		return &ValidationError{
			Message: "Category name should be at-least 3 char long",
			Code:    401,
		}
	}
	return nil
}
