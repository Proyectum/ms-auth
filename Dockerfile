FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go mod download
RUN GOARCH=amd64 GOOS=linux go build -o application ./cmd/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/application .
COPY resources/ ./resources/

RUN chmod +x ./application

EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./application"]