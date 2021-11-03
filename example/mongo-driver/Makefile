all:
	go run main.go

build:
	docker-compose build
up: build
	docker-compose up --detach mongo-server
down:
	docker-compose down
logs:
	docker-compose logs mongo-server
