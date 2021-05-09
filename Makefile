account:
	go run api.go account

notify:
	go run api.go account

env:
	docker-compose up -d

gen:
	go generate ./...
