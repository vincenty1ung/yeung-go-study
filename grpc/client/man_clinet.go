package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/vincenty1ung/yeung-go-study/grpc/pb"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)

	for {
		// 初始化etcd客户端
		var addr = "localhost:2379"
		cli, _ := clientv3.New(
			clientv3.Config{
				Endpoints:   []string{addr},
				DialTimeout: 5 * time.Second,
			},
		)

		// 新建builder，etcd官方实现的Builder对象
		r, _ := resolver.NewBuilder(cli)
		// 向grpc注册builder，这样Dial时，就可以按照Scheme查找到此Builder，
		// grpcResolver.Register(r) 多余以下的WithResolvers做了代替
		conn, err := grpc.Dial(
			r.Scheme()+"://"+addr+"/project_test", grpc.WithTransportCredentials(insecure.NewCredentials()),
			// grpc.WithBlock(),
			// grpc.WithAuthority(""),
			grpc.WithResolvers(r),
		)
		if err != nil {
			log.Fatal("连接 gPRC 服务失败,", err)
		}

		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		manServiceClient := pb.NewManServiceClient(conn)
		// 发送请求，调用 MyTest 接口
		response, err := manServiceClient.GetMan(
			context.Background(), &pb.GetManRequest{Id: 1, Bytes: []byte{1, 2, 4}, ManMap: nil},
		)
		if err != nil {
			log.Fatal("发送请求失败，原因是:", err)
		}
		log.Println(response)

		// line, _, _ := reader.ReadLine()
		// log.Println(string(line))
		time.Sleep(time.Millisecond * 100)
	}
}
