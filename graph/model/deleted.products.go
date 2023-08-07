package model

type DeletedProducts struct {
	Base
	Product *Product `json:"product,omitempty" gorm:"serializer:json"`
}
