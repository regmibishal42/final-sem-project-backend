package model

type User struct {
	Base
	Email      string     `json:"email" gorm:"uniqueIndex"`
	UserType   UserType   `json:"userType"`
	IsVerified bool       `json:"isVerified" gorm:"default:false"`
	Status     UserStatus `json:"status"`
	Password   string     `json:"password"`
}
