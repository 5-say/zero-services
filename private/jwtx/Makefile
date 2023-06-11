.PHONY: dao
.PHONY: sql
.PHONY: rpc

run:
	TZ='Asia/Shanghai'; export TZ
	cd rpc && go run jwtx.go

dao:
	cd db && go run .

rpc:
	goctl rpc protoc ../define/jwtx.proto -I=../define --go_out=../../ --go-grpc_out=../../ --zrpc_out=./rpc --style goZero
