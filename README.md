# Blockchain

Simple gRPC server and client to create block and get blockchain

#### Get deps

- export GO_PATH=~/go
- export PATH=$PATH:/$GO_PATH/bin
- go get -u google.golang.org/grpc
- go get -u github.com/golang/protobuf/protoc-gen-go

#### Start server

- go run server/main.go

#### Start server

- go run client/main.go
