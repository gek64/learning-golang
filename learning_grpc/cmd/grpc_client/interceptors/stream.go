package interceptors

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"io"
	"log"
)

func MyStreamClientInterceptor1(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Println("[pre] my stream client interceptor 1", method)

	stream, err := streamer(ctx, desc, cc, method, opts...)
	return &MyClientStreamWrapper1{stream}, err
}

type MyClientStreamWrapper1 struct {
	grpc.ClientStream
}

func (s *MyClientStreamWrapper1) SendMsg(m interface{}) error {
	log.Println("[pre message] my stream client interceptor 1: ", m)

	return s.ClientStream.SendMsg(m)
}

func (s *MyClientStreamWrapper1) RecvMsg(m interface{}) error {
	err := s.ClientStream.RecvMsg(m)

	if !errors.Is(err, io.EOF) {
		log.Println("[post message] my stream client interceptor 1: ", m)
	}
	return err
}

func (s *MyClientStreamWrapper1) CloseSend() error {
	err := s.ClientStream.CloseSend()

	log.Println("[post] my stream client interceptor 1")
	return err
}
