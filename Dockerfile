# Этап  1: Сборка зависимостей
FROM golang:1.19-alpine AS builder

# Установка рабочей директории
WORKDIR /app

# Копирование исходного кода в контейнер
COPY . .

# Установка зависимостей
RUN go mod download

# Этап  2: Сборка приложения
FROM builder AS appbuilder

# Компиляция приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main cmd/editor/main.go

# Этап  3: Запуск приложения
FROM alpine:latest

# Установка рабочей директории
WORKDIR /app

# Копирование исполняемого файла из этапа сборки
COPY --from=appbuilder /app/main /app/main

# Установка gRPC reflection
RUN apk add --no-cache ca-certificates

# Запуск приложения
CMD ["/app/main"]