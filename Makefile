

run:
	go run ./cmd/

build-image:
	docker build -t efner/session-microservice:1.0 .

deploy: build-image
	docker push efner/session-microservice:1.0

.PHONY: run build-image deploy