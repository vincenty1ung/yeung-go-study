package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

// 客户端
func client() {
	var group sync.WaitGroup
	// 出初始化五个客户端
	for i := 0; i < 5; i++ {
		group.Add(1)
		logx.Info(fmt.Sprintf("client: 开始执行携程%d", i))
		go func(int2 int, w *sync.WaitGroup) {
			dial, err := net.Dial("tcp", "127.0.0.1:8099")
			if err != nil {
				logx.Error(err)
				return
			}
			defer func(dial net.Conn) {
				err := dial.Close()
				if err != nil {
					logx.Error(err)
				}
			}(dial)
			index := 0
			for {
				r := &req{
					Msg:       fmt.Sprintf("我是客户端%d", int2),
					Num:       index,
					LocalAddr: dial.LocalAddr().String(),
				}
				marshal, err := json.Marshal(r)
				if err != nil {
					logx.Error(err)
					return
				}
				// io多路复用 内核阻塞 g挂起
				_, err = dial.Write(marshal)
				if err != nil {
					logx.Error(err)
					return
				}

				bytes := make([]byte, 1024)
				// io多路复用 内核阻塞 g挂起
				rsn, err := dial.Read(bytes)
				if err != nil {
					logx.Error(err)
					return
				}
				res := new(resp)
				err = json.Unmarshal(bytes[:rsn], res)
				if err != nil {
					logx.Error(err)
					return
				}

				logx.Info(fmt.Sprintf("client:收到服务端响应: %v", res))
				time.Sleep(time.Millisecond * 10)
				index++
				if index == 20 {
					w.Done()
					return
				}
			}

		}(i, &group)
		// logx.Info(fmt.Sprintf("client: 休息20秒在开始新的携程"))
		// time.Sleep(time.Second * 20)
	}

	// 结束服务端
	group.Wait()
	close(chanEXIT)

}
