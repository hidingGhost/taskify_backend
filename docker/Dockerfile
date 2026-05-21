FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Hier der Fix: Baue aus dem scripts-Ordner
RUN go build -o taskify-backend ./scripts/*.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/taskify-backend .
EXPOSE 8080
CMD ["./taskify-backend"]