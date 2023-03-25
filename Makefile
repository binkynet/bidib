all:
	go build ./...
	mkdir -p build
	go build -o build/test github.com/binkynet/bidib/test
	go build -o build/bidibWizard github.com/binkynet/bidib/cmd/bidibWizard