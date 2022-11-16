package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"learning_grpc/pkg/grpc/chat"
	"learning_grpc/pkg/grpc/product"
	"learning_grpc/pkg/grpc/user"
	"log"
	"time"
)

func main() {
	fmt.Println("一对一 rpc")
	err := startRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("服务端流 rpc")
	err = startServerStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("客户端流 rpc")
	err = startClientStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("双向流 rpc")
	err = startBiStreamsRpc()
	if err != nil {
		log.Panicln(err)
	}
}

// 建立网络连接
func getNetConn() (conn *grpc.ClientConn, err error) {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	return grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// 一对一 rpc
func startRpc() (err error) {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := getNetConn()
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(conn)

	// 新建上下文,3s超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 生成客户端
	productClient := product.NewProductClient(conn)
	userClient := user.NewUserClient(conn)

	// 客户端远程调用函数 GetProduct
	productResp, err := productClient.GetProduct(ctx, &product.ProductReq{Id: "测试ID"})
	if err != nil {
		return err
	}
	fmt.Println(productResp)

	// 客户端远程调用函数 GetUser
	userResp, err := userClient.GetUser(ctx, &user.UserReq{Id: "测试ID"})
	if err != nil {
		return err
	}
	fmt.Println(userResp)

	return nil
}

// 服务端流 rpc
func startServerStreamRpc() (err error) {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := getNetConn()
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(conn)

	// 新建上下文,3s超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 生成客户端
	chatClient := chat.NewChatClient(conn)

	// 客户端远程调用服务器流式服务下的 ChatServerStream
	// 返回的 chatRespStream 是一个流客户端,需要使用 chatRespStream.Recv() 来接受服务传来的信息
	chatRespStream, err := chatClient.ChatServerStream(ctx, &chat.ChatReq{Msg: "你好"})
	if err != nil {
		log.Fatal(err)
	}
	// 通过无限for循环接受所有的服务传来的信息
	for {
		// 每次接受信息都是一个 chatResp
		chatResp, err := chatRespStream.Recv()
		if err != nil {
			// 当 err 为 EOF 的时候表示已经接收完成服务传来的所有信息
			if errors.Is(err, io.EOF) {
				break
			}
			// 不为 EOF 则可能是发生了错误,记录报错
			log.Panicln(err)
		}
		fmt.Println(chatResp)
	}

	return nil
}

// 客户端流 rpc
func startClientStreamRpc() (err error) {
	var sendMessages = []string{"哈哈", "你好", "这是多个客户端发送的信息"}

	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := getNetConn()
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(conn)

	// 新建上下文,3s超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 生成客户端
	chatClient := chat.NewChatClient(conn)

	// 客户端远程调用客户端流式服务下的 ChatClientStream
	// 返回的是流
	stream, err := chatClient.ChatClientStream(ctx)
	if err != nil {
		return err
	}

	// 客户端通过流发送多个信息
	for _, message := range sendMessages {
		err := stream.Send(&chat.ChatReq{Msg: message})
		if err != nil {
			return err
		}
	}
	// 发送完成信息后,接收服务端通过流传递来的消息,并关闭流
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	// 服务端通过流传递来的消息
	fmt.Println(resp)

	return nil
}

// 双向流 rpc
func startBiStreamsRpc() (err error) {
	var sendMessages = []string{"哈哈", "你好", "这是多个客户端发送的信息"}

	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := getNetConn()
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(conn)

	// 新建上下文,3s超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 生成客户端
	chatClient := chat.NewChatClient(conn)

	streams, err := chatClient.ChatBiStreams(ctx)
	if err != nil {
		return err
	}

	for _, message := range sendMessages {
		err := streams.Send(&chat.ChatReq{Msg: message})
		if err != nil {
			return err
		}
	}
	err = streams.CloseSend()
	if err != nil {
		return err
	}

	for {
		resp, err := streams.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		fmt.Println(resp.GetMsg())
	}
}
