package routers

import (
	"github.com/gin-gonic/gin"

	"ot/api"
	"ot/middleware"
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
	router.Use(cors.Cors())

	allow := middleware.AllowPathPrefixSkipper("/api/v1/login")
	a, _, _ := mdAuth.InitAuth()
	router.Use(mdAuth.UserAuthMiddleware(a, allow))

	router.GET("/health", api.Health)
	v1 := router.Group("api/v1", api.Login)
	v1.POST("/login")
	return router
}
