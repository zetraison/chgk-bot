# chgk-bot

![Go](https://github.com/zetraison/chgk-bot/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/zetraison/chgk-bot)](https://goreportcard.com/report/github.com/zetraison/chgk-bot)
![Docker Telegram Bot](https://github.com/zetraison/chgk-bot/workflows/Docker%20Telegram%20Bot/badge.svg)
![Docker ICQ Bot](https://github.com/zetraison/chgk-bot/workflows/Docker%20ICQ%20Bot/badge.svg)


Telegram Bot based on http://db.chgk.info questions database

### Docker
```bash
docker run -it --rm -e TELEGRAM_BOT_TOKEN=<token> zetraison/chgk-telegram-bot
```

```bash
docker run -it --rm -e ICQ_BOT_TOKEN=<token> zetraison/chgk-icq-bot
```
