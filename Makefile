# enable module support across all go commands.
export GO111MODULE = on

compile:
	@echo "Compile icq for linux 386..."
	GOOS=linux GOARCH=386 go build -o build/linux/icq-386 cmd/icq/main.go

	@echo "Compile icq for windows 386..."
	GOOS=windows GOARCH=386 go build -o build/windows/icq-386.exe cmd/icq/main.go

	@echo "Compile icq for macos 386..."
	GOOS=darwin GOARCH=386 go build -o build/mac/icq-386 cmd/icq/main.go

	@echo "Compile telegram for linux amd64..."
	GOOS=linux GOARCH=amd64 go build -o build/linux/telegram-amd64 cmd/telegram/main.go

	@echo "Compile telegram for windows amd64..."
	GOOS=windows GOARCH=amd64 go build -o build/windows/telegram-amd64.exe cmd/telegram/main.go

	@echo "Compile telegram for macos amd64..."
	GOOS=darwin GOARCH=amd64 go build -o build/mac/telegram-amd64 cmd/telegram/main.go

lint:
	go fmt ./...
	go vet -v ./...
	golangci-lint run --fast

run_icq_bot:
	go run cmd/icq/main.go

run_telegram_bot:
	go run cmd/telegram/main.go