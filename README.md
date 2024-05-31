# grpc-search

This project demonstrates a gRPC server implementation in Go. The project includes protobuf definitions, a server, and a client.

## Project Structure

```
/myproject
  /proto
    service.proto
  /cmd
    server.go
    main.go
    di.go
  go.mod
  go.sum
```

## Prerequisites

Ensure you have the following installed on your system:
- Go (version 1.16 or later)
- Protocol Buffers compiler (`protoc`)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

To install the `protoc` plugins, run:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## GRPC endpoints Definition
```
service TransacUser {
  rpc AddUser (UserModel) returns (APIResp) {}
  rpc GetUser (GetUserRequest) returns (GetUserResponse) {}
}
```

AddUser accepts user Model with schema: 
```
message UserModel {
  int64 ID = 1;
  string Name = 2;
  string Email = 3;
  string City = 4;
  string Phone = 5;
  string Married = 6;
}
```

GetUser can be used as a dynamic search query executor: 
It accepts argument with schema: 
```
message GetUserRequest {
  oneof criteria {
    int64 id = 1;
    IdsRequest idsRequest = 2;
    SearchCriteria search = 3;
  } 
}
```

### Further functional definition can be found under proto/service.proto

```
export PATH=$PATH:$(go env GOPATH)/bin
```
## Makefile Targets

The Makefile provides several targets to streamline the development process. Below are the main targets and their descriptions:
```
make
```
This is the default target. It generates the protobuf files and builds the server binary.

```
make
```
Generate protobuf files:
```
make proto
```
Generates Go files from the .proto files in the proto directory.

```
make proto
make build
```
Builds the server binary.

```
make build
make clean
```
Cleans up the generated files and binaries.

```
make clean
make run
```
Runs the server binary.

```
make run
```

### Steps to Build and Run the gRPC Server
Generate protobuf files and build the project:

```
make
```
Run the server:
```
make run
```
