package logic

import (
	"context"

	"main/base/service/jwtx"
	"main/base/service/jwtx/private/rpc/internal/svc"

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
	// todo: add your logic here and delete this line

	return &jwtx.CheckToken_Response{}, nil
}
