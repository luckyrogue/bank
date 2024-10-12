APP_NAME = bank
PORT = 8080

build:
	@echo "Building the Go application..."
	go build -o $(APP_NAME)

run: build
	@echo "Running the Go application..."
	./$(APP_NAME)

docker-build:
	@echo "Building the Docker image..."
	docker build -t $(APP_NAME) .

docker-run:
	@echo "Running the Docker container..."
	docker run -p $(PORT):$(PORT) $(APP_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

test:
	@echo "Running tests..."
	go test ./...

fmt:
	@echo "Formatting the Go code..."
	go fmt ./...

lint:
	@echo "Linting the Go code..."
	golangci-lint run

help:
	@echo "Available commands:"
	@echo "  make build        - Сборка Go-приложения"
	@echo "  make run          - Сборка и запуск приложения"
	@echo "  make docker-build - Сборка Docker-образа"
	@echo "  make docker-run   - Запуск Docker-контейнера"
	@echo "  make clean        - Очистка скомпилированных файлов"
	@echo "  make test         - Запуск тестов"
	@echo "  make fmt          - Форматирование кода"
	@echo "  make lint         - Линтер кода"

.PHONY: build run docker-build docker-run clean test fmt lint help
