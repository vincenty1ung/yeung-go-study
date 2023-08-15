package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	udpPort = 8000
)

func main() {
	clients := make(map[string]*net.UDPAddr)
	messages := make(chan string)

	// 启动服务端
	go startServer(clients, messages)

	// 启动客户端
	startClient(clients, messages)
}

// 服务端
func startServer(clients map[string]*net.UDPAddr, messages chan string) {
	// 创建UDP地址
	address, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", udpPort))
	if err != nil {
		log.Fatal(err)
	}

	// 创建UDP连接
	conn, err := net.ListenUDP("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Printf("UDP server started on port %d\n", udpPort)

	// 启动接收消息的协程
	go receiveMessages(conn, clients, messages)

	// 广播接收到的消息
	for msg := range messages {
		fmt.Println(msg)
	}
}

// 接收消息
func receiveMessages(conn *net.UDPConn, clients map[string]*net.UDPAddr, messages chan<- string) {
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error while receiving data:", err)
			continue
		}

		message := string(buffer[:n])

		// 将客户端添加到客户端列表中
		clients[addr.String()] = addr

		// 格式化消息
		msg := fmt.Sprintf("[%s]: %s", addr.String(), message)

		// 将消息发送到广播通道
		messages <- msg

		// 发送消息给所有其他客户端
		for _, clientAddr := range clients {
			if clientAddr.String() != addr.String() {
				_, err := conn.WriteToUDP([]byte(msg), clientAddr)
				if err != nil {
					log.Println("Error while sending data:", err)
				}
			}
		}
	}
}

// 客户端
func startClient(clients map[string]*net.UDPAddr, messages chan<- string) {
	// 创建UDP地址
	address, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", udpPort))
	if err != nil {
		log.Fatal(err)
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp", nil, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 从标准输入读取用户输入并发送消息
	sendMessages(conn, messages)
}

// 发送消息
func sendMessages(conn *net.UDPConn, messages chan<- string) {
	for {
		// 从标准输入读取用户输入
		input := make([]byte, 1024)
		n, err := os.Stdin.Read(input)
		if err != nil {
			log.Println("Error while reading input:", err)
			continue
		}

		// 发送消息到服务器
		_, err = conn.Write(input[:n])
		if err != nil {
			log.Println("Error while sending data:", err)
			continue

		}
	}
}
