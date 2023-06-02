package logic

import (
	"context"
	"time"

	"github.com/5-say/zero-services/services/jwtx"
	"github.com/5-say/zero-services/services/jwtx/private/db/query"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/svc"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	// 初始化数据库
	gormdb, _ := gorm.Open(mysql.Open("root:root@(mysql:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	q := query.Use(gormdb)

	// 查找 token
	m := q.JwtxToken
	token, err := m.Where(m.ID.Eq(in.Tid)).First()
	if err != nil {
		return nil, err
	}

	// IP 一致性校验
	if in.CheckIP {
		if in.RequestIP != token.MakeTokenIP {
			return nil, errors.New("check ip fail")
		}
	}

	// token 刷新校验
	var (
		newToken = ""
		now      = time.Now()
	)
	if in.Iat == token.FinalRefreshAt.Unix() { // token 未刷新

		// 需要刷新 token
		if in.Iat+in.RefreshInterval > now.Unix() {

			// 构造 token 字符串（过期时间不变，签发时间顺延）
			newToken, err = makeTokenStr(in.AccessSecret, now.Unix(), in.Exp, token.ID)
			if err != nil {
				return nil, err
			}

			// 更新数据库
			_, err = m.Where(m.ID.Eq(token.ID)).UpdateSimple(
				m.LastRefreshAt.Value(token.FinalRefreshAt),
				m.FinalRefreshAt.Value(now),
			)
			if err != nil {
				return nil, err
			}
		}

		// 验证通过
		return &jwtx.CheckToken_Response{
			NewToken:        newToken,
			RandomAccountID: token.RandomAccountID,
		}, nil

	} else if in.Iat == token.LastRefreshAt.Unix() { // token 已刷新

		// 当前时间 未超出 并发容错时间（允许继续使用）
		if now.Unix() < token.FinalRefreshAt.Unix()+in.FaultTolerance {
			// 验证通过
			return &jwtx.CheckToken_Response{
				NewToken:        "",
				RandomAccountID: token.RandomAccountID,
			}, nil
		}
	}

	return nil, errors.New("token refreshed")
}
