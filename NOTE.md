Todo
1.healthcheck readiness
6.redis
11.testing
12.DI

Check todo

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

Build in
1.Error wrapper
2.Swagger
3.APM
4.Linter
5.Husky
6.Live reload + air
7.keycloak IAM
8.Golang sqs