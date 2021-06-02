package api

import (
	"net/http"
	"ot/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	service := service.User{}
	
	res := service.Login()
	c.JSON(http.StatusOK, res)
}
