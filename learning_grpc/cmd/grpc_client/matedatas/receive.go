package matedatas

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learning_grpc/cmd/grpc_client/clients"
	"learning_grpc/pkg/grpc/login"
	"log"
)

func Logout() (err error) {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := clients.GetNetConn()
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(conn)
	// 创建客户端连接
	client := login.NewLoginClient(conn)

	// 封装元数据内容
	md := metadata.New(map[string]string{
		"type": "logout",
		"from": "client",
	})
	// 元数据附加到上下文中
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 请求服务,并附带元数据
	var header, trailer metadata.MD
	resp, err := client.Logout(ctx, &login.LoginReq{Msg: "上传元数据"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return err
	}

	fmt.Println(resp.GetMsg())
	fmt.Printf("header:%v\n", header)
	fmt.Printf("trailer:%v\n", trailer)

	return nil
}
