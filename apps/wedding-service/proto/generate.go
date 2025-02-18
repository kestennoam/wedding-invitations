package proto

//go:generate echo "Hello, World!"
//go:generate protoc --go_out=. --go-grpc_out=. --proto_path=. wedding.proto
//go:generate echo "Hello, World!"
