.PHONY: dev build test clean docker-up docker-down migrate-up migrate-down

# Development
dev:
	docker-compose -f docker/docker-compose.yml up postgres redis -d
	cd backend && air &
	cd frontend && npm run dev

# Build
build:
	cd frontend && npm run build
	cd backend && go build -o bin/server cmd/server/main.go

# Testing
test:
	cd backend && go test ./...
	cd frontend && npm run test

# Database migrations
migrate-up:
	cd backend && migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	cd backend && migrate -path migrations -database "$(DATABASE_URL)" down

# Docker
docker-up:
	docker-compose -f docker/docker-compose.yml up --build -d

docker-down:
	docker-compose -f docker/docker-compose.yml down

docker-logs:
	docker-compose -f docker/docker-compose.yml logs -f

# Clean
clean:
	cd backend && go clean
	cd frontend && rm -rf dist node_modules
	docker-compose -f docker/docker-compose.yml down -v

# Install dependencies
install:
	cd backend && go mod tidy
	cd frontend && npm install

# Generate API docs
docs:
	cd backend && swag init -g cmd/server/main.go -o docs/

# Production deployment
deploy:
	docker-compose -f docker/docker-compose.yml -f docker/docker-compose.prod.yml up --build -d