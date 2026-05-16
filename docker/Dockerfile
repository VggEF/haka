# Этап 1: Сборка
FROM golang:1.21-alpine AS builder

# Установка необходимых пакетов
RUN apk add --no-cache git ca-certificates

# Установка рабочей директории
WORKDIR /app

# Копирование go.mod и go.sum
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download
RUN go mod verify

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/bin/student-app ./cmd/api

# Этап 2: Финальный образ
FROM alpine:3.18

# Установка CA certificates и timezone
RUN apk --no-cache add ca-certificates tzdata

# Создание рабочей директории
WORKDIR /root/

# Копирование бинарного файла из этапа сборки
COPY --from=builder /app/bin/student-app .

# Копирование конфигурационных файлов
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/migrations ./migrations

# Создание папок для загрузок и логов
RUN mkdir -p uploads logs

# Открытие порта
EXPOSE 8080

# Запуск приложения
CMD ["./student-app"]