# Dev Dependencies
godep:
	@go get -u github.com/golang/dep/cmd/dep
	dep ensure -v -vendor-only

pretty:
	gofmt -s -w .

run:
	go run main.go

bin: pretty
	go build -o go_chat main.go

coverage:
	go test -race -v -cover -coverprofile=coverage.out $$(go list ./... | grep -Ev "cmd")

cover:
	go tool cover -html=coverage.out

test:
	go test ./...

dep:
	dep ensure -v -vendor-only