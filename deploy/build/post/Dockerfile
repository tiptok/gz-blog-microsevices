FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
# RUN go mod download
COPY . .
COPY deploy/config/post /app/etc
RUN go mod tidy
RUN go build -ldflags="-s -w" -o /app/post app/post/cmd/rpc/post.go


FROM ubuntu:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/post /app/post
COPY --from=builder /app/etc /app/etc

EXPOSE 8080

CMD ["./post", "-f", "etc/post.yaml"]
