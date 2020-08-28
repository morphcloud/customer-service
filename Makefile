.PHONY: build
build:
	go build -o ./bin/app ./main.go

.PHONY: run
run:
	./bin/app

.PHONY: build-and-run
build-and-run: build run

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: test-bench
test-bench:
	go test -bench=. ./...

.PHONY: test-cover
test-cover:
	go test -cover ./...

.PHONY: format
format:
	gofmt -w ./..

.PHONY: docker-build
docker-build:
	docker build -t hzhyvinskyi/morphcloud-customer-service:1.0.0 .

.PHONY: docker-run
docker-run:
	docker run -d -p 8082:8082 --name morphcloud-customer-service-container hzhyvinskyi/morphcloud-customer-service:1.0.0

.PHONY: docker-push
docker-push:
	docker push hzhyvinskyi/morphcloud-customer-service:1.0.0

.DEFAULT_GOAL := build-and-run
