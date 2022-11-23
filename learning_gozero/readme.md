# gozero

## 安装
```shell
# 项目中添加 gozero 依赖
go get -u github.com/zeromicro/go-zero

# 安装 goctl 工具
go install github.com/zeromicro/go-zero/tools/goctl
```

## 生成grpc及zrpc文件(用于rpc服务)
```shell
goctl rpc protoc api/user.proto --go_out=pkg/user/types --go-grpc_out=pkg/user/types --zrpc_out=pkg/user
```

## 生成api文件(用于web服务)
```shell
goctl api go -api api/order.api -dir pkg/order
```