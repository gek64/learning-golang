package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_client/product"
	"grpc_client/user"
	"log"
	"time"
)

func main() {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	// 建立客户端连接
	productClient := product.NewProductClient(conn)
	userClient := user.NewUserClient(conn)

	// 新建上下文,3s超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 客户端远程调用函数 GetProduct
	productResp, err := productClient.GetProduct(ctx, &product.ProductReq{Id: "测试ID"})
	if err != nil {
		log.Fatal(err)
	}
	// 客户端远程调用函数 GetUser
	userResp, err := userClient.GetUser(ctx, &user.UserReq{Id: "测试ID"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n%v\n%v\n", productResp, userResp, userResp.Birthday.AsTime())
}
