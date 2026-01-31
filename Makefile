docker-redis:
	# * Pull the latest Redis official image.
	docker pull redis:latest
	# * Launch the Redis container.
	docker run -d --name my-redis -p 6379:6379 redis

build:
	go build -o ./bin/main ./src/main.go

install:
	go mod tidy

run:
	go run ./src/main.go

run-k6:
	k6 run ./src/main-k6.js
