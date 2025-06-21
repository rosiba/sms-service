package api

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	// TODO: add other middlewares if necessary

	r.GET("/ping", Ping)
	r.POST("/messages", SendMessage)

	return r
}
