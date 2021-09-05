
.PHONY: protoc
protoc:
		protoc -I. -I ./third_party  \
			--go_out . --go_opt paths=source_relative \
			--go-grpc_out . --go-grpc_opt paths=source_relative \
			 ./api/user/v1/*.proto ./api/video/v1/*.proto ./api/comment/v1/*.proto
##		--go-gin_out . --go-gin_opt=paths=source_relative \


.PHONY: baseProto
baseProto:
		protoc -I. -I ./third_party  \
			--go_out=plugins=grpc:./api/user/v1/ \
			 ./api/user/v1/*.proto
#	@for f in api/*; \
#	do \
#	  for v in $$f/*; \
#	  	do \
#  	  		for proto in $$v/*; \
#  	  		do \
#			  protoc --proto_path=. -I ./third_party \
#				--go_out . --go_opt paths=source_relative \
#				--go-grpc_out . --go-grpc_opt paths=source_relative \
#				--go-gin_out . --go-gin_opt=paths=source_relative \
#				$$proto; \
#		   done ;\
#  		done ;\
#  	done


