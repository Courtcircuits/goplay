build:
  go build -o bin/goplay cmd/main.go
dev:
  go run cmd/main.go
run:
  just build
  ./bin/main
