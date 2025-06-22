package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *Handler) (*gin.Engine, error) {
	// TODO: set mode to final value before submission
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("failed to set trusted proxies: %v", err)
	}

	// TODO: add other middlewares if necessary

	r.GET("/ping", Ping)
	r.POST("/messages", h.SendMessage)
	r.GET("/messages/sent", h.ListSentMessages)
	r.POST("/service", h.StartStopDelivery)

	return r, nil
}
