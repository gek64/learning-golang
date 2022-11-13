package main

import (
	"encoding/json"
	"net"
	"net/rpc"
)

type ProductService struct {
	Id    string
	Name  string
	SKU   string
	Price float64
}

func (p *ProductService) GetSKU(request string, reply *[]byte) (err error) {
	*reply, err = json.Marshal(ProductService{
		Id:    request,
		Name:  "测试用商品",
		SKU:   "A1000",
		Price: 100,
	})
	return err
}

func StartRPCServer() (server *rpc.Server, err error) {
	// 注册tcp网络监听器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	// rpc服务建立
	server = rpc.NewServer()
	// rpc服务注册
	err = server.RegisterName("ProductService", new(ProductService))
	if err != nil {
		return nil, err
	}
	// rpc服务使用建立的tcp网络监听器(程序会停留在此处接收网络tcp数据包)
	server.Accept(listener)

	return server, nil
}
