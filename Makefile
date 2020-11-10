pb:
	@protoc --go_out=. --go-grpc_out=. \
         api/protocols/connectivity.proto