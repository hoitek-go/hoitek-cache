testcov:
	go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
test:
	go test ./... -v