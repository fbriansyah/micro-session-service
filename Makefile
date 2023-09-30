

run:
	go run ./cmd/

build-image:
	docker build -t efner/session-microservice:1.0 .

.PHONY: run build-image