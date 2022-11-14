# grpc

- https://www.liwenzhou.com/posts/Go/gRPC/
- https://zenn.dev/hsaki/books/golang-grpc-starting/viewer
- https://connect.build/docs/introduction/
- https://segmentfault.com/a/1190000040917752
- https://hypc.github.io/2019/08/16/golang-project-structure/

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
protoc --go_out="./pkg/grpc" --go-grpc_out="./pkg/grpc" api/product.proto
protoc --go_out="./pkg/grpc" --go-grpc_out="./pkg/grpc" api/user.proto
```

### 添加grpc依赖

```shell
go get -u google.golang.org/grpc
```

### 运行安装
```shell
# 运行
go run learning_grpc/cmd/grpc_server
go run learning_grpc/cmd/grpc_client
# 安装所有
go install learning_grpc/cmd/...
```