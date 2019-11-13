# Soupbintcp client

## requirement

- [x] The client should be compiled to a binary that can run in alpine linux.

- [x] The binary can use arguments (or flags) to specify which server will be connected
- [x] The client should exit when it receives an OS interrupt signal
- [x] Implement error handling logic and don't just panic
- [x] Write some unit tests
- [x] Do the following actions (in order) after the client connects to the server

# Development

`$ go get -u https://github.com/yhsiang/soupclient`

`$ go run main.go`

# Build

Make sure your golang with cross compile support and choose your target platform.

* Linux   `$ make build-linux`
* MacOS   `$ make build-darwin-amd64`
* Windows `$ make build-windows-amd64`

# Docker

You need build executable first then use the below command.

`$ make docker`

# Test

Please run test first to make sure all's fine in your environment.

`$ make test`