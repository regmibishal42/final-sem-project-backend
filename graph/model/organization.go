package model

type Organization struct {
	Base
	Email              string              `json:"email,omitempty"`
	Contact            string              `json:"contact,omitempty"`
	Address            string              `json:"Address,omitempty"`
	CreatedByID        string              `json:"createdByID"`
	CreatedBy          *User               `json:"createdBy,omitempty" gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VerificationStatus *VerificationStatus `json:"verificationStatus,omitempty" gorm:"default:NOT_VERIFIED"`
}

func (input *CreateOrganizationInput) Validator() (Organization, *ValidationError) {
	organization := Organization{
		Email:   input.Email,
		Contact: input.Contact,
	}
	if input.Address != nil {
		//city := util.Ref()
		organization.Address = input.Address.City + "," + input.Address.District + "," + input.Address.State
	}
	return organization, nil

}
