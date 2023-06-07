.PHONY: dao
.PHONY: sql
.PHONY: rpc

run:
	cd rpc && go run jwtx.go

dao:
	cd db && go run .

sql:
	cd ../define/db && > structure.sql && cat structure/*.sql >> structure.sql

rpc:
	goctl rpc protoc ../define/jwtx.proto -I=../define --go_out=../../ --go-grpc_out=../../ --zrpc_out=./rpc --style goZero
