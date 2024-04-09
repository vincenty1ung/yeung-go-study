package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定Content-Type为text/event-stream
	w.Header().Set("Content-Type", "text/event-stream")
	// 设置缓存控制，禁用缓存
	w.Header().Set("Cache-Control", "no-cache")
	// 设置连接保持活动
	w.Header().Set("Connection", "keep-alive")
	// 允许所有来源跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 创建通道用于通知连接状态
	notify := w.(http.CloseNotifier).CloseNotify()
	// 通道用于停止事件发送
	stopChan := make(chan struct{})

	// 向客户端发送初始数据
	fmt.Fprintf(w, "data: %s\n\n", "Initial message")

	// 每秒发送一次数据
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-notify:
			// 客户端连接断开，清理资源
			log.Println("Client disconnected")
			close(stopChan)
			return
		case <-ticker.C:
			// 向客户端发送新的数据
			fmt.Fprintf(w, "data: %s\n\n", time.Now().Format("15:04:05"))
			w.(http.Flusher).Flush()
		case <-stopChan:
			// 停止事件发送
			return
		}
	}
}

func main() {
	http.HandleFunc("/sse", handleSSE)

	log.Println("Server is running on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
