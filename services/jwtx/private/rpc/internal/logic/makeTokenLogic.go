package logic

import (
	"context"
	"time"

	"github.com/5-say/zero-services/services/jwtx"
	"github.com/5-say/zero-services/services/jwtx/private/db/model"
	"github.com/5-say/zero-services/services/jwtx/private/db/query"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/svc"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	gormdb, _ := gorm.Open(mysql.Open("root:root@(mysql:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	q := query.Use(gormdb)

	// 清除旧 token，默认单端登录；若 group 混入设备码，则可实现多端登录
	m := q.JwtxToken
	if _, err := m.Where(
		m.TokenGroup.Eq(in.Group),
		m.RandomAccountID.Eq(in.RandomAccountID),
	).Delete(); err != nil {
		return nil, err
	}

	// 创建新 token
	now := time.Now()
	token := model.JwtxToken{
		TokenGroup:      in.Group,
		RandomAccountID: in.RandomAccountID,
		MakeTokenIP:     in.RequestIP,
		CreatedAt:       now,
		RefreshAt:       now,
		ExpirationAt:    now.Add(time.Duration(in.ExpSecond) * time.Second),
	}
	if err := q.JwtxToken.Create(&token); err != nil {
		return nil, err
	}

	// 构造 token 字符串
	claims := make(jwt.MapClaims)
	claims["iat"] = now.Unix()                // 签发时间
	claims["exp"] = now.Unix() + in.ExpSecond // 过期时间
	claims["tid"] = token.ID                  // token ID
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = claims
	tokenStr, err := jwtToken.SignedString([]byte(in.SecretKey))

	return &jwtx.MakeToken_Response{
		Token: tokenStr,
	}, err
}
