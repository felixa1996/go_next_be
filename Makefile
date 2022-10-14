BINARY_NAME=go_next_be

build:
	GOARCH=amd64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin app/main.go
	GOARCH=amd64 GOOS=linux go build -o build/${BINARY_NAME}-linux app/main.go
	GOARCH=amd64 GOOS=window go build -o build/${BINARY_NAME}-windows app/main.go

run:
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	# go test ./...
	go test -v -race -covermode=atomic -coverprofile coverage.out ./...

lint:
	golangci-lint run ./...

oas:
	swag init -g app/common/app.go --output docs

test-coverage:
	go test -race -covermode=atomic ./... -coverpkg=./... -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
	# open coverage/coverage.html

test-coverage-verbose:
	go test -v -race -covermode=atomic ./... -coverpkg=./... -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
	# open coverage/coverage.html