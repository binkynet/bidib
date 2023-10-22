all: dependencies
	go build ./...
	mkdir -p build
	go build -o build/test github.com/binkynet/bidib/test
	go build -o build/bidibWizard github.com/binkynet/bidib/cmd/bidibWizard
	gox \
		-osarch="darwin/arm64 linux/amd64 linux/arm linux/arm64" \
		-output="build/{{.OS}}/{{.Arch}}/bidibWizard" \
		-tags="netgo" \
		github.com/binkynet/bidib/cmd/bidibWizard


dependencies:
	go install golang.org/x/tools/cmd/stringer

deploy:
	scp build/linux/arm/bidibWizard pi@192.168.77.1:/home/pi/