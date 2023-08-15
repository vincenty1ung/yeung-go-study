package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	go tcpPacketFragmentation()
	httpSOCKET()
}

func tcpPacketFragmentation() {
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8081")
	if err != nil {
		logs.Error("main错误： %s", err)
		return
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		logs.Error("main错误： %s", err)
		return
	}
	if err != nil {
		logs.Error("main错误： %s", err)
		return
	}
	go createClient()
	for true {
		conn, err := listen.Accept()
		if err != nil {
			logs.Error("main错误： %s", err)
			return
		}
		go getRead(conn, '\n')
	}
}

var (
	sended = make(chan struct{})
	recved = make(chan struct{})
)

func createClient() {
	time.Sleep(3 * time.Second)
	tcpConn, err := net.DialTCP(
		"tcp", nil, &net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8081,
		},
	)
	defer tcpConn.Close()
	if err != nil {
		logs.Error("createClient错误： %s", err)
		return
	}
	for {
		txt := "Hw\n"
		_, err := tcpConn.Write([]byte(txt))
		// fmt.Println(write)
		if err != nil {
			fmt.Println(err)
			break
		}
		// tcpConn.Write([]byte(txt))
		// tcpConn.Write([]byte(txt))
		// sended <- struct{}{}
		// <-recved
		/*reads := make([]byte, 1024)
		tcpConn.Read(reads)

		fmt.Println(string(reads))*/
	}

}
func getRead(conn net.Conn, key byte) {
	reads := make([]byte, 1024)
	c := 0
	for true {
		// <-sended
		n, err := conn.Read(reads)
		if err != nil {
			logs.Error("getRead错误： %s", err)
		}
		/*	i := bytes.IndexByte(reads, key)
			if i != -1 {
				fmt.Println(fmt.Sprintf("%s", reads[:i]))
				reads = reads[i+1:]
			} else {
				fmt.Println("ssss")
			}*/
		fmt.Println(fmt.Sprintf("%s", reads[:n]))

		// recved <- struct{}{}
		if n < 0 {
			fmt.Println("n 小于 0")
			continue
		}
		c++
		if c == 100000 {
			fmt.Println(len(reads))
		}
		reads = make([]byte, 1024)
		fmt.Println(fmt.Sprintf("第%v次", c))
	}
}

func httpSOCKET() {
	// 创建一个新的http路由管理器
	mux := http.NewServeMux()
	mux.HandleFunc("/index", indexHandler)
	// 只是监听8080端口
	_ = http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	GetBody := r.GetBody
	logs.Info("这是一个很开心的脸 %s %v", GetBody, 3)
	_, _ = fmt.Fprint(w, "你好世界")
}

type Gf func(ctx2 context.Context)

func goFunc(ctx context.Context, gf Gf) {
	go func() {
		// defer api-log
		gf(ctx)
	}()
}
