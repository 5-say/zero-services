package logic

import (
	"context"
	"time"

	"github.com/5-say/go-tools/random"
	"github.com/5-say/zero-services/services/jwtx"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/svc"
	"github.com/golang-jwt/jwt"

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
	var (
		secretKey       = in.SecretKey
		iat             = time.Now().Local().Unix()
		seconds         = in.ExpSecond
		randomAccountID = random.Simple().Encode(int64(in.AccountID), 8)
	)

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["rai"] = randomAccountID
	claims["grp"] = in.Group
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenStr, err := token.SignedString([]byte(secretKey))

	return &jwtx.MakeToken_Response{
		Token: tokenStr,
	}, err
}
