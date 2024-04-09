package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/uncleyeung/yeung-go-study/utils/pprof"
)

type Client struct {
	conn     net.Conn
	nickname string
}

type Message struct {
	client  *Client
	content string
}

var (
	clients      = make(map[*Client]bool)
	join         = make(chan *Client)
	leave        = make(chan *Client)
	messages     = make(chan Message)
	activeClient = make(chan bool)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("debug:Server started. Listening on localhost:8000")
	pprof.Launch()

	go broadcastMessages()
	go imSend()
	for {
		conn, err := listener.Accept()
		// fmt.Println("debug:一个用户加入")
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{conn: conn}
		clients[client] = true

		go handleClient(client)
	}
}

func handleClient(client *Client) {
	defer func() {
		leave <- client
		client.conn.Close()
	}()

	client.conn.Write([]byte("Enter your nickname: "))

	nicknameBuf := make([]byte, 1024)
	n, err := client.conn.Read(nicknameBuf)
	if err != nil {
		log.Println(err)
		return
	}
	client.nickname = string(nicknameBuf[:n])
	fmt.Println(fmt.Sprintf("debug:当前用户已经加入，用户名字：%s", client.nickname))
	join <- client
	// activeClient <- true

	for {
		messageBuf := make([]byte, 1024)
		n, err := client.conn.Read(messageBuf)
		if err != nil {
			log.Println(err)
			break
		}

		content := string(messageBuf[:n])
		messages <- Message{client: client, content: content}
	}
}

func broadcastMessages() {
	for {
		select {
		case client := <-join:
			message := "User " + client.nickname + " joined the chat\n"
			fmt.Println(fmt.Sprintf("debug:通知所有用户，%s", message))
			sendMessageToAllClients(message, client)
		case client := <-leave:
			delete(clients, client)
			message := "User " + client.nickname + " left the chat\n"
			sendMessageToAllClients(message, nil)
		case message := <-messages:
			fmt.Println(fmt.Sprintf("debug:某个用户发送消息后，推送所有用户，%s", message))
			content := message.client.nickname + ": " + message.content + "\n"
			sendMessageToAllClients(content, message.client)
		}
	}
}

func sendMessageToAllClients(message string, sender *Client) {
	for client := range clients {
		if client != sender {
			_, err := client.conn.Write([]byte(message))
			if err != nil {
				log.Println(err)
				client.conn.Close()
				delete(clients, client)
			}
		}
	}
}
func imSend() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			sendMessageToAllClients("系统定时推送，没10秒一次", nil)

		}
	}

}

// This code sets up a TCP server that listens on `localhost:8000` and accepts incoming client connections.
// Each client is represented by a `Client` struct, which contains the connection and the client's chosen nickname.
// When a client joins the chat, their nickname is read from the connection and they are added to the `clients` map.
// When a client sends a message, it is broadcasted to all connected clients using the `broadcastMessages` goroutine.
// Please note that this is a basic example and doesn't include error handling or advanced features like authentication or message persistence.
// It's intended to demonstrate the basic structure of an IM server in Go.
