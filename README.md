# prerequisite: gRPC code generator

## install protobuf and the related go extension

```
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## install protoc-gen-go-grpc

```
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## generate pb & gRPC codes

```
$ protoc \
    --go_out=./transition --go_opt=paths=source_relative \
    --go-grpc_out=./transition --go-grpc_opt=paths=source_relative \
    transition.proto
```

# install go protobuf dependencies

```
$ go get google.golang.org/grpc
$ go get google.golang.org/protobuf
```
