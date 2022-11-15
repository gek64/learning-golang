package main

import (
	"fmt"
	"learning_grpc/pkg/grpc/chat"
)

// 实现grpc服务
type serverStream struct {
	chat.UnimplementedChatServer
}

func (s *serverStream) ChatServerStream(req *chat.ChatReq, stream chat.Chat_ChatServerStreamServer) (err error) {
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
