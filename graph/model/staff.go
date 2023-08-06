package model

import "time"

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
