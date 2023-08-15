package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

// 简易池化
var (
	_cPool *sync.Pool
)

func clientPoolInit(num int) {
	_cPool = &sync.Pool{
		New: func() interface{} {
			dial, err := net.Dial("tcp", "127.0.0.1:8099")
			if err != nil {
				logx.Error(err)
			}
			return dial
		},
	}

	for i := 0; i < num; i++ {
		dial, err := net.Dial("tcp", "127.0.0.1:8099")
		if err != nil {
			logx.Error(err)
			continue
		}
		putClientPool(dial)
	}
}

func getClientPool() net.Conn {
	if conn, ok := _cPool.Get().(net.Conn); ok {
		return conn
	}
	return nil
}

func putClientPool(c net.Conn) {
	_cPool.Put(c)
}

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

// 池化客户端
func pollclient() {
	clientPoolInit(5)
	var group sync.WaitGroup
	// 出初始化五个客户端
	for i := 0; i < 5; i++ {
		group.Add(1)
		logx.Info(fmt.Sprintf("阶段1:client: 开始执行携程%d", i))
		go func(int2 int, w *sync.WaitGroup) {
			dial := getClientPool()
			defer func(dialf net.Conn) {
				putClientPool(dialf)
			}(dial)
			index := 0
			for {
				r := &req{
					Msg:       fmt.Sprintf("阶段1:我是客户端%d", int2),
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

				logx.Info(fmt.Sprintf("阶段1:client:收到服务端响应: %v", res))
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

	// 阶段1结束
	logx.Info(fmt.Sprintf("阶段1结束"))
	var group1 sync.WaitGroup
	for i := 0; i < 5; i++ {
		group1.Add(1)
		logx.Info(fmt.Sprintf("阶段2:client: 开始执行携程%d", i))
		go func(int2 int, w *sync.WaitGroup) {
			dial := getClientPool()
			defer func(dialf net.Conn) {
				putClientPool(dialf)
			}(dial)
			index := 0
			for {
				r := &req{
					Msg:       fmt.Sprintf("阶段2:我是客户端%d", int2),
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

				logx.Info(fmt.Sprintf("阶段2:client:收到服务端响应: %v", res))
				time.Sleep(time.Millisecond * 10)
				index++
				if index == 20 {
					w.Done()
					return
				}
			}

		}(i, &group1)
		// logx.Info(fmt.Sprintf("client: 休息20秒在开始新的携程"))
		// time.Sleep(time.Second * 20)
	}
	group1.Wait()
	close(chanEXIT)
}
