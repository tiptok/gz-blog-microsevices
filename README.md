#

## 安装工具

```
GOCTL_VERSION=1.3.8
PROTOC_VERSION=22.2



go install github.com/envoyproxy/protoc-gen-validate@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
```

## gencode

```
生成 proto
goctl rpc protoc api/protobuf/user/v1/*.proto --go_out=. --go-grpc_out=. --zrpc_out=app/user/cmd/rpc
生成 model
goctl model mysql ddl -s deploy/sql/user/000001_create_users_table.up.sql -d app/user

goctl rpc protoc -I ./third_party -I ./api/protobuf  api/protobuf/blog/v1/*.proto  \
    --go_out=. \
    --go-grpc_out=. \
    --zrpc_out=app/blog/cmd/rpc
    
protoc -I ./third_party -I ./api/protobuf  api/protobuf/blog/v1/*.proto  \
    --go_out=. \
    --go-grpc_out=. \
    --validate_out=.,lang=go:. \
    --grpc-gateway_out=.  
```


## test

```
curl -XGET http://127.0.0.1:9999/api.rest.blog.v1/comments
```


## 生成配置
生成Dockerfile
```
goctl docker --go blog.go --port 8080 --version 1.19 --base ubuntu:latest
```