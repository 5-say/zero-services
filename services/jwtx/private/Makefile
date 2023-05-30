.PHONY: api
.PHONY: doc
.PHONY: init
.PHONY: tidy
.PHONY: dao
.PHONY: sql
.PHONY: rpc

run: tidy
	cd api && go run api.go

api:
	goctl api go --home ./tpl --api api/define/api.api --dir api --style goZero

doc:
	goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api api/define/api.api -dir api/doc

tidy:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod tidy

init:
	sh shell/install-goctl.sh
	sh shell/install-gopls.sh
	sh shell/install-vscode-extension.sh
	cd api/etc && [ -f api.yaml ] || cp example.api.yaml api.yaml
	make doc
	make api
	make tidy

dao:
	cd db && go run .

sql:
	cd db && > structure.sql && cat sql/structure/*.sql >> structure.sql

rpc:
	goctl rpc protoc define/rpc.proto --go_out=../../ --go-grpc_out=../../ --zrpc_out=./rpc --style goZero
