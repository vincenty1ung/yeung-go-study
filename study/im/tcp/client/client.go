package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	sendChan              = make(chan string, 120)
	quitChan              = make(chan string)
	readServerMessageChan = make(chan string)
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 监听输入源头
	// fmt.Println("debug:Connected to server.")
	go listenConsole()
	// 从服务器接收响应
	go readServerMessage(err, conn)

	// 接受和发送用两个chan循环 关闭用一个chan 手动关闭输入退出事件exit
	for {
		select {
		case text := <-sendChan:
			// fmt.Println("debug:当前用户发送：" + text)
			sendServerMsg(conn, text, err)
		case <-quitChan:
			// 当前用户退出聊天室
			return
		case text := <-readServerMessageChan:
			// fmt.Println("服务器说:" + text)
			fmt.Println(text)

		}
	}
	// sendServerMsg(err, conn)

}

func sendServerMsg(conn net.Conn, text string, err error) {
	// 向服务器发送消息
	// message := "Hello, server!"
	_, err = conn.Write([]byte(text))
	if err != nil {
		log.Fatal(err)
	}
}

func readServerMessage(err error, conn net.Conn) {
	for {
		select {
		case <-quitChan:
			return
		default:
			response := make([]byte, 1024)
			n, err := conn.Read(response)
			if err != nil {
				log.Fatal(err)
			}
			readServerMessageChan <- fmt.Sprintf("%s\n", response[:n])
		}

	}
}

func listenConsole() {
	// fmt.Println("debug:Listening for console input. Enter 'quit' to exit.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			// fmt.Println("debug:Error reading input:", err)
			continue
		}

		// 去除输入行的换行符
		input = input[:len(input)-1]

		if input == "quit" {
			// fmt.Println("debug:Exiting...")
			// 关闭连接
			quitChan <- "quit"
			break
		}
		// 发送消息
		// fmt.Println("debug:", input)
		sendChan <- input
		// return
	}
}
