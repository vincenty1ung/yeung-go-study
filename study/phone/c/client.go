package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gordonklaus/portaudio"
	"github.com/pion/webrtc/v3"
)

const (
	SampleRate  = 44100
	BufferSize  = 256
	NumChannels = 1
)

func main() {
	// 初始化PortAudio
	if err := portaudio.Initialize(); err != nil {
		fmt.Printf("PortAudio initialization failed: %v\n", err)
		return
	}
	defer portaudio.Terminate()

	// 打开默认音频输入设备
	inputDevice, err := portaudio.DefaultInputDevice()
	if err != nil {
		fmt.Printf("Error opening default input device: %v\n", err)
		return
	}

	// 打开默认音频输出设备
	outputDevice, err := portaudio.DefaultOutputDevice()
	if err != nil {
		fmt.Printf("Error opening default output device: %v\n", err)
		return
	}

	// 设置音频参数
	streamParameters := portaudio.HighLatencyParameters(inputDevice, outputDevice)

	// 打开音频流
	stream, err := portaudio.OpenStream(streamParameters, audioCallback)
	if err != nil {
		fmt.Printf("Error opening audio stream: %v\n", err)
		return
	}
	defer stream.Close()

	// 开始音频流
	if err := stream.Start(); err != nil {
		fmt.Printf("Error starting audio stream: %v\n", err)
		return
	}

	// 连接到服务器
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:12345")
	if err != nil {
		log.Fatalf("Error resolving UDP address: %v", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("Error dialing UDP: %v", err)
		return
	}
	defer conn.Close()

	// 发送连接成功消息给服务器
	_, err = conn.Write([]byte{1})
	if err != nil {
		log.Fatalf("Failed to send UDP data: %v", err)
		return
	}

	// 接收offer和ICE候选信息
	offerBuf := make([]byte, 2048)
	_, err = conn.Read(offerBuf)
	if err != nil {
		log.Fatalf("Failed to read offer from server: %v", err)
		return
	}

	// 创建WebRTC连接配置
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{},
	}

	// 创建WebRTC连接
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&webrtc.MediaEngine{}))
	connection, err := api.NewPeerConnection(config)
	if err != nil {
		log.Fatalf("Error creating WebRTC connection: %v", err)
		return
	}
	defer connection.Close()

	// 解析offer
	offer := webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  string(offerBuf),
	}
	if err := connection.SetRemoteDescription(offer); err != nil {
		log.Fatalf("Error setting remote description: %v", err)
		return
	}

	// 创建answer
	answer, err := connection.CreateAnswer(nil)
	if err != nil {
		log.Fatalf("Error creating answer: %v", err)
		return
	}

	// 设置本地描述
	if err := connection.SetLocalDescription(answer); err != nil {
		log.Fatalf("Error setting local description: %v", err)
		return
	}

	// 将ICE候选信息发送给服务器
	connection.OnICECandidate(
		func(candidate *webrtc.ICECandidate) {
			if candidate != nil {
				// 将ICE候选信息通过UDP发送给服务器
				_, err := conn.Write([]byte(candidate.ToJSON().Candidate))
				if err != nil {
					log.Printf("Failed to send ICE candidate to server: %v", err)
				}
			}
		},
	)

	// 创建ICE监听器
	iceChan := make(chan *webrtc.ICECandidate, 1)
	connection.OnICECandidate(
		func(candidate *webrtc.ICECandidate) {
			if candidate != nil {
				iceChan <- candidate
			}
		},
	)

	// 发送answer给服务器
	_, err = conn.Write([]byte(answer.SDP))
	if err != nil {
		log.Fatalf("Failed to send answer to server: %v", err)
		return
	}

	// 等待ICE候选信息
	iceCandidate := <-iceChan

	// 将ICE候选信息通过UDP发送给服务器
	_, err = conn.Write([]byte(iceCandidate.ToJSON().Candidate))
	if err != nil {
		log.Printf("Failed to send ICE candidate to server: %v", err)
	}

	// 捕获Ctrl+C信号，以便在程序退出时关闭音频流
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Exiting...")
}

/*
// 音频回调函数，处理音频数据
func audioCallback(_, outBuf []int32) {
	// 在这里你可以处理音频数据，例如进行音频解码、播放等
	// 在这个简化的示例中，我们只是简单地将输出缓冲区清零
	for i := range outBuf {
		outBuf[i] = 0
	}

	// TODO: 在这里进行音频解码和处理逻辑
}
*/

// 音频回调函数，处理音频数据
func audioCallback(inBuf, outBuf []int32) {
	// 在这里你可以处理音频数据，例如进行音频编码、网络传输等
	// 在这个简化的示例中，我们只是简单地将输入缓冲区的数据发送给客户端

	// TODO: 在这里进行音频编码和发送给客户端的逻辑
}
