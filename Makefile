build:
	go build -o out/tiny_url cmd/tiny_url/main.go

run:
	go run cmd/tiny_url/main.go

build-docker:
	docker compose build

run-docker: build-docker
	docker compose up -d

stop-docker:
	docker compose stop

restart-docker: stop-docker run-docker

