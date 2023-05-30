// Code generated by goctl. DO NOT EDIT.
// Source: rpc.proto

package jwtxclient

import (
	"context"

	"main/base/service/jwtx"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckToken_Request  = jwtx.CheckToken_Request
	CheckToken_Response = jwtx.CheckToken_Response
	MakeToken_Request   = jwtx.MakeToken_Request
	MakeToken_Response  = jwtx.MakeToken_Response

	Jwtx interface {
		// 生成 token
		MakeToken(ctx context.Context, in *MakeToken_Request, opts ...grpc.CallOption) (*MakeToken_Response, error)
		// 校验 token
		CheckToken(ctx context.Context, in *CheckToken_Request, opts ...grpc.CallOption) (*CheckToken_Response, error)
	}

	defaultJwtx struct {
		cli zrpc.Client
	}
)

func NewJwtx(cli zrpc.Client) Jwtx {
	return &defaultJwtx{
		cli: cli,
	}
}

// 生成 token
func (m *defaultJwtx) MakeToken(ctx context.Context, in *MakeToken_Request, opts ...grpc.CallOption) (*MakeToken_Response, error) {
	client := jwtx.NewJwtxClient(m.cli.Conn())
	return client.MakeToken(ctx, in, opts...)
}

// 校验 token
func (m *defaultJwtx) CheckToken(ctx context.Context, in *CheckToken_Request, opts ...grpc.CallOption) (*CheckToken_Response, error) {
	client := jwtx.NewJwtxClient(m.cli.Conn())
	return client.CheckToken(ctx, in, opts...)
}
