package logic

import (
	"context"
	"time"

	"main/base/service/jwtx"
	"main/base/service/jwtx/private/rpc/internal/svc"

	"github.com/5-say/go-tools/random"
	"github.com/golang-jwt/jwt/v4"
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
	// todo: add your logic here and delete this line

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
