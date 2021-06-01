package routers

import (
	"github.com/gin-gonic/gin"

	"ot/api"
	mdAuth "ot/middleware/auth"
	"ot/middleware/cors"
	"ot/middleware/log"
)

// SetupRouter 初始化gin入口，路由信息
func SetupRouter() *gin.Engine {
	router := gin.New()
	if err := log.InitLogger(); err != nil {
		panic(err)
	}
	router.Use(log.GinLogger(log.Logger),
		log.GinRecovery(log.Logger, true))
	a, _, _ := mdAuth.InitAuth()

	router.Use(mdAuth.UserAuthMiddleware(a))
	router.Use(cors.Cors())
	router.GET("/health", api.Health)
	return router
}
