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

```
goctl rpc protoc -I ./third_party -I ./api/protobuf  api/protobuf/auth/v1/*.proto  \
    --go_out=. \
    --go-grpc_out=. \
    --zrpc_out=app/auth/cmd/rpc 
    

goctl rpc protoc -I ./third_party -I ./api/protobuf  api/protobuf/postv1/*.proto  \
    --go_out=./api/protobuf/post \
    --go-grpc_out=./api/protobuf/post \
    --zrpc_out=app/post/cmd/rpc \
    -v
        
protoc -I ./third_party -I ./api/protobuf  api/protobuf/post/v1/*.proto  \
    --go_out=. \
    --go-grpc_out=. \
    --validate_out="lang=go,paths=source_relative:." \
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

生成镜像
```
docker build -f deploy/build/user/Dockerfile -t tiptok/gz-blog-user:1.0.0 .
docker build -f deploy/build/blog/Dockerfile -t tiptok/gz-blog:1.0.0 .
```

## 部署

### 部署redis

```
k apply -k ./deploy/platform/kube/redis
kubectl get all -n redis
k -n redis exec -it redis-66fd8f7cd7-2j8gg -- sh
# redis-cli
127.0.0.1:6379> auth Hello1234!!!!
OK
127.0.0.1:6379> config get requirepass

访问

在同一个 Kubernetes 集群里面的应用, 可以通过 service 连接 redis

service name: redis-svc.redis
port: 6379
```

### 部署mysql
```
k apply -k deploy/platform/kube/mysql/mysql.yaml
```


### 部署etcd
```
k apply -f deploy/platform/kube/etcd/etcd.yaml
k apply -f deploy/platform/kube/etcd/etcd-web-ui.yaml 
查看 etcd v3 web ui
http://106.52.103.187:30081/
etcdctl get user.rpc --prefix
etcdctl  get --prefix ""
```

### 部署服务
```
# 使用k8s负载均衡需要先配置svc account绑定查看endpoints权限
k apply -f deploy/platform/kube/svc-account -n app

# 配置
k apply -f deploy/platform/kube/gz-blog-user-config.yaml -n app
k apply -f deploy/platform/kube/gz-blog-config.yaml -n app

k apply -f deploy/platform/kube/gz-blog-user.yaml -n app
k apply -f deploy/platform/kube/gz-blog.yaml -n app
```