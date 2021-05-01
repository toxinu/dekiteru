binaries:
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build -o build/dekiteru-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o build/dekiteru-linux-arm64
	GOOS=darwin GOARCH=amd64 go build -o build/dekiteru-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o build/dekiteru-windows-amd64
