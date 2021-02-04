SERVICE_NAME=github.com/echoturing/go-yuque-sdk

.PHONY:fmt
fmt:
	@find . -name "*.go" | xargs goimports -w -l --local $(SERVICE_NAME) --private "mockprivate"