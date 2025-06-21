package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessage(c *gin.Context) {
	var requestBody SendMessageRequest
	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Push to DB

	c.JSON(http.StatusOK, SendMessageResponse{
		Message: "Accepted",
		// TODO: Get UUID for message
		MessageID: "123",
	})
}

type SendMessageRequest struct {
	To      string `json:"to" binding:"required"`
	Content string `json:"content" binding:"required,min=2,max=180"`
}

type SendMessageResponse struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}
