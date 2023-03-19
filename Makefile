all:
	go build ./...
	mkdir -p build
	go build -o build/test github.com/binkynet/bidib/test