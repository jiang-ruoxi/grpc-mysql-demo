package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-mysql-demo/proto"
	"grpc-mysql-demo/service"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Printf("net.listen tcp connect failed, err::%s\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	//注册实现的服务实例
	var goodsService service.GoodsService
	goodsService = service.NewGoodsService()
	proto.RegisterGoodsServiceServer(grpcServer, goodsService)

	//启动gRPC服务端
	fmt.Println("gRPC is running...")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("gRPC server err:%s\n", err)
	}
}
