package logic

import (
	"context"

	"github.com/5-say/zero-services/services/jwtx"
	"github.com/5-say/zero-services/services/jwtx/private/db/dao"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/svc"

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
	_, err := m.Where(m.ID.Eq(in.Tid)).Delete()

	return &jwtx.DeleteToken_Response{}, err
}
