package model

type Category struct {
	Base
	Name           string        `json:"name" gorm"uniqueIndex:Idx"`
	OrganizationID string        `json:"organizationID" gorm"uniqueIndex:Idx"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
