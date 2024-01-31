// model/address.go
package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address represents an address model
type Address struct {
	UserID        primitive.ObjectID `json:"user_id"`
	DocumentPic   string             `json:"document_pic,omitempty"`
	Address       string             `json:"address"`
	VerifyStatus  int                `json:"verifyStatus,omitempty"`
	City          string             `json:"city"`
	StateProvince string             `json:"stateProvince"`
	ZipCode       string             `json:"zipCode"`
	Country       map[string]string  `json:"country,omitempty"` // Change type to map[string]string
	CreateDate    time.Time          `json:"create_date"`
	UpdateDate    time.Time          `json:"update_date"`
	IsActive      bool               `json:"isActive"`
}
