package delivery

import (
	"log"
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
	s.Start()
	ticker := time.NewTicker(Interval)
	resultChan := make(chan DeliveryResult)
	errorChan := make(chan error)
	for {
		select {
		case <-ticker.C:
			if s.status == DeliveryStatusON {
				messages, err := s.mr.GetPendingMessages(ConcurrentMessageCount)
				if err != nil {
					log.Println("error getting unsent messages", err)
					continue
				}
				for _, message := range messages {
					go s.deliver(message, resultChan, errorChan)
				}
			}
		case err := <-errorChan:
			log.Println(err.Error())
		}
	}
}

func (s *Service) Start() {
	s.status = DeliveryStatusON
	log.Println("starting delivery service")
}

func (s *Service) Stop() {
	s.status = DeliveryStatusOFF
	log.Println("stopping delivery service")
}
