package repository

import "time"

type CacheRepository interface {
	SaveMessageTime(id string, sentAt time.Time) error
}
