# grpc

- https://www.liwenzhou.com/posts/Go/gRPC/
- https://zenn.dev/hsaki/books/golang-grpc-starting/viewer
- https://connect.build/docs/introduction/
- https://segmentfault.com/a/1190000040917752

## 环境安装

```shell
# protoc
## https://github.com/protocolbuffers/protobuf/releases
## 下载对应版本protoc并添加到环境变量中
# go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 例子中使用

### grpc服务与结构体代码生成

```shell
mkdir -p product && mv product.proto product
protoc --go_out="." --go-grpc_out="." product/product.proto

mkdir -p user && mv user.proto user
protoc --go_out="." --go-grpc_out="." user/user.proto
```

### 添加grpc依赖

```shell
go get -u google.golang.org/grpc
```

