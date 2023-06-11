package logic

import (
	"context"

	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/private/jwtx/db/dao"
	"github.com/5-say/zero-services/private/jwtx/rpc/internal/svc"
	"github.com/5-say/zero-services/public/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 移除 token（安全退出）
func (l *DeleteTokenLogic) DeleteToken(in *jwtx.DeleteToken_Request) (*jwtx.DeleteToken_Response, error) {
	// 初始化数据库
	q := dao.Common()

	// 移除 token
	m := q.JwtxToken
	_, err := m.Where(m.ID.Eq(in.TokenID)).Delete()
	if err != nil {
		return nil, t.RPCError(err.Error(), "delete fail")
	}

	return &jwtx.DeleteToken_Response{}, nil
}
