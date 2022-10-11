# All the commands below should be execute at a sepecific folder with sepecific args.
# Forgive my laziness. Cause the identifying of path spreads every where. 

#
# API Related
# apps/service/cmd/api/def
# goctl api -o xxx.api
# goctl api go -api xxx.api -dir ../

# RPC Related
# apps/service/cmd/rpc/pb
# goctl rpc -o xxx.proto
# goctl rpc protoc xxx.proto --go_out=. --go-grpc_out=. --zrpc_out=../

#
# Model Related
# apps/service/model
# the target folder path, instead of current relative path, 
# is not required unless there are multiple models 
# and you wanna put them into seperate folders.
# goctl model mysql datasource --url="user:pwd@tcp(host:port)/db" -table="table" -c -dir=./model_dir

#
# Template Related
# deploy/goctl
# goctl template init --home=./template