DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: deps bindata fmt ui

ui:
	go install .

bindata: bindata-deps
	-rm assets/assets.go
	go-bindata -o assets/assets.go -pkg assets assets/...

bindata-deps:
	go get github.com/jteeuwen/go-bindata/...

fmt:
	go fmt ./...

deps: bindata-deps
	#FIXME cleanup this
	go get github.com/mailhog/http
	go get github.com/mailhog/MailHog/config
	go get github.com/ian-kent/go-log/log
	go get github.com/ian-kent/envconf
	go get github.com/ian-kent/goose
	go get github.com/ian-kent/linkio
	go get labix.org/v2/mgo

test-deps:
	go get github.com/smartystreets/goconvey

.PNONY: all ui bindata bindata-deps fmt deps test-deps
