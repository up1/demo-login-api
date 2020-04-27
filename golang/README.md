# How to run ?
```
$go run main.go
```

## Testing and code coverage
```
$go test ../..
$go test ./... --cover
$go test ./...  -coverprofile=coverage.out
$go tool cover -html=coverage.out
```