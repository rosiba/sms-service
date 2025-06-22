package api

import (
	"sms-service/internal/delivery"
	"sms-service/internal/repository"
)

type Handler struct {
	mr repository.MessageRepository
	ds *delivery.Service
}

func NewHandler(mr repository.MessageRepository, ds *delivery.Service) *Handler {
	return &Handler{
		mr: mr,
		ds: ds,
	}
}
