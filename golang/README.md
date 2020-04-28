# How to run ?
```
$go run main.go
```

## Testing and code coverage
```
$go test ./...
$go test ./... --cover
$go test ./...  -coverprofile=coverage.out
$go tool cover -html=coverage.out
```

## Lint with go (Static code analysis)
```
$go get -u golang.org/x/lint/golint
$golint ./...
```
