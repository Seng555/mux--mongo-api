// model/reset_token.go
package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ResetToken represents a reset token model
type ResetToken struct {
	UserID     primitive.ObjectID `json:"user_id"`
	Email      string             `json:"email,omitempty"`
	Token      string             `json:"token,omitempty"`
	OnModel    string             `json:"on_model,omitempty"`
	Expiration time.Time          `json:"expiration,omitempty"`
	CreateDate time.Time          `json:"create_date,omitempty"`
	UpdateDate time.Time          `json:"update_date,omitempty"`
}
