package api

import "sms-service/internal/repository"

type Handler struct {
	mr repository.MessageRepository
}

func NewHandler(mr repository.MessageRepository) *Handler {
	return &Handler{mr: mr}
}
