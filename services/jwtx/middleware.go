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

func RefreshTokenMiddleware(
	errHandler func(err error) error,
	simpleRandomConfig random.SimpleRandomConfig,
	jwtxConfig JWTXConfig,
	jwtxRpc JwtxClient,
) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// 校验 token（拓展校验、刷新 token）
			resp, err := jwtxRpc.CheckToken(r.Context(), &CheckToken_Request{
				AccessSecret:    jwtxConfig.AccessSecret,
				RefreshInterval: jwtxConfig.RefreshInterval,
				FaultTolerance:  jwtxConfig.FaultTolerance,
				CheckIP:         jwtxConfig.CheckIP,
				RequestIP:       ip.GetRequestIP(r),
				RequestToken:    r.Header.Get("Authorization"),
			})

			if err != nil {

				// fmt.Println(err, resp) // 可以通过外部传参补充日志

				httpx.ErrorCtx(r.Context(), w, errHandler(err))

			} else {
				fmt.Println("- - - - - - - - - ", resp) // 可以通过外部传参补充日志

				// 响应头填充刷新后的 token
				w.Header().Add("NEW-TOKEN", resp.NewToken)

				// 解析真实的账户 ID
				accountID := random.Simple(simpleRandomConfig).Decode(resp.RandomAccountID)

				// 附加上下文
				ctx := r.Context()
				ctx = context.WithValue(ctx, "accountID", accountID)

				// 继续执行
				next(w, r.WithContext(ctx))
			}

		}
	}
}
