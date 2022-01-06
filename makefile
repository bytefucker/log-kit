
all: clean  build

clean:
	go clean -i ./...
	rm -rf ${GOPATH}/bin/log-kit

build:
	go build  -v -o ${GOPATH}/bin/log-kit .

build-linux:
	GOOS=linux GOARCH=amd64 go build  -v -o ${GOPATH}/bin/log-kit-linux .
