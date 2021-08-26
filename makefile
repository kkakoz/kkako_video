
.PHONY: protoc
protoc:
	@for dir in api/*; \
	do \
	  for v in $$dir/*; \
	  	do \
  	  		for proto in $$v/*; \
  	  		do \
			  protoc --proto_path=. -I ./third_party \
				--go_out . --go_opt paths=source_relative \
				--go-grpc_out . --go-grpc_opt paths=source_relative \
				--go-gin_out . --go-gin_opt=paths=source_relative \
				$$proto; \
		   done ;\
  		done ;\
  	done


