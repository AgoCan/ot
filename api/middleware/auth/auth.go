package auth

import (
	"fmt"
	"net/http"
	"ot/config"
	"ot/middleware"
	"ot/pkg/auth"
	"ot/pkg/auth/jwtauth"
	"ot/pkg/auth/jwtauth/store/buntdb"
	"ot/pkg/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userIDCtx struct{}

func UserAuthMiddleware(a auth.Auther, skippers ...middleware.SkipperFunc) gin.HandlerFunc {
	if !config.Conf.Jwt.Enable {
		return func(c *gin.Context) {
			c.Next()
		}
	}
	return func(c *gin.Context) {
		if middleware.SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		var token string
		auth := c.GetHeader("Authorization")
		prefix := "Bearer "
		if auth != "" && strings.HasPrefix(auth, prefix) {
			token = auth[len(prefix):]
		}
		if token == "" {
			c.JSON(http.StatusOK, response.ErrInvalidToken)
			c.Abort()
			return
		}
		userID, err := a.ParseUserID(c.Request.Context(), token)
		fmt.Println(userID, err)

		c.Next()
	}
}

func InitAuth() (auth.Auther, func(), error) {
	cfg := config.Conf.Jwt

	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(cfg.Expired))
	opts = append(opts, jwtauth.SetSigningKey([]byte(cfg.SigningKey)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(cfg.SigningKey), nil
	}))

	var method jwt.SigningMethod
	switch cfg.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	opts = append(opts, jwtauth.SetSigningMethod(method))

	var store jwtauth.Storer
	switch cfg.Store {
	case "redis":
		fmt.Println(123)
	default:
		s, err := buntdb.NewStore(cfg.FilePath)
		if err != nil {
			return nil, nil, err
		}
		store = s
	}

	auth := jwtauth.New(store, opts...)
	cleanFunc := func() {
		auth.Release()
	}
	return auth, cleanFunc, nil
}
