package main

import (
	"errors"
	"google.golang.org/grpc"
	"io"
	"log"
)

type myServerStreamWrapper1 struct {
	grpc.ServerStream
}

// RecvMsg 通过重写标准方法插入接收消息之前的预处理
func (s *myServerStreamWrapper1) RecvMsg(m interface{}) error {
	err := s.ServerStream.RecvMsg(m)
	if !errors.Is(err, io.EOF) {
		log.Println("[pre message] 流拦截器1:", m)
	}
	return err
}

// SendMsg 通过重写标准方法插入发送消息之后的后处理
func (s *myServerStreamWrapper1) SendMsg(m interface{}) error {
	log.Println("[post message] 流拦截器1:", m)
	return s.ServerStream.SendMsg(m)
}

func myStreamServerInterceptor1(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// 打开流时的预处理
	log.Println("[pre stream] 流拦截器1:", info.FullMethod)

	// 执行服务函数
	err := handler(srv, &myServerStreamWrapper1{ss})

	// 关闭流的后处理
	log.Println("[post stream] 流拦截器1 运行结束")

	return err
}
