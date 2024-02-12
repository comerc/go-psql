# Используем официальный образ Golang как базовый
FROM golang:latest as builder

# Копируем исходный код в контейнер
WORKDIR /app
COPY . .

# Собираем приложение
RUN go build -o main .

# Запускаем приложение
CMD ["./main"]

# Запустите команду, которая не завершается
# CMD tail -f /dev/null