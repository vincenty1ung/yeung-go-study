package etcd

import (
	"context"
	"strconv"
	"time"

	client "go.etcd.io/etcd/client/v3"
)

var cli *client.Client
var interval = 5

// Register 注册服务
// etcdAdders:etcd地址
// serviceName:注册到etcd的服务名称
// serviceAddr:自己grpc业务监听的地址端口
func Register(etcdAddrs []string, serviceName string, serviceAddr string) error {
	// 获取链接
	var err error
	if cli == nil {
		cli, err = client.New(
			client.Config{
				Endpoints:   etcdAddrs,
				DialTimeout: 5 * time.Second,
			},
		)
		if err != nil {
			return err
		}
	}
	// 注册续租
	return register(serviceName, serviceAddr)
}

// etcd服务发现时，底层解析的是一个json串，且包含Addr字段
func getValue(addr string) string {
	return "{\"Addr\":\"" + addr + "\"}"
}

func register(serviceName, serviceAddr string) error {
	// 注册服务
	leaseResp, err := cli.Grant(context.Background(), int64(interval+1))
	if err != nil {
		return err
	}
	fullKey := serviceName + "/" + strconv.Itoa(int(leaseResp.ID))
	_, err = cli.Put(context.Background(), fullKey, getValue(serviceAddr), client.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}
	keepAlive(serviceName, serviceAddr, leaseResp)
	return nil
}

// 异步续约
func keepAlive(name string, addr string, leaseResp *client.LeaseGrantResponse) {
	// 永久续约，续约成功后，etcd客户端和服务器会保持通讯，通讯成功会写数据到返回的通道中
	// 停止进程后，服务器链接不上客户端，相应key租约到期会被服务器自动删除
	c, err := cli.KeepAlive(cli.Ctx(), leaseResp.ID)
	go func() {
		if err == nil {
			defer func(cli *client.Client, ctx context.Context, id client.LeaseID) {
				_, err := cli.Revoke(ctx, id)
				if err != nil {

				}
			}(cli, cli.Ctx(), leaseResp.ID)
			for {
				select {
				case _, ok := <-c:
					if !ok { // 续约失败
						_, err := cli.Revoke(cli.Ctx(), leaseResp.ID)
						if err != nil {
							return
						}
						err = register(name, addr)
						if err != nil {
							return
						}
						return
					}
				}
			}

		}
	}()
}
