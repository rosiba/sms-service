package delivery

import (
	"fmt"
	"log"
	"sms-service/internal/model"
	"sms-service/internal/repository"
	"time"
)

const (
	DeliveryStatusON       = "ON"
	DeliveryStatusOFF      = "OFF"
	Interval               = 2 * time.Second
	ConcurrentMessageCount = 2
)

type Service struct {
	mr     repository.MessageRepository
	status string
}

type DeliveryResult struct {
	MessageID  string
	ExternalID string
}

func NewDeliveryService(mr repository.MessageRepository) *Service {
	return &Service{
		mr: mr,
	}
}

func (s *Service) Run() {
	log.Println("delivery service started")
	ticker := time.NewTicker(Interval)
	resultChan := make(chan DeliveryResult)
	errorChan := make(chan error)
	for {
		select {
		case <-ticker.C:
			if s.status == DeliveryStatusON {
				messages, err := s.mr.GetUnsentMessages(ConcurrentMessageCount)
				if err != nil {
					log.Println("error getting unsent messages", err)
					continue
				}
				for _, message := range messages {
					go deliver(message, resultChan, errorChan)
				}
			}
		case r := <-resultChan:
			log.Println(r.ExternalID)
			if err := s.mr.SetMessageStatus(r.MessageID, model.MessageStatusSent); err != nil {
				errorChan <- fmt.Errorf("error setting message status: %v", err)
				continue
			}
			// TODO: cache to redis
		case err := <-errorChan:
			log.Println(err.Error())
		}
	}
}

func (s *Service) Start() {
	s.status = DeliveryStatusON
}

func (s *Service) Stop() {
	s.status = DeliveryStatusOFF
}
