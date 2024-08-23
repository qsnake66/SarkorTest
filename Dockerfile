# Установка базового образа
FROM golang:1.22

# Установка рабочей директории
WORKDIR /app

# Копирование файлов
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Сборка приложения
RUN go build -o main ./cmd

# Запуск приложения
CMD ["./main"]