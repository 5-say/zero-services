package jwtx

import (
	"context"
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

				// 通过外部回调 加工错误对象、记录错误日志
				appErr := errHandler(err)
				// 为响应对象写入错误信息
				httpx.ErrorCtx(r.Context(), w, appErr)

			} else {

				// 响应头填充刷新后的 token
				w.Header().Add("NEW-TOKEN", resp.NewToken)

				// 解析真实的账户 ID
				accountID := random.Simple(simpleRandomConfig).Decode(resp.RandomAccountID)

				// 附加上下文
				ctx := r.Context()
				ctx = context.WithValue(ctx, "accountID", accountID)
				ctx = context.WithValue(ctx, "tokenID", resp.TokenID)

				// 继续执行
				next(w, r.WithContext(ctx))
			}

		}
	}
}
