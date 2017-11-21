package pjrouter

import (
	"github.com/gin-gonic/gin"
	"logics/register"
	"logics/session"
)

var pjRouter = gin.Default()

func Load() *gin.Engine {
	apiV1 := pjRouter.Group("/v1")
	apiV1.POST("/register/:way", register.Register)
	apiV1.POST("/session/:way", session.Create)
	return pjRouter
}
