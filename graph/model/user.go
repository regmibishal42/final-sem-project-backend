package model

import "time"

type User struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	UserType   UserType   `json:"userType"`
	IsVerified bool       `json:"isVerified"`
	Status     UserStatus `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
}
