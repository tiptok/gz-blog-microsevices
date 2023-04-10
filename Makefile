redeployed-at:=$(shell date +%s)

.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/google/wire/cmd/wire@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golang/mock/mockgen@latest
	@echo "Installing protoc-gen-validate (PGV) can currently only be done from source. See: https://github.com/envoyproxy/protoc-gen-validate#installation"

.PHONY: protoc
protoc:
	for file in $$(find api -name '*.proto'); do \
		protoc \
		-I $$(dirname $$file) \
		-I ./third_party \
		--go_out=:$$(dirname $$file) --go_opt=paths=source_relative \
		--go-grpc_out=:$$(dirname $$file) --go-grpc_opt=paths=source_relative \
		--validate_out="lang=go:$$(dirname $$file)" --validate_opt=paths=source_relative \
		--grpc-gateway_out=:$$(dirname $$file) --grpc-gateway_opt=paths=source_relative \
		$$file; \
	done

.PHONY: zrpc
zrpc:
	for file in $$(find api -name '*.proto'); do \
		goctl rpc protoc \
		-I $$(dirname $$file) \
		-I ./third_party \
		--go_out=$$(dirname $$file) --go_opt=paths=source_relative \
		--go-grpc_out=$$(dirname $$file) --go-grpc_opt=paths=source_relative \
		--zrpc_out="app/$$(basename $$file | cut -f 1 -d '.')/cmd/rpc" \
		-v \
		$$file; \
	done

.PHONY: rpc-gen
rpc-gen: protoc zrpc


.PHONY: model-gen
model-gen:
	goctl model mysql ddl -s deploy/sql/comment/000001_create_comments_table.up.sql -d app/comment

.PHONY: wire
wire:
	wire ./...


.PHONY: test
test:
	go test -cover -race -covermode=atomic -coverprofile=coverage.txt ./...


.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations/user -database "mysql://root:@tcp(localhost:3306)/users" -verbose up
	migrate -path ./migrations/post -database "mysql://root:@tcp(localhost:3306)/posts" -verbose up
	migrate -path ./migrations/comment -database "mysql://root:@tcp(localhost:3306)/comments" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations/user -database "mysql://root:@tcp(localhost:3306)/users" -verbose down -all
	migrate -path ./migrations/post -database "mysql://root:@tcp(localhost:3306)/posts" -verbose down -all
	migrate -path ./migrations/comment -database "mysql://root:@tcp(localhost:3306)/comments" -verbose down -all

.PHONY: migrate-refresh
migrate-refresh: migrate-down migrate-up

.PHONY: auth-server
auth-server:
	go run ./app/auth/cmd/rpc/auth.go -f ./app/auth/cmd/rpc/etc/auth.yaml

.PHONY: user-server
user-server:
	go run ./app/user/cmd/rpc/user.go -f ./app/user/cmd/rpc/etc/user.yaml

.PHONY: post-server
post-server:
	go run ./app/post/cmd/rpc/post.go -f ./app/post/cmd/rpc/etc/post.yaml

.PHONY: comment-server
comment-server:
	go run ./app/comment/cmd/rpc/comment.go -f ./app/comment/cmd/rpc/etc/comment.yaml

.PHONY: blog-server
blog-server:
	go run ./app/blog/cmd/rpc/blog.go -f ./app/blog/cmd/rpc/etc/blog.yaml

.PHONY: win-up
win-up:
	start make user-server
	start make auth-server
	start make post-server
	start make comment-server
	start make blog-server

.PHONY: docker-build
docker-build:
	docker build -f deploy/build/blog/Dockerfile -t tiptok/gz-blog:1.0.0 .
	docker build -f deploy/build/user/Dockerfile -t tiptok/gz-blog-user:1.0.0 .
#    docker build -f deploy/build/auth/Dockerfile -t tiptok/gz-blog-auth:1.0.0 .
#    docker build -f deploy/build/comment/Dockerfile -t tiptok/gz-blog-comment:1.0.0 .
#    docker build -f deploy/build/post/Dockerfile -t tiptok/gz-blog-post:1.0.0 .

.PHONY: kube-deploy
kube-deploy:
	kubectl apply -f ./deployments/
	kubectl apply -f ./deployments/dtm/
	kubectl apply -f ./deployments/blog/
	kubectl apply -f ./deployments/user/
	kubectl apply -f ./deployments/post/
	kubectl apply -f ./deployments/auth/
	kubectl apply -f ./deployments/comment/
	kubectl apply -f ./deployments/addons/

.PHONY: kube-delete
kube-delete:
	kubectl delete -f ./deployments/
	kubectl delete -f ./deployments/dtm/
	kubectl delete -f ./deployments/blog/
	kubectl delete -f ./deployments/user/
	kubectl delete -f ./deployments/post/
	kubectl delete -f ./deployments/auth/
	kubectl delete -f ./deployments/comment/
	kubectl delete -f ./deployments/addons/

.PHONY: kube-redeploy
kube-redeploy:
	@echo "redeployed at ${redeployed-at}"
	kubectl patch deployment blog-server -p '{"spec": {"template": {"metadata": {"annotations": {"redeployed-at": "'${redeployed-at}'" }}}}}'
	kubectl patch deployment user-server -p '{"spec": {"template": {"metadata": {"annotations": {"redeployed-at": "'${redeployed-at}'" }}}}}'
	kubectl patch deployment auth-server -p '{"spec": {"template": {"metadata": {"annotations": {"redeployed-at": "'${redeployed-at}'" }}}}}'
	kubectl patch deployment post-server -p '{"spec": {"template": {"metadata": {"annotations": {"redeployed-at": "'${redeployed-at}'" }}}}}'
	kubectl patch deployment comment-server -p '{"spec": {"template": {"metadata": {"annotations": {"redeployed-at": "'${redeployed-at}'" }}}}}'