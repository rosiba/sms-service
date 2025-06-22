package model

import "time"

const (
	MessageStatusPending = "pending"
	MessageStatusSent    = "sent"
)

type Message struct {
	ID        string     `json:"id"`
	Status    string     `json:"status,omitempty"`
	Recipient string     `json:"recipient"`
	Content   string     `json:"content"`
	SentAt    *time.Time `json:"sentAt,omitempty"`

	Model
}
