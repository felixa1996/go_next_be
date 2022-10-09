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
	go test ./...

lint:
	golangci-lint run