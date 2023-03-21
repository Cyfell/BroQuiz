# BroQuiz
Le quiz des Bros

## How to launch

    docker-compose -f deployments/docker-compose.yml up

Open website on [http://localhost:8080/infos](http://localhost:8080/infos) to check up.

## How to test

```go
# Answer a question
go run ./cmd/broquizctl answer 1

# More commands
go run ./cmd/broquizctl help
```
