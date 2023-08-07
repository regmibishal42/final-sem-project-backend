package model

import "time"

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
