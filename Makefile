## lint: Lint the code base.
.PHONY: lint
lint:
	golangci-lint run --build-tags integration --disable-all -E revive -E staticcheck -E unused -E gocritic -E gocyclo -E gofmt -E misspell -E stylecheck -E unconvert -E unparam

## test: Run unit tests.
test:
	go test ./...

## build: Build and generate executable.
build:
	docker build . -t mc:dev

## run: Build and run code.
run: build
	docker run -it mc:dev