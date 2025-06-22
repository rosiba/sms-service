package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ActionStart string = "start"
	ActionStop  string = "stop"
)

func (h *Handler) StartStopDelivery(c *gin.Context) {
	var requestBody StartStopDeliveryRequest
	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			ErrorMessage: InvalidRequestBody,
			Error:        err,
		})
		return
	}

	switch requestBody.Action {
	case ActionStart:
		h.ds.Start()
		c.JSON(http.StatusOK, Response{
			Message: DeliveryStartSuccessful,
		})
	case ActionStop:
		h.ds.Stop()
		c.JSON(http.StatusOK, Response{
			Message: DeliveryStopSuccessful,
		})
	default:
		c.JSON(http.StatusBadRequest, Response{
			Message: InvalidAction,
		})
	}
}

type StartStopDeliveryRequest struct {
	Action string `json:"action"`
}

type StartStopDeliveryResponse struct {
	Message string `json:"message"`
}
