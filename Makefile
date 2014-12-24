DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: deps bindata fmt ui

ui:
	go install .

bindata:
	go-bindata -o assets/assets.go -pkg assets assets/...

fmt:
	go fmt ./...

deps:
	#FIXME cleanup this
	go get github.com/ian-kent/gotcha/gotcha
	go get github.com/ian-kent/go-log/log
	go get github.com/ian-kent/envconf
	go get github.com/ian-kent/goose
	go get github.com/ian-kent/linkio
	go get github.com/jteeuwen/go-bindata/...
	go get labix.org/v2/mgo
	# added to fix travis issues
	go get code.google.com/p/go-uuid/uuid
	go get code.google.com/p/go.crypto/bcrypt

test-deps:
	go get github.com/smartystreets/goconvey

.PNONY: all ui bindata fmt deps test-deps
