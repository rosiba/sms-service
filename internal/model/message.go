package model

import "time"

const (
	MessageStatusPending = "pending"
	MessageStatusSent    = "sent"
)

type Message struct {
	ID        string
	Status    string
	Recipient string
	Content   string
	SentAt    *time.Time

	Model
}
