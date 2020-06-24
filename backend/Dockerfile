FROM golang:1.12-alpine as builder

WORKDIR /app
COPY . .
RUN ls

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -a -o ./main ./main.go

FROM alpine:3.9
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app .
RUN ls

ENTRYPOINT ./main