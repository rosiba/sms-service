package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) ListSentMessages(c *gin.Context) {
	const DefaultLimit = 10
	queryCount := c.Query("count")
	count, err := strconv.ParseUint(queryCount, 10, 32)
	if err != nil || queryCount == "" {
		count = DefaultLimit
	}

	messages, err := h.mr.GetSentMessages(uint(count))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			ErrorMessage: FailedGetUnsentMessages,
			Error:        err,
		})
		return
	}

	c.JSON(http.StatusOK, messages)
}
