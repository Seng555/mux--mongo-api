// model/logs.go
package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Log represents a log document
type Log struct {
	UserID      primitive.ObjectID `json:"user_id"`
	Action      string             `json:"action"`
	OnModel     string             `json:"on_model,omitempty"`
	Status      int                `json:"status"`
	Description string             `json:"des"`
	CreateDate  time.Time          `json:"create_date,omitempty"`
}
