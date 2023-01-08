How to Generate the Proto

protoc -Igreet/proto --go_out=. --go_opt=module=github.com/behroozz/grpc-microservice --go-grpc_out=. --go-grpc_opt=module=github.com/behroozz/grpc-microservice greet/proto/dummy.proto
