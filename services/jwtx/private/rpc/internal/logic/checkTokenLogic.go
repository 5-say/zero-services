package logic

import (
	"context"
	"time"

	"github.com/5-say/zero-services/services/jwtx"
	"github.com/5-say/zero-services/services/jwtx/private/db/query"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/svc"
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

	// 确认是否需要刷新
	now := time.Now()
	if in.Iat+in.RefreshInterval < now.Unix() { // 当前 token 需要刷新

		if token.RefreshAt.Unix() == in.Iat { // 当前 token 有效
			// 更新数据库
			// 生成新 token
		} else { // 当前 token 无效，token 已被并发刷新
			// 确认并发容错时间
			if false {
				// 超出并发容错时间，返回错误信息
			}
		}
	}

	return &jwtx.CheckToken_Response{}, nil
}
