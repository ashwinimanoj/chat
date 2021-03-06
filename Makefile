MAIN=./cmd/main.go

build:
	go build -o bin/chat ${MAIN}

run:
	./bin/chat serve

start:	build run

