package jwtx

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func RefreshTokenMiddleware(config JWTXConfig, jwtxRpc JwtxClient) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			logx.Infof("iat: %v", r.Context().Value("iat"))
			logx.Infof("exp: %v", r.Context().Value("exp"))
			logx.Infof("tid: %v", r.Context().Value("tid"))

			resp, err := jwtxRpc.CheckToken(r.Context(), &CheckToken_Request{
				SecretKey: config.AccessSecret,
			})
			if err != nil {
			}

			fmt.Println(err, resp)

			next(w, r)
		}
	}
}
