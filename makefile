setup :
	protoc --go_out=. --go-grpc_out=. --proto_path=. api.proto
	protoc --grpc-gateway_out=. --go_out=. --go-grpc_out=. --proto_path=. api.proto
	protoc --openapiv2_out=./swagger --proto_path=. api.proto

start :
	go run server.go

test :
	go test -cover ./...
