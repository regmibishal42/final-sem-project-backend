package model

import "time"

type Profile struct {
	UserID        string     `json:"userID" gorm:"primaryKey"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	ContactNumber string     `json:"contactNumber"`
	DateOfBirth   *time.Time `json:"DateOfBirth,omitempty"`
	Address       *string    `json:"Address,omitempty"`
}
