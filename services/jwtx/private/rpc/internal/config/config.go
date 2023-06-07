package config

import (
	"github.com/5-say/go-tools/tools/db"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB db.Config
}
