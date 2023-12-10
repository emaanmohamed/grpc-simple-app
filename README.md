Needed dependencies go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 To generate the needed golang files from the proto files. protoc -I./protofiles/greetpb
--go_out=./server
--go-grpc_out=require_unimplemented_servers=false:./server
--go_out=./client
--go-grpc_out=require_unimplemented_servers=false:./client
protofiles/greetpb/greet.proto
