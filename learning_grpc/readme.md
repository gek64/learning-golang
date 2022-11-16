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

### `grpc`服务与结构体代码生成

```shell
protoc --go_out="./pkg/grpc" --go-grpc_out="./pkg/grpc" api/product.proto
protoc --go_out="./pkg/grpc" --go-grpc_out="./pkg/grpc" api/user.proto
protoc --go_out="./pkg/grpc" --go-grpc_out="./pkg/grpc" api/chat.proto
```

### 添加`grpc`依赖

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

### 使用`grpcurl`

- https://github.com/fullstorydev/grpcurl/releases
- 需要在`grpc`服务中注册`reflection`服务

```shell
# 查看所有注册的服务
grpcurl -plaintext localhost:8080 list
# 查看注册的服务下的所有方法函数
grpcurl -plaintext localhost:8080 list product.Product
# 发送请求(windows环境下不可用)
grpcurl -plaintext -d '{"msg": "哈哈"}' localhost:8080 chat.Chat.ChatServerStream
# 发送多个请求
grpcurl -plaintext -d '{"msg": "哈哈"}{"msg": "你好"}' localhost:8080 chat.Chat.ChatServerStream
```