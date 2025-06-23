package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *Handler) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("failed to set trusted proxies: %v", err)
	}

	r.Use(cors.Default())

	r.GET("/ping", Ping)
	r.POST("/messages", h.SendMessage)
	r.GET("/messages/sent", h.ListSentMessages)
	r.PUT("/service/status", h.StartStopDelivery)

	return r, nil
}
