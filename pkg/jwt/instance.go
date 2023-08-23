package jwt

import (
	"sync"

	"github.com/tiktokSpeed/tiktokSpeed/conf"
)

var (
	once    sync.Once
	jwtInst *JWT
)

func GetJWTInst() *JWT {
	once.Do(func() {
		jwtInst = NewJWT([]byte(conf.GetConf().Jwt.SignKey))
	})
	return jwtInst
}

func CreateToken(claims CustomClaims) (string, error) {
	return GetJWTInst().CreateToken(claims)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	return GetJWTInst().ParseToken(tokenString)
}
