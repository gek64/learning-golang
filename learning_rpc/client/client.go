package main

import "net/rpc"

func StartRPCClient() (reply []byte, err error) {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	err = client.Call("ProductService.GetSKU", "客户端远程调用输入id", &reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
