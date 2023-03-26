all: dependencies
	go build ./...
	mkdir -p build
	go build -o build/test github.com/binkynet/bidib/test
	go build -o build/bidibWizard github.com/binkynet/bidib/cmd/bidibWizard

dependencies:
	go install golang.org/x/tools/cmd/stringer