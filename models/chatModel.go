package models

import "time"

// Chat represents a one-on-one chat message
type Chat struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	From      string    `json:"from" bson:"from"`
	To        string    `json:"to" bson:"to"`
	Msg       string    `json:"message" bson:"message"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
