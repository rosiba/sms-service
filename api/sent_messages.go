package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) ListSentMessages(c *gin.Context) {
	queryCount := c.Query("count")
	count, err := strconv.ParseUint(queryCount, 10, 32)
	if err != nil || queryCount == "" {
		// TODO: set default count as a constant
		count = 100
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
