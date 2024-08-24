FROM golang:1.23 as builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED 0
ENV GOOS linux

RUN go mod download
RUN go build -o service ./cmd/service/main.go



FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/service .

RUN apk add tzdata
RUN ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

ENV CONFIG_PATH "/app/config/config.yaml"

EXPOSE 8080

ENTRYPOINT [ "/app/service" ]
