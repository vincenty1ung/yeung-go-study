package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vincenty1ung/yeung-go-study/grpc/etcd"
	"github.com/vincenty1ung/yeung-go-study/grpc/pb"
	"github.com/vincenty1ung/yeung-go-study/grpc/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "p", "8080", "设置端口")
	// flag.StringVar(&port, "p", "", "设置端口")
	flag.Parse()
}

func main() {

	// 创建 Tcp 连接
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 创建gRPC服务
	grpcServer := grpc.NewServer()

	// Tester 注册服务实现者
	ser := new(server.ManServiceServerImpl)
	pb.RegisterManServiceServer(grpcServer, ser)

	// 在 gRPC 服务上注册反射服务
	// func Register(s *grpc.Server)
	reflection.Register(grpcServer)
	err = etcd.Register([]string{"localhost:2379"}, "project_test/", "localhost:"+port)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
