FROM golang:1.18-alpine AS builder
WORKDIR /app/dbo-golang
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o dbo-golang.exe

FROM alpine
RUN apk update && apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/dbo-golang /app/
ENTRYPOINT ["./dbo-golang.exe"]