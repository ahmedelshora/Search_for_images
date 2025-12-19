.PHONY: up down build logs clean restart

# Start all services
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

# Build and start services
build:
	docker-compose up -d --build

# View logs
logs:
	docker-compose logs -f app

# Clean everything (including volumes)
clean:
	docker-compose down -v
	rm -rf images/*

# Restart the app
restart:
	docker-compose restart app

# Run the app once
run:
	docker-compose run --rm app