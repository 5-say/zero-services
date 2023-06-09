package logic

import (
	"context"
	"time"

	"github.com/5-say/zero-services/services/jwtx"
	"github.com/5-say/zero-services/services/jwtx/private/db/dao"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/common"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/svc"

	"github.com/5-say/go-tools/tools/t"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验 token
func (l *CheckTokenLogic) CheckToken(in *jwtx.CheckToken_Request) (*jwtx.CheckToken_Response, error) {
	// 解析 token 字符串
	claims, err := common.ParseToken(in.RequestToken, []byte(in.AccessSecret))
	if err != nil {
		return nil, t.RPCError(err.Error(), "parse fail")
	}
	iat := int64(claims["iat"].(float64))
	tid := uint64(claims["tid"].(float64))
	randomAccountID := claims["rai"].(string)

	// 初始化数据库
	q := dao.Common()

	// 查找 token
	m := q.JwtxToken
	token, err := m.Where(m.ID.Eq(tid)).First()
	if err != nil {
		return nil, t.RPCError(err.Error(), "not found")
	}

	// IP 一致性校验
	if in.CheckIP {
		if in.RequestIP != token.MakeTokenIP {
			return nil, t.RPCError("", "ip fail")
		}
	}

	// token 刷新校验
	var (
		newToken = ""
		now      = time.Now()
	)
	// l.Debug("iat", iat)
	// l.Debug("最后刷新时间", token.FinalRefreshAt.Unix())
	// l.Debug("刷新时间间隔", in.RefreshInterval)
	// l.Debug("上次刷新时间", token.LastRefreshAt.Unix())
	// l.Debug("并发容错时间", in.FaultTolerance)
	if iat == token.FinalRefreshAt.Unix() { // token 未刷新

		// 需要刷新 token
		if iat+in.RefreshInterval < now.Unix() {

			// 构造 token 字符串（过期时间不变，签发时间顺延）
			newToken, err = common.MakeTokenStr(in.AccessSecret, now, token.ExpirationAt, token.ID, randomAccountID)
			if err != nil {
				return nil, t.RPCError(err.Error(), "make fail")
			}

			// 更新数据库
			_, err = m.Where(m.ID.Eq(token.ID)).UpdateSimple(
				m.LastRefreshAt.Value(token.FinalRefreshAt),
				m.FinalRefreshAt.Value(now),
			)
			if err != nil {
				return nil, t.RPCError(err.Error(), "refresh fail")
			}
		}

		// 验证通过
		return &jwtx.CheckToken_Response{
			NewToken:        newToken,
			RandomAccountID: randomAccountID,
		}, nil

	} else if iat == token.LastRefreshAt.Unix() { // token 已刷新

		// 当前时间 未超出 并发容错时间（允许继续使用）
		if now.Unix() < token.FinalRefreshAt.Unix()+in.FaultTolerance {
			// 验证通过
			return &jwtx.CheckToken_Response{
				NewToken:        "",
				RandomAccountID: randomAccountID,
			}, nil
		}
	}

	return nil, t.RPCError("", "token refreshed")
}
