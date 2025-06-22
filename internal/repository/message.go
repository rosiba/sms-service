package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
	"sms-service/internal/model"
	"time"
)

type MessageRepository interface {
	GetPendingMessages(count uint) ([]model.Message, error)
	GetSentMessages(count uint) ([]model.Message, error)
	SaveMessage(message model.Message) (string, error)
	SetMessageAsSent(id pgtype.UUID, sentAt time.Time) error
}
