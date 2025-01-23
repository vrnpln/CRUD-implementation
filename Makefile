# Makefile для создания миграций
.PHONY: migrate migrate-down migrate-new run

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:vanere39@localhost:5432/postgres?sslmode=disable"

# Абсолютный путь к папке с миграциями
MIGRATE := migrate -source "file:///Users/voronetskaya/Desktop/Project1/migrations" \
        -database "postgres://postgres:vanere39@localhost:5432/postgres?sslmode=disable"

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations -seq $(NAME)

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down
	
# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер
	
