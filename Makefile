lint:
	golangci-lint run

events:
	go run cmd/eventservice/main.go

mongo:
	docker-compose --project-directory ./build/package up