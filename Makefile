.PHONY: all clean

all: linux windows osx


osx:
	mkdir -p ./build
	env GOOS=darwin GOARCH=amd64 go build -o ./build/netuitive-event-handler-osx main.go

linux:
	mkdir -p ./build
	env GOOS=linux GOARCH=amd64 go build -o ./build/netuitive-event-handler-linux main.go

windows:
	mkdir -p ./build
	env GOOS=windows GOARCH=amd64 go build -o ./build/netuitive-event-handler-windows main.go

clean:
	rm -rf ./build

setup:
	glide install