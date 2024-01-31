// model/user_Model.go
package model

import (
	"time"
)

// User represents a user model
type User struct {
	Email        string     `json:"email" validate:"required,email"`
	Password     string     `json:"password" validate:"required"`
	FirstName    string     `json:"firstName,omitempty"`
	LastName     string     `json:"lastName,omitempty"`
	IdCard       string     `json:"idCard,omitempty"`
	IdCardType   string     `json:"idCardType,omitempty"`
	IdCardPic    string     `json:"idCardPic,omitempty"`
	SelfiePic    string     `json:"selfiePic,omitempty"`
	VerifyStatus int        `json:"verifyStatus,omitempty"`
	DateOfBirth  *time.Time `json:"dateOfBirth,omitempty"`
	UserType     int        `json:"userType,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	Profile      string     `json:"profile,omitempty"`
	CreateDate   time.Time  `json:"createDate"`
	UpdateDate   time.Time  `json:"updateDate"`
	IsActive     bool       `json:"isActive"`
}

// NewUser creates a new User instance with default values
func NewUser(email, password string) *User {
	return &User{
		Email:        email,
		Password:     password,
		CreateDate:   time.Now(),
		UpdateDate:   time.Now(),
		VerifyStatus: VerifyStatusNotVerified,
		IsActive:     true,
	}
}

// Constants for VerifyStatus
const (
	VerifyStatusNotVerified = 0
	VerifyStatusIdentity    = 1
	VerifyStatusFull        = 2
)

// Constants for UserType
const (
	UserTypeStandard = 0
	UserTypeVIP      = 1
)
