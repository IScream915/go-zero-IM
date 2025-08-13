package ctxData

import "github.com/golang-jwt/jwt/v4"

const Identify = "userCtx"

func GetJwtToken(secretKey string, iat, expiredTime int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expiredTime
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
