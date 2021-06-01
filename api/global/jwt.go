package global

import (
	"ot/config"
	"ot/pkg/auth"
	"ot/pkg/auth/jwtauth"
	"ot/pkg/auth/jwtauth/store/buntdb"

	"github.com/dgrijalva/jwt-go"
)

func InitJwt() (auth.Auther, func(), error) {
	var opts []jwtauth.Option
	jwtConfig := config.Conf.Jwt
	opts = append(opts, jwtauth.SetExpired(jwtConfig.Expired))
	opts = append(opts, jwtauth.SetSigningKey([]byte(jwtConfig.SigningKey)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(jwtConfig.SigningKey), nil
	}))
	var method jwt.SigningMethod
	switch jwtConfig.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	opts = append(opts, jwtauth.SetSigningMethod(method))

	var store jwtauth.Storer
	switch jwtConfig.Store {
	case "redis":

	default:
		s, err := buntdb.NewStore(jwtConfig.FilePath)
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
