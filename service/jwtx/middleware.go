package jwtx

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func RefreshTokenMiddleware(secretKey string, jwtxRpc JwtxClient) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			logx.Infof("exp: %v", r.Context().Value("exp"))
			logx.Infof("iat: %v", r.Context().Value("iat"))
			logx.Infof("rai: %v", r.Context().Value("rai"))
			logx.Infof("grp: %v", r.Context().Value("grp"))

			jwtxRpc.CheckToken(r.Context(), &CheckToken_Request{
				SecretKey: secretKey,
			})
			next(w, r)
		}
	}
}
