#

## gencode

```
生成 proto
goctl rpc protoc api/protobuf/user/v1/*.proto --go_out=. --go-grpc_out=. --zrpc_out=app/user/cmd/rpc
生成 model
goctl model mysql ddl -s deploy/sql/user/000001_create_users_table.up.sql -d app/user
```
