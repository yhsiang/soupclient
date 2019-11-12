all: build-linux build-darwin-amd64 build-windows-amd64 docker
test:
	go test ./...
build-linux:
	GOOS=linux GOARCH=amd64 go build -o go-soupclient .
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o go-soupclient-darwin-amd64 .
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o go-soupclient-window-amd64.exe .
docker:
	docker build -t soupclient:0.0.2 .
clean:
	rm go-soupclient go-soupclient-darwin-amd64 go-soupclient-window-amd64.exe