sub_track
A lightweight REST API for managing subscriptions.
The service stores subscription records in PostgreSQL and exposes CRUD endpoints plus a placeholder for calculating the total sum of active subscriptions.

The project was built with Go 1.22+, Gin, Jmoiron/SQLX and Viper.

Table of Contents
Section	Description
Features	What the service does
Architecture	How the code is structured
Installation	How to build & run locally
Configuration	YAML & env vars
Database Migration	Running the SQL scripts
API Reference	End‑points, payloads & examples
Error Handling	Custom error type
Testing	Run the test suite
Development	Common commands & guidelines
License	© 2024
Features
Feature	Description
CRUD	Create, read, update and delete subscriptions
Pagination & filtering	(Future work)
Total sum calculation	Endpoint stub – ready for business logic
Validation	Built‑in Gin binding + custom validator for isDateValid
Logging	Structured slog logger
Clean architecture	Separation of concerns: handler → service → repository
Architecture
.
├── cmd
│   └── app
│       └── main.go          # entry point
├── internal
│   ├── handler             # HTTP handlers (Gin)
│   ├── repository          # Database access
│   └── service             # Business logic
├── schema                  # SQL migration files
├── server.go               # HTTP server wrapper
├── sub_track               # Public package (models, errors)
├── configs
│   └── config.yaml         # Default config
├── go.mod
└── README.md
main.go – loads config, connects to PostgreSQL, wires up repository → service → handler, starts the HTTP server.
handler – routes, validation, JSON marshalling and error handling.
repository – SQLX wrappers, date conversions and PostgresNormalDate helper.
service – orchestrates repository calls and implements business rules.
sub_track – shared structs (Subscription, SubscriptionUpdate, SumResponse) and the ErrorPlus interface.
Installation
Prerequisites

Go 1.22 or newer
PostgreSQL 13+
(Optional) Docker & docker‑compose for dev
1. Clone the repo
git clone https://github.com/BOBAvov/sub_track.git
cd sub_track
2. Install dependencies
go mod download
3. Set up the database
# Create a new PostgreSQL user & database
psql -U postgres
CREATE USER sub_user WITH PASSWORD 'qwerty';
CREATE DATABASE sub_db OWNER sub_user;

# Apply migrations
psql -U sub_user -d sub_db -f schema/000001_init.up.sql
Tip – If you prefer Docker, run:

docker run --name sub_pg -e POSTGRES_USER=sub_user -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=sub_db -p 5436:5432 -d postgres:15
4. Configure the service
Copy the default config and edit if needed:

cp configs/config.yaml configs/config.yaml
# edit configs/config.yaml to match your DB host/port/etc.
Environment variables can override config values (e.g., PORT, DB_HOST, …).
See the Configuration section for details.

5. Run the server
go run ./cmd/app
The service will be reachable at http://localhost:8080.

Configuration
The project uses Viper to read a config.yaml file located in the configs/ folder.
All values can also be overridden by environment variables.

Section	Key	Type	Default	Example
port	port	string	8080	9090
db	host	string	localhost	db.example.com
port	string	5436	5432
username	string	postgres	sub_user
password	string	qwerty	supersecret
dbname	string	postgres	sub_db
sslmode	string	disable	require
Example env override

export PORT=9090
export DB_HOST=db.example.com
export DB_USER=sub_user
export DB_PASSWORD=supersecret
go run ./cmd/app
Database Migration
The only migration currently is the subs table.
Feel free to add more files to the schema/ folder and follow the naming convention:

000001_init.up.sql – apply migration
000001_init.down.sql – rollback
If you ever need to reset the database:

psql -U sub_user -d sub_db -f schema/000001_init.down.sql
psql -U sub_user -d sub_db -f schema/000001_init.up.sql
API Reference
Base URL
http://localhost:8080/api
1. Create Subscription
Method	Endpoint	Body	Response
POST	/subs/	```json	
{			
"user_id": "11111111-2222-3333-4444-555555555555",			
"service_name": "Premium",			
"price": 99,			
"start_date": "2024-01-01",			
"end_date": "2024-12-31"			
}			
```	200 OK
{"id": 1}		
user_id must be a valid UUID.
Dates must be in YYYY-MM-DD format (validator isDateValid).

