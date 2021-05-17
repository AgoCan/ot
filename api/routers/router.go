package routers

import (
	"github.com/gin-gonic/gin"

	"ot/api"
	"ot/middleware/log"
)

// SetupRouter 初始化gin入口，路由信息
func SetupRouter() *gin.Engine{
	router := gin.New()
	if err := log.InitLogger(); err != nil {
		panic(err)
	}
	router.Use(log.GinLogger(log.Logger),
		log.GinRecovery(log.Logger, true))

	router.GET("/health", api.Health)
	return router
}