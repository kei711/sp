.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/sp
	# GOOS=linux GOARCH=386 go build -o ./bin/linux_386/sp

	# GOOS=darwin GOARCH=386 go build -o ./bin/darwin_386/sp
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/sp
