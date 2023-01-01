BINARY_NAME=gosh
BINARY_DIR=bin

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_DIR}/${BINARY_NAME}-mac main.go
	GOARCH=amd64 GOOS=linux  go build -o ${BINARY_DIR}/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_DIR}/${BINARY_NAME}-win main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_DIR}/${BINARY_NAME}-*
