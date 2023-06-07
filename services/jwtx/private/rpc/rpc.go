package main

import (
	"flag"
	"fmt"

	"github.com/5-say/go-tools/tools/db"
	"github.com/5-say/zero-services/services/jwtx"
	"github.com/5-say/zero-services/services/jwtx/private/db/dao"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/config"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/server"
	"github.com/5-say/zero-services/services/jwtx/private/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化公共数据库连接
	dao.SetCommon(c.DB, db.LogWriter())

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		jwtx.RegisterJwtxServer(grpcServer, server.NewJwtxServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
