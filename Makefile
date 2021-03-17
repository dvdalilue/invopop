
all: get build
	@echo "Done"

get:
	go get -d

build: build-server build-client
	@echo "Build finished"

build-server:
	go build -o server cmd/server/main.go

build-client:
	go build -o client cmd/client/main.go

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go

unit-test:
	go test ./test

clean:
	rm server client