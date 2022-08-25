build:
	go install -mod=mod ./...

modvendor:
	go build -mod=mod ./...
	go mod tidy -compat=1.17
	go mod vendor

test: 
	go test -race -cover -timeout=300s ./...

run: build
	echoserver

runhttp: build
	echoserver http