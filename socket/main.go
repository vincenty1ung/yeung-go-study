package main

import (
	"fmt"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

var (
	chanEXIT = make(chan struct{}, 1)
	// chanStart = make(chan struct{}, 1)
)

type (
	req struct {
		Msg       string `json:"msg"`
		Num       int    `json:"num"`
		LocalAddr string `json:"local_addr"`
	}

	resp struct {
		Msg       string `json:"msg"`
		LocalAddr string `json:"local_addr"`
	}
)

func (r req) String() string {
	return fmt.Sprintf("Msg:%s,Num:%d,LocalAddr:%s", r.Msg, r.Num, r.LocalAddr)
}
func (r resp) String() string {
	return fmt.Sprintf("Msg:%s,LocalAddr:%s", r.Msg, r.LocalAddr)
}

func main() {
	// 服务端
	go server()
	// 客户端
	time.Sleep(time.Second * 2)
	//client()
	pollclient()
	time.Sleep(time.Second * 10)
	logx.Info("main: exit")
}
