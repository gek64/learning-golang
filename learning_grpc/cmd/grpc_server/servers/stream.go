package servers

import (
	"errors"
	"fmt"
	"io"
	"learning_grpc/pkg/grpc/chat"
	"strings"
)

// ChatServerStream 服务器流服务,服务器可以多次发送信息
func (s *Server) ChatServerStream(req *chat.ChatReq, stream chat.Chat_ChatServerStreamServer) (err error) {
	// 发送5次消息
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("你已经发送的信息是：%s,本次回复是第%d次", req.GetMsg(), i)
		// 调用发送方法来发送msg给客户端
		err = stream.Send(&chat.ChatResp{Msg: msg})
		if err != nil {
			return err
		}
	}
	// 使用return语句来结束流
	return nil
}

// ChatClientStream 客户端流服务,客户端可以多次发送信息,服务端收完信息后使用 stream.SendAndClose() 发送反馈给客户端
func (s *Server) ChatClientStream(stream chat.Chat_ChatClientStreamServer) (err error) {
	// 用于接受汇总所有的客户端传送来的信息
	var messages []string

	// 接受所有的客户端传送来的信息
	for {
		// stream.Recv() 每次接收的都是一个 chatReq
		chatReq, err := stream.Recv()
		if err != nil {
			// 当 err 为 EOF 的时候表示已经接收完成客户端传来的所有信息
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		// 存储接收的所有信息
		messages = append(messages, chatReq.GetMsg())
	}

	// 所有接收的信息合并为一条用于返回到客户端的信息
	respMsg := strings.Join(messages, ",")

	// 发送流的结尾的反馈信息到客户端
	return stream.SendAndClose(&chat.ChatResp{Msg: respMsg})
}

// ChatBiStreams 双向流服务,服务端、客户端都可以多次发送信息,发送完成均需要发送结束反馈给对方
func (s *Server) ChatBiStreams(stream chat.Chat_ChatBiStreamsServer) (err error) {
	for {
		// 接收信息
		req, err := stream.Recv()
		if err != nil {
			// 客户端发送一条信息,服务端回复一条信息,客户端不发送了就终止
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		// 组合服务端返回给客户端的信息
		respMessage := fmt.Sprintf("你好收到的信息是：%s", req.Msg)

		// 发送信息
		err = stream.Send(&chat.ChatResp{Msg: respMessage})
		if err != nil {
			return err
		}
	}
}
