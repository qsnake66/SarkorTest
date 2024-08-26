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

# Установка утилиты migrate для выполнения миграций
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && mv migrate.linux-amd64 /usr/local/bin/migrate
# Копирование файлов миграций
COPY db/migrations /app/migrations

# Запуск миграций и приложения
CMD ["sh", "-c", "migrate -path /app/migrations -database postgres://postgres:yourpassword@db:5432/warehouse?sslmode=disable up && ./main"]