2. Get All Subscriptions
Method	Endpoint	Response
GET	/subs/	200 OK
{ "data": [ {...}, {...} ] }
3. Get Subscription By ID
Method	Endpoint	Response
GET	/subs/:id	200 OK
{ ...subscription... }
4. Update Subscription
Method	Endpoint	Body	Response
PUT	/subs/:id	```json	
{			
"price": 120,			
"end_date": "2025-06-30"			
}			
```	200 OK
{ "status": "ok" }		
Only price and/or end_date may be provided.

5. Delete Subscription
Method	Endpoint	Response
DELETE	/subs/:id	200 OK
{ "status": "ok" }
6. Total Sum (stub)
The endpoint exists but currently returns a 400 Bad Request until the business logic is added.

curl -X POST http://localhost:8080/api/total/ -d '{"start_date":"2024-01-01","end_date":"2024-12-31","users_id":[1,2],"services_name":["Premium"]}'
Error Handling
The project uses a custom error type:

type MyError struct {
    Text     string
    Location string
}
All repository errors are wrapped in sub_track.MyError and forwarded to the HTTP layer. The handler will translate these into a JSON error response:

{ "message": "database connect error" }
The HTTP status code is set according to the context (e.g., 500 for server errors).

Testing
Run the unit test suite:

go test ./...
Note: No integration tests are included yet – feel free to add them.

Development
Task	Command
Lint & format	go vet ./... && go fmt ./...
Generate go.mod	go mod tidy
Run with env vars	export DB_HOST=... ; go run ./cmd/app
Build binary	go build -o bin/sub_track ./cmd/app
Run the binary	./bin/sub_track
Code Quality
Linting: golangci-lint run
Static Analysis: go vet
Contribution
Fork the repo.
Create a feature branch (git checkout -b feat/xxx).
Run tests & linters.
Submit a pull request.
License
MIT © 2025 BOBAvov
(See LICENSE file for details)

Happy hacking! 🚀

# SubTrack  

> Удобный сервис для хранения и аналитики подписок.  
> Реализован на Go, использует Gin‑фреймворк, Viper для конфигурации и PostgreSQL как БД.  

---

## 📦 Структура репозитория

```
SubTrack/
├─ cmd/
│  └─ app/
│     └─ main.go          # точка входа, инициализация всего приложения
├─ configs/
│  └─ config.yaml        # параметры запуска
├─ internal/
│  ├─ handler/           # HTTP‑контроллеры
│  ├─ repository/        # доступ к БД (PostgreSQL)
│  └─ service/           # бизнес‑логика
├─ schema/               # миграции
│  ├─ 000001_init.up.sql
│  └─ 000001_init.down.sql
├─ server.go             # обёртка над http.Server
├─ todo_sub.go           # модели, которые отдаём через API
├─ errors.go             # собственный тип ошибки
└─ go.mod / go.sum
```

---

## ⚙️ Что делает приложение

| Маршрут | Метод | Описание |
|---------|-------|----------|
| `/api/subs` | `POST` | Создать подписку |
| `/api/subs` | `GET`  | Получить список всех подписок |
| `/api/subs/:id` | `GET`  | Получить подписку по id |
| `/api/subs/:id` | `PUT`  | Обновить `price` и/или `end_date` |
| `/api/subs/:id` | `DELETE` | Удалить подписку |
| `/api/total` | `POST` | (TODO) подсчитать сумму по дате и фильтрам |

Все запросы/ответы сериализуются в JSON.

---

## 🚀 Как запустить

> Требуется Go 1.22+ и PostgreSQL 15+

### 1. Клонирование

```bash
git clone https://github.com/BOBAvov/sub_track.git
cd sub_track
```

### 2. Настройка БД

1. Установите PostgreSQL.  
2. Создайте базу `postgres` (или другую и укажите в `configs/config.yaml`).  
3. Примените миграции:

