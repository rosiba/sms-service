package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

const (
	MessageStatusPending = "pending"
	MessageStatusSent    = "sent"
)

type Message struct {
	ID        pgtype.UUID `json:"id"`
	Status    string      `json:"status,omitempty"`
	Recipient string      `json:"recipient"`
	Content   string      `json:"content"`
	SentAt    *time.Time  `json:"sentAt,omitempty"`

	Model
}
