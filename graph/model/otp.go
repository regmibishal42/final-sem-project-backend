package model

import "time"

type Otp struct {
	UserId    string    `json:"userId" gorm:"primary_key"`
	User      *User     `json:"user,omitempty"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"createdAt"`
}
