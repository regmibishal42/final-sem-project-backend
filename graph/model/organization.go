package model

type Organization struct {
	Base
	Email              *string             `json:"email,omitempty"`
	Contact            *string             `json:"contact,omitempty"`
	Address            *string             `json:"Address,omitempty"`
	CreatedByID        string              `json:"createdByID"`
	CreatedBy          *User               `json:"createdBy,omitempty" gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VerificationStatus *VerificationStatus `json:"verificationStatus,omitempty"`
}
