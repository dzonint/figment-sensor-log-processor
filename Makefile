.PHONY: test
test:
	go clean -testcache ./... && go test -race -v ./...

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: build
build: vendor
	go build -trimpath -o build/figment-sensor-log-processor ./cmd/figment-sensor-log-processor/*.go

.PHONY: clean
clean:
	rm -rf build

.PHONY: run
run: build
	build/figment-sensor-log-processor

.PHONY: run-docker
run-docker:
	docker build -t figment-sensor-log-processor .
	docker run --rm figment-sensor-log-processor
