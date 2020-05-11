# build state
FROM golang:1.14 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o icqbot cmd/icq/main.go

# final stage
FROM scratch
COPY --from=builder /app/icqbot /usr/bin/
CMD ["icqbot"]