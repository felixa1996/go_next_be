Todo
1.healthcheck readiness
2.golang response formatter https://adityarama1210.medium.com/create-golang-api-doc-with-swag-d73be1767d39
3.golang keycloak
 4.golang error
5.golang sqs
6.redis
8.kafka
11.testing

Cmd:
golangci-lint run
go test -cover -coverprofile coverage.out ./...    
go tool cover -func=coverage.out
go tool cover -html=coverage.out

Air


Husky
1.go install github.com/automation-co/husky@latest
2.add alias in zsh alias husky='$(go env GOPATH)/bin/air'
3.husky init
4.husky add pre-commit "golangci-lint run"

Swagger
1. add alias in zsh alias swag='$(go env GOPATH)/bin/swag'
swag init -g app/common/app.go --output docs