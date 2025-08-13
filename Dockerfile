FROM golang:1.24-bookworm AS builder

WORKDIR /src

# Зависимости отдельно для кеша
COPY go.mod go.sum ./
RUN go mod download

# Исходники
COPY . .

# Сборка бинарника (статически, без CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/subtrack ./cmd/app

# -------- Runtime --------
FROM gcr.io/distroless/static:nonroot

WORKDIR /app

# Режим gin = release
ENV GIN_MODE=release

# Бинарник
COPY --from=builder /out/subtrack /app/subtrack

# Конфигурация и схема (если нужно примонтировать/использовать)
COPY configs ./configs
COPY schema ./schema

# Порт из конфигурации
EXPOSE 8080

# Запуск
ENTRYPOINT ["/app/subtrack"]


#docker build -t subtrack . сборка
#docker run -d --name subtrack -p 8080:8080 subtrack запуск
#docker run --name=todo-db --network=subtrack-net -e POSTGRES_PASSORD='qwerty' -p 5436:5432 -d --rm postgres
#migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up