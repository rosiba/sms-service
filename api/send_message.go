package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sms-service/internal/model"
)

func (h *Handler) SendMessage(c *gin.Context) {
	var requestBody SendMessageRequest
	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageID, err := h.mr.SaveMessage(model.Message{
		Recipient: requestBody.Recipient,
		Content:   requestBody.Content,
	})
	if err != nil {
		log.Println("failed to save message:", err)
		c.JSON(http.StatusInternalServerError, Response{
			ErrorMessage: FailedSaveMessage,
			Error:        err,
		})
	}

	c.JSON(http.StatusCreated, SendMessageResponse{
		Message:   MessageCreated,
		MessageID: messageID,
	})
}

type SendMessageRequest struct {
	Recipient string `json:"recipient" binding:"required"`
	Content   string `json:"content" binding:"required,min=2,max=180"`
}

type SendMessageResponse struct {
	Message   string `json:"message"`
	MessageID string `json:"message_id"`
}
