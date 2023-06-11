
help:
	@echo ""
	@echo "------------------------------------------------"
	@echo "make j   |  make jwtx-run  :  运行 jwtx     服务"
	@echo "make jr  |  make jwtx-rpc  :  生成 jwtx rpc 文件"
	@echo "make jd  |  make jwtx-dao  :  生成 jwtx dao 文件"
	@echo "------------------------------------------------"
	@echo ""

run:
	make help

j:jwtx-run
jwtx-run:
	TZ='Asia/Shanghai'; export TZ
	cd private/jwtx/rpc && go run jwtx.go

jr:jwtx-rpc
jwtx-rpc:
	goctl rpc protoc define/jwtx.proto -I=define --go_out=public --go-grpc_out=public --zrpc_out=private/jwtx/rpc --style goZero

jd:jwtx-dao
jwtx-dao:
	cd private/jwtx/db && go run main.go -f ../rpc/etc/jwtx.yaml
