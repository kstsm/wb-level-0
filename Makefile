# Docker
docker:
	cd consumer && docker-compose up -d
stop:
	cd consumer && docker-compose down

# Goose
goose-up:
	cd consumer && goose -dir=./migrations up
goose-down:
	cd consumer && goose -dir=./migrations down

# Fast start
producer:
	cd producer && go run main.go
consumer:
	cd consumer && go run main.go
run-all:
	cd producer && go run main.go &
	sleep 2
	cd consumer && go run main.go







