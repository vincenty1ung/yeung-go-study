package main

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/net"
)

var chs = make(chan int, 1)
var chs1 = make(chan struct{}, 1)
var chs2 = make(chan struct{})

func write() {
	// 三秒后go1结束
	time.Sleep(3 * time.Second)
	chs1 <- struct{}{}
	// 11秒后go2结束
	time.Sleep(11 * time.Second)
	chs <- 1
}

func go1() {
	// 定时器每800毫秒往通道里放入一个值
	timer2 := time.NewTicker(time.Millisecond * 1000)
	defer timer2.Stop()
	for {
		select {
		case <-timer2.C:
			fmt.Println("goread1")
			// cpu使用率
			times, _ := cpu.Times(true)
			fmt.Println(fmt.Sprintf("%v", times))
			/*info, _ := cpu.Info()
			fmt.Println(fmt.Sprintf("%v",info))*/
			// 连接信息
			connections, _ := net.Connections("tcp")
			fmt.Println(fmt.Sprintf("%v", connections))
		case s := <-chs1:
			fmt.Println(unsafe.Sizeof(s))
			fmt.Println("g1:结束")
			return
		}
	}

}

func go2() {
	for {
		// select无数据阻塞 随机取一个case 当数据可达执行
		select {
		case ch1 := <-chs:
			fmt.Println(unsafe.Sizeof(ch1))
			fmt.Println(ch1)
			close(chs2)
			// select中 return 直接退出函数
			return
			// 获取一个1秒的timer读取通道
		case <-time.After(time.Second):
			fmt.Println("每秒检查数据是否就绪")
			// select中 break 无效
			// default:
			// 	continue

		}

		fmt.Println("select 阻塞拿到数据后执行")
		// 结束循环
		// fmt.Println("结束循环")
		// break
	}
	// fmt.Println("循环之外")
}

var (
	start = make(chan int)
	stop  = make(chan struct{})
)

type handl struct {
	num string
}

func (receiver handl) handler() {
	<-start
	for {
		select {
		case <-time.After(time.Millisecond * 400):
			fmt.Println("执行" + receiver.num)
		case <-stop:
			fmt.Println("结束")
			return
		}
	}
}

func go3() {

	defer close(stop)
	go func() {
		handl{num: "g1"}.handler()
	}()
	go func() {
		handl{num: "g2"}.handler()
	}()
	go func() {
		handl{num: "g3"}.handler()
	}()
	go func() {
		handl{num: "g4"}.handler()
	}()
	go func() {
		handl{num: "g5"}.handler()
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("开始执行")
	close(start)
	time.Sleep(time.Second * 10)
}

func main() {

	// chan 学习
	go go1()
	go go2()
	// var s struct{}
	// fmt.Println(unsafe.Sizeof(s))
	write()

	<-chs2

	name := int64(1)
	fmt.Println("==================")
	// unsafe.Sizeof()返回字节单位
	fmt.Println(unsafe.Sizeof(name)) // 8字节
	names := strconv.Itoa(1)
	fmt.Println(unsafe.Sizeof(names)) // 16
	namel := []string{"1"}
	fmt.Println(unsafe.Sizeof(namel)) // 24
	name8 := uint8(1)
	fmt.Println(unsafe.Sizeof(name8)) // 1
	name32 := uint32(1)
	fmt.Println(unsafe.Sizeof(name32)) // 4

	go3()

}
