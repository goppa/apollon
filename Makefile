all: build

build:
	go build -o ./build ./main.go

run:
	go run ./main.go

test:
	go test -v ./
