package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	// TODO: set mode to final value before submission
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("failed to set trusted proxies: %v", err)
	}

	// TODO: add other middlewares if necessary

	r.GET("/ping", Ping)
	r.POST("/messages", SendMessage)

	return r, nil
}
