package jwtx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/5-say/go-tools/tools/ip"
	"github.com/5-say/go-tools/tools/random"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshTokenMiddleware(simpleRandomConfig random.SimpleRandomConfig, jwtxConfig JWTXConfig, jwtxRpc JwtxClient) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// 获取客户端 IP 地址
			requestIP := ip.GetRequestIP(r)

			// 校验 token（拓展校验、刷新 token）
			resp, err := jwtxRpc.CheckToken(r.Context(), &CheckToken_Request{
				AccessSecret:    jwtxConfig.AccessSecret,
				RefreshInterval: jwtxConfig.RefreshInterval,
				FaultTolerance:  jwtxConfig.FaultTolerance,
				CheckIP:         jwtxConfig.CheckIP,
				RequestIP:       requestIP,
				Tid:             r.Context().Value("tid").(uint64),
				Iat:             r.Context().Value("iat").(int64),
				Exp:             r.Context().Value("exp").(int64),
			})

			if err != nil {

				fmt.Println(err, resp) // 可以通过外部传参补充日志
				httpx.ErrorCtx(r.Context(), w, err)

			} else {

				// 响应头填充刷新后的 token
				w.Header().Add("NEW-TOKEN", resp.NewToken)

				// 解析真实的账户 ID
				accountID := random.Simple(simpleRandomConfig).Decode(resp.RandomAccountID)

				// 附加上下文
				ctx := r.Context()
				ctx = context.WithValue(ctx, "requestIP", requestIP)
				ctx = context.WithValue(ctx, "accountID", accountID)

				// 继续执行
				next(w, r.WithContext(ctx))
			}

		}
	}
}
