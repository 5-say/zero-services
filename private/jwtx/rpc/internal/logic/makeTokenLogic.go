package logic

import (
	"context"
	"time"

	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/private/jwtx/db/dao"
	"github.com/5-say/zero-services/private/jwtx/db/dao/model"
	"github.com/5-say/zero-services/private/jwtx/rpc/common"
	"github.com/5-say/zero-services/private/jwtx/rpc/internal/svc"
	"github.com/5-say/zero-services/public/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeTokenLogic {
	return &MakeTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成 token
func (l *MakeTokenLogic) MakeToken(in *jwtx.MakeToken_Request) (*jwtx.MakeToken_Response, error) {
	// 初始化数据库
	q := dao.Common()

	// 清除旧 token，默认单端登录；若 group 混入设备码，则可实现多端登录
	m := q.JwtxToken
	if _, err := m.Where(
		m.TokenGroup.Eq(in.Group),
		m.AccountID.Eq(in.AccountID),
	).Delete(); err != nil {
		return nil, t.RPCError(err.Error(), "clear fail")
	}

	// 创建新 token
	now := time.Now()
	token := model.JwtxToken{
		TokenGroup:     in.Group,
		AccountID:      in.AccountID,
		MakeTokenIP:    in.RequestIP,
		CreatedAt:      now,
		LastRefreshAt:  now,
		FinalRefreshAt: now,
		ExpirationAt:   now.Add(time.Duration(in.AccessExpire) * time.Second),
	}
	if err := q.JwtxToken.Create(&token); err != nil {
		return nil, t.RPCError(err.Error(), "create fail")
	}

	// 构造 token 字符串
	tokenStr, err := common.MakeTokenStr(in.AccessSecret, now, now.Add(time.Second*time.Duration(in.AccessExpire)), token.ID, in.RandomAccountID)
	if err != nil {
		return nil, t.RPCError(err.Error(), "make fail")
	}

	return &jwtx.MakeToken_Response{
		Token: tokenStr,
	}, nil
}
