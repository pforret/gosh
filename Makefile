BINARY_NAME=gosh
BINARY_DIR=bin
FLAGS=-ldflags "-s -w"

build:
	GOARCH=amd64 GOOS=darwin  go build ${FLAGS} -o ${BINARY_DIR}/${BINARY_NAME}-mac main.go
	GOARCH=amd64 GOOS=linux   go build ${FLAGS} -o ${BINARY_DIR}/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build ${FLAGS} -o ${BINARY_DIR}/${BINARY_NAME}-win main.go
	ls -rtl --color=auto ${BINARY_DIR}

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_DIR}/${BINARY_NAME}-*