```bash
psql -U postgres -d postgres -f schema/000001_init.up.sql
```

> При необходимости можно написать простую скриптовую миграцию с помощью `golang-migrate`.

### 3. Конфигурация

Файл `configs/config.yaml` содержит:

```yaml
port : "8080"
db :
    host: "localhost"
    port: "5436"
    username: "postgres"
    password: "qwerty" 
    dbname: "postgres"
    sslmode: "disable"
```

Замените значения под свою среду.

### 4. Сборка и запуск

```bash
go build -o sub_track cmd/app/main.go
./sub_track
```

Приложение запустится на `http://localhost:8080`.

> В продакшн‑окружении обычно запускают `sub_track` в контейнере Docker:

```yaml
# Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o sub_track cmd/app/main.go

FROM alpine
COPY --from=builder /app/sub_track /usr/local/bin/sub_track
CMD ["sub_track"]
```

```bash
docker build -t sub_track .
docker run -p 8080:8080 -e DB_HOST=pg -e DB_PORT=5432 ... sub_track
```

---

## 🔧 Технические детали

### Архитектура

```
┌──────────────────────┐
│   HTTP Router (Gin)   │
│    (handler)          │
└───────┬────────────────┘
        │
┌───────▼────────────────┐
│     Service Layer      │
│ (business logic)       │
└───────┬────────────────┘
        │
┌───────▼────────────────┐
│   Repository Layer     │
│  (PostgreSQL access)   │
└────────────────────────┘
```

* **handler** – отвечает за маршалинг JSON и валидацию входящих данных.  
* **service** – бизнес‑логика (например, проверка входных параметров).  
* **repository** – прямой доступ к БД; реализует CRUD через `sqlx`.

### Логирование

`log/slog` используется по всему проекту. В консоль выводятся сообщения уровня `INFO`, `ERROR`.  

### Валидация

Для полей `start_date` и `end_date` используется кастомный валидатор `isDateValid` (формат `MM.yyyy`).  

---

## 📚 API‑примеры

### Создание подписки

```bash
curl -X POST http://localhost:8080/api/subs \
  -H "Content-Type: application/json" \
  -d '{
        "user_id":"123e4567-e89b-12d3-a456-426614174000",
        "service_name":"Netflix",
        "price": 1499,
        "start_date":"01.2023",
        "end_date":"01.2024"
      }'
```

Ответ:

```json
{"id":1}
```

### Получить все подписки

```bash
curl http://localhost:8080/api/subs
```

```json
{
  "data":[
    {
      "id":1,
      "user_id":"123e4567-e89b-12d3-a456-426614174000",
      "service_name":"Netflix",
      "price":1499,
      "start_date":"2023-01-01",
      "end_date":"2024-01-28"
    }
  ]
}
```

### Обновление подписки

```bash
curl -X PUT http://localhost:8080/api/subs/1 \
  -H "Content-Type: application/json" \
  -d '{
        "price": 1599,
        "end_date":"01.2025"
      }'
```

Ответ:

```json
{"status":"ok"}
```

---

## 🔎 Как это работает

* В `internal/repository/postgres.go` реализованы функции `PostgresNormalDate` и `PostgresNormalDate` – конвертируют строку `MM.yyyy` в полноценную дату.  
* Таблица `subs` хранит все поля: `user_id`, `service_name`, `price`, `start_date`, `end_date`.  
* При обновлении можно менять только `price` и/или `end_date`.  
* При запросе `GET /api/subs` данные возвращаются в виде структуры `sub_track.Subscription`.  

---

## 📦 Зависимости

| Пакет | Версия |
|-------|--------|
| gin-gonic/gin | latest |
| spf13/viper | latest |
| go-playground/validator/v10 | latest |
| jmoiron/sqlx | latest |
| lib/pq | latest |
| log/slog | стандартный пакет (Go 1.22+) |

Все зависимости управляются через Go Modules (`go.mod`).

---

## 🧪 Тесты

> Пока тесты не реализованы, но при необходимости можно добавить unit‑тесты для сервисов и репозиториев.

```bash
go test ./...
```

---

