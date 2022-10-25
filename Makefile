test:
	go test -v ./... -coverprofile coverage.out

test-html:
	go tool cover -html=coverage.out

build:
	go build -v ./...
