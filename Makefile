.PHONY: init

run:
	go mod tidy

init:
	sh shell/install-goctl.sh
	sh _shell/install-gopls.sh
	sh _shell/install-vscode-extension.sh
