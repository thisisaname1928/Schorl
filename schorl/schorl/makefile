OUTPUT=../initramfs/init
GO_SRC=$(shell find . -type f -name "*.go")

$(OUTPUT): $(GO_SRC)
	@echo "[GOC] -> schorlSysInit"
	@go build .
	@cp schorlSysInit $(OUTPUT)