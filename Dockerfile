# Build bosqichi
FROM golang:1.22.4-alpine  AS builder

# Ishchi katalog
WORKDIR /app

# Go module fayllarini nusxalash
COPY go.mod go.sum ./

# Dependency larni yuklab olish
RUN go mod download

# Loyiha kodini nusxalash
COPY . .

# Ilovani build qilish
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Runtime image
FROM alpine:latest

# Kerakli utilities o'rnatish
RUN apk --no-cache add ca-certificates

# Ishchi katalog
WORKDIR /root/

# Build bosqichidan binary faylni nusxalash
COPY --from=builder /app/main .

# Port ochish
EXPOSE 8080

# Ilovani ishga tushirish
ENTRYPOINT ["./main"]