all: bindata fmt ui

ui:
	go install .

bindata: bindata-deps
	-rm assets/assets.go
	go-bindata -o assets/assets.go -pkg assets assets/...

bindata-deps:
	go get github.com/jteeuwen/go-bindata/...

fmt:
	go fmt ./...

.PHONY: all ui bindata bindata-deps fmt
