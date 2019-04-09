setup:
	go get github.com/golang/dep/cmd/dep
	go get golang.org/x/lint/golint

build: 
	go build

test:
	go test $(go list ./... | grep -v /vendor)

dep: setup
	dep ensure

