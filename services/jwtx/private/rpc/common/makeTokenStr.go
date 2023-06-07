package common

import "github.com/golang-jwt/jwt"

// MakeTokenStr .. 构造 token 字符串
func MakeTokenStr(secretKey string, iat, exp int64, tokenID uint64) (tokenStr string, err error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat     // 签发时间
	claims["exp"] = exp     // 过期时间
	claims["tid"] = tokenID // token ID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
