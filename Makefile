# enable module support across all go commands.
export GO111MODULE = on

version = 1.0.0

compile:
	@echo "Compile icq app..."
	GOOS=linux GOARCH=386 go build -o bin/icq-bot-${version}-linux-386 cmd/icq/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/icq-bot-${version}-linux-amd64 cmd/icq/main.go
	GOOS=darwin GOARCH=386 go build -o bin/icq-bot-${version}-darwin-386 cmd/icq/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/icq-bot-${version}-darwin-amd64 cmd/icq/main.go
	GOOS=windows GOARCH=386 go build -o bin/icq-bot-${version}-windows-386.exe cmd/icq/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/icq-bot-${version}-windows-amd64.exe cmd/icq/main.go

	@echo "Compile telegram app..."
	GOOS=linux GOARCH=386 go build -o bin/telegram-bot-${version}-linux-386 cmd/telegram/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/telegram-bot-${version}-linux-amd64 cmd/telegram/main.go
	GOOS=darwin GOARCH=386 go build -o bin/telegram-bot-${version}-darwin-386 cmd/telegram/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/telegram-bot-${version}-darwin-amd64 cmd/telegram/main.go
	GOOS=windows GOARCH=386 go build -o bin/telegram-bot-${version}-windows-386.exe cmd/telegram/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/telegram-bot-${version}-windows-amd64.exe cmd/telegram/main.go

clean:
	rm -rf bin
	rm -rf dist

lint:
	golangci-lint run --fast

run_icq_bot:
	go run cmd/icq/main.go

run_telegram_bot:
	go run cmd/telegram/main.go

docker_build:
	docker build -f build/Dockerfile.telegram -t zetraison/chgk-telegram-bot .
	docker build -f build/Dockerfile.icq -t zetraison/chgk-icq-bot .

docker_run_telegram_bot:
	docker run -it --rm -e TELEGRAM_BOT_TOKEN=<TELEGRAM_BOT_TOKEN> --name chgk_telegram_bot zetraison/chgk-telegram-bot

docker_run_icq_bot:
	docker run -it --rm -e ICQ_BOT_TOKEN=<ICQ_BOT_TOKEN> --name chgk_icq_bot zetraison/chgk-icq-bot

release_dry_run:
	goreleaser --snapshot --skip-publish --rm-dist

release:
	goreleaser