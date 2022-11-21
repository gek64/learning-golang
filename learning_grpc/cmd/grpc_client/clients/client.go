package clients

import (
	"context"
	"errors"
	"fmt"
	// 需要匿名导入来避免 [proto: not found] 错误,即使不使用这个包的内容
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"io"
	"learning_grpc/pkg/grpc/chat"
	"learning_grpc/pkg/grpc/errorCode"
	"learning_grpc/pkg/grpc/product"
	"learning_grpc/pkg/grpc/user"
	"log"
	"time"
)

// GetNetConn 建立网络连接
func GetNetConn() (conn *grpc.ClientConn, err error) {
	// 建立连接
	return grpc.Dial("localhost:8080",
		// 不使用 tls/ssl 必选
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 等待连接建立成功后再执行接下来的操作,同步执行(可选,默认后台异步建立连接)
		grpc.WithBlock(),

		// 自定义一元拦截器(可选)
		//grpc.WithUnaryInterceptor(interceptor.MyUnaryClientInterceptor1),
		// 自定义流拦截器(可选)
		//grpc.WithStreamInterceptor(interceptor.MyStreamClientInterceptor1),

		// 多个拦截器链式启动(可选)
		grpc.WithChainUnaryInterceptor(),
		grpc.WithChainStreamInterceptor(),
	)
}

// StartRpc 一对一 rpc
func StartRpc() (err error) {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := GetNetConn()
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
	fmt.Println(userResp, userResp.GetBirthday().AsTime())

	return nil
}

// StartServerStreamRpc 服务端流 rpc
func StartServerStreamRpc() (err error) {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := GetNetConn()
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

// StartClientStreamRpc 客户端流 rpc
func StartClientStreamRpc() (err error) {
	var sendMessages = []string{"哈哈", "你好", "这是多个客户端发送的信息"}

	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := GetNetConn()
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

// StartBiStreamsRpc 双向流 rpc
func StartBiStreamsRpc() (err error) {
	var sendMessages = []string{"哈哈", "你好", "这是客户端发送的多个信息"}

	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := GetNetConn()
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

	// 调用客户端下 ChatBiStreams 服务
	streams, err := chatClient.ChatBiStreams(ctx)
	if err != nil {
		return err
	}

	// 发送所有信息
	for i, message := range sendMessages {
		// 使用 streams.Send() 来发送信息到服务端
		err := streams.Send(&chat.ChatReq{Msg: message})
		if err != nil {
			return err
		}

		// 所有消息发送完成
		if i+1 == len(sendMessages) {
			// 终止发送(发送 io.EOF)
			err = streams.CloseSend()
			if err != nil {
				return err
			}
		}

		// 接收服务端反馈的信息
		resp, err := streams.Recv()
		if err != nil {
			// 读取到服务端发送的 io.EOF 表示接收完成所有服务端的反馈信息
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		fmt.Println(resp.GetMsg())
	}

	return nil
}

// TestRpcErrorCode 错误测试
func TestRpcErrorCode() (err error) {
	// 建立连接,配置选项忽略传输层凭据(tls/ssl)
	conn, err := GetNetConn()
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
	errorClient := errorCode.NewErrorClient(conn)

	// 客户端远程调用函数,获取自定义的错误
	_, err = errorClient.ErrorServer(ctx, &errorCode.ErrorReq{Msg: ""})
	// 提取错误
	s, _ := status.FromError(err)
	fmt.Printf("错误代码：%d\n", s.Code())
	fmt.Printf("错误信息：%s\n", s.Message())
	fmt.Printf("错误详情：%s\n", s.Details())

	return nil
}
