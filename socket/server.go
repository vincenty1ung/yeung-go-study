package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

func server() {
	logx.Info("server: 启动服务端...")
	// net.Listen()
	// ListenAndBind
	tcp, _ := net.Listen("tcp", "127.0.0.1:8099")
	// accept
	defer func(tcp net.Listener) {
		err := tcp.Close()
		if err != nil {
			logx.Error(err)
		}
	}(tcp)

	select {
	case <-chanEXIT:
		logx.Info("server: 服务端停止...")
		return
	default:

	}
	logx.Info("server: acceptTCP")
	go accept(tcp)
}

func accept(listener net.Listener) {
	for {
		// 阻塞在这里等待新的全连接套接字 io多路复用 内核阻塞 g挂起
		acceptTCP, err := listener.Accept()
		if err != nil {
			logx.Error(err)
			return
		}
		logx.Info("server: 阻塞了吗?...上面的select会被阻塞在这里,导致客户端发送结束后服务端会在这里阻塞读取最新的套接字,go底层的accept和read和Write都使用了epoll,linux事件处理器处理调用的就绪io调度")
		go handleConn(acceptTCP)
	}
}

// 读取
func handleConn(c net.Conn) {
	defer func(cx net.Conn) {
		err := cx.Close()
		if err != nil {
			logx.Error(err)
		}
	}(c)

	for {
		bytes := make([]byte, 1024)
		// io多路复用 内核阻塞 g挂起
		rn, err := c.Read(bytes)
		if err != nil {
			logx.Error(err)
			return
		}

		sreq := new(req)
		err = json.Unmarshal(bytes[:rn], sreq)
		if err != nil {
			logx.Error(err)
			return
		}

		logx.Info(fmt.Sprintf("server:收到请求: %v", sreq))

		time.Sleep(time.Second * 2)
		logx.Info("server: 业务操作处理耗时2秒")
		logx.Info("server: 业务操作处理完成")

		sresp := &resp{
			Msg:       fmt.Sprintf("你好啊:%s号客户端.", sreq.Msg[len(sreq.Msg)-1:len(sreq.Msg)]),
			LocalAddr: c.LocalAddr().String(),
		}
		srespmarshal, err := json.Marshal(sresp)
		if err != nil {
			logx.Error(err)
			return
		}
		// io多路复用 内核阻塞 g挂起
		_, err = c.Write(srespmarshal)
		if err != nil {
			logx.Error(err)
			return
		}
	}
}
