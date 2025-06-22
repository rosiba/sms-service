package repository

import "sms-service/internal/model"

type MessageRepository interface {
	GetUnsentMessages(count uint) ([]model.Message, error)
	GetSentMessages() ([]model.Message, error)
	SetMessageStatus(messageID string, status string) error
	SaveMessage(message model.Message) error
}
