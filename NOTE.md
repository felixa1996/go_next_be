1.glide
2.viper

Done
1.air
2.linter


Todo
1.Sample entity
2.Sample Handler
3.Setup sigterm
4.Disconnect mongo
5.log
6.redis
7.env
8.kafka
10.precommit
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