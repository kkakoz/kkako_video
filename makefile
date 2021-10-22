
.PHONY: protoc
protoc:
		protoc -I. -I ./third_party -I ./api \
			--go_out . --go_opt paths=source_relative \
			--go-grpc_out . --go-grpc_opt paths=source_relative \
			--grpc-gateway_out . \
			--grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt paths=source_relative \
			--grpc-gateway_opt generate_unbound_methods=true \
			 ./api/user/v1/*.proto ./api/video/v1/*.proto ./api/comment/v1/*.proto
##		--go-gin_out . --go-gin_opt=paths=source_relative \


.PHONY: mock
mock:
	mockgen -source=api/user/v1/user_grpc.pb.go -destination=api/user/v1/user_grpc_mock.go -package=v1


