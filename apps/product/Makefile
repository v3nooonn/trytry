#gen-pb:
#	57ctl rpc -o ./pb/$(pb).proto --remote="https://github.com/v3nooonn/trytry-template.git"
#gen-code:
#	57ctl rpc protoc ./pb/$(pb).proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
#gen-model:
#	57ctl model pg datasource --url="postgres://pgadmin:pgadmin@127.0.0.1:5434/trytry?sslmode=disable" --schema=$(schema) -table=$(table) -c -dir=./internal/model --remote="https://github.com/v3nooonn/trytry-template"

proto:
	#goctl rpc -o ./pb/$(pb).proto
	goctl rpc -o ./pb/product.proto

model:
	goctl model mysql datasource --url="root:root@tcp(127.0.0.1:3306)/trytry" -table=$(table) -c -dir=./internal/model

code:
	#goctl rpc protoc ./pb/$(pb).proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
	goctl rpc protoc ./pb/product.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
