# chgk-bot

![Go](https://github.com/zetraison/chgk-bot/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/zetraison/chgk-bot)](https://goreportcard.com/report/github.com/zetraison/chgk-bot)
![Docker Telegram Bot](https://github.com/zetraison/chgk-bot/workflows/Docker%20Telegram%20Bot/badge.svg)
![Docker ICQ Bot](https://github.com/zetraison/chgk-bot/workflows/Docker%20ICQ%20Bot/badge.svg)


Telegram Bot based on http://db.chgk.info questions database.

## Features
Chgk bot supports next commands:

- `/start` - starts bot and shows help
- `/help` - shows help
- `/question` - sends random question to chat
- `/round` - starts round of game
- `/stop` - stops round of game
- `/score` - shows results

## Building
You need either Docker and make, or go in order to build binary.

### Build with Go
```bash
GOOS=linux GOARCH=amd64 go build -o bin/chgk-telegram-bot cmd/telegram/main.go
```
```bash
GOOS=linux GOARCH=amd64 go build -o bin/chgk-icq-bot cmd/icq/main.go
```
or
```bash
make compile
```

### Build with Docker
```bash
docker build -f build/Dockerfile.telegram -t zetraison/chgk-telegram-bot .
```
```bash
docker build -f build/Dockerfile.icq -t zetraison/chgk-icq-bot .
```
or
```bash
make docker_build
```

## Running
You need setup ENV variables `$(TELEGRAM_BOT_TOKEN)` and `$(ICQ_BOT_TOKEN)` in `.env` file or in your environment to run binaries.

### Run with Go
```bash
export TELEGRAM_BOT_TOKEN=$(TELEGRAM_BOT_TOKEN) go run cmd/telegram/main.go
```
```bash
export ICQ_BOT_TOKEN=$(ICQ_BOT_TOKEN) go run cmd/icq/main.go
```

or

```bash
make docker_run_telegram_bot
make docker_run_icq_bot
```

### Run with Docker
```bash
docker run -it --rm -e TELEGRAM_BOT_TOKEN=$(TELEGRAM_BOT_TOKEN) zetraison/chgk-telegram-bot
```
```bash
docker run -it --rm -e ICQ_BOT_TOKEN=$(ICQ_BOT_TOKEN) zetraison/chgk-icq-bot
```

or

```bash
make docker_run_telegram_bot
make docker_run_icq_bot
```
