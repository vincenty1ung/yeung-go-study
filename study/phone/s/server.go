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

	// 设置音频参数
	/*inputParams := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{Device: inputDevice, Channels: NumChannels},
	}*/
	streamParameters := portaudio.HighLatencyParameters(inputDevice, nil)

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

	// 监听UDP连接
	addr, err := net.ResolveUDPAddr("udp", ":12345")
	if err != nil {
		log.Fatalf("Error resolving UDP address: %v", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Error listening on UDP: %v", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server is ready. Waiting for client...")

	// 接收客户端连接
	_, clientAddr, err := conn.ReadFromUDP(make([]byte, 1))
	if err != nil {
		log.Fatalf("Error reading from UDP: %v", err)
		return
	}

	fmt.Printf("Client connected: %s\n", clientAddr.String())

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

	// 创建音轨
	/*audioTrack, err := connection.NewAudioTrack(webrtc.DefaultPayloadTypeOpus, 1, "audio", "pion")
	if err != nil {
		log.Fatalf("Error creating audio track: %v", err)
		return
	}*/
	audioTrack, err := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, "audio", "pion1",
	)
	if err != nil {
		log.Fatalf("Error creating audio track: %v", err)
		return
	}

	// 添加音轨到连接
	if _, err = connection.AddTrack(audioTrack); err != nil {
		log.Fatalf("Error adding audio track to connection: %v", err)
		return
	}

	// 创建数据通道
	dataChannel, err := connection.CreateDataChannel("data", nil)
	if err != nil {
		log.Fatalf("Error creating data channel: %v", err)
		return
	}

	// 设置数据通道的处理函数
	dataChannel.OnMessage(
		func(msg webrtc.DataChannelMessage) {
			// 在这里处理从客户端接收到的数据
			fmt.Printf("Received message from client: %s\n", string(msg.Data))
		},
	)

	// 将ICE候选信息发送给客户端
	connection.OnICECandidate(
		func(candidate *webrtc.ICECandidate) {
			if candidate != nil {
				fmt.Printf("Sending ICE candidate to client\n")
				// 将ICE候选信息通过数据通道发送给客户端
				if err := dataChannel.SendText(candidate.ToJSON().Candidate); err != nil {
					log.Printf("Failed to send ICE candidate to client: %v", err)
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

	// 创建offer
	offer, err := connection.CreateOffer(nil)
	if err != nil {
		log.Fatalf("Error creating offer: %v", err)
		return
	}

	// 设置本地描述
	if err := connection.SetLocalDescription(offer); err != nil {
		log.Fatalf("Error setting local description: %v", err)
		return
	}

	// 等待ICE候选信息
	iceCandidate := <-iceChan

	// 发送offer和ICE候选信息给客户端
	if _, err := conn.Write([]byte(offer.SDP)); err != nil {
		log.Fatalf("Failed to send offer to client: %v", err)
		return
	}
	if _, err := conn.Write([]byte(iceCandidate.ToJSON().Candidate)); err != nil {
		log.Fatalf("Failed to send ICE candidate to client: %v", err)
		return
	}

	// 等待Ctrl+C信号，以便在程序退出时关闭音频流
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Exiting...")
}

// 音频回调函数，处理音频数据
func audioCallback(inBuf, _ []int32) {
	// 在这里你可以处理音频数据，例如进行音频编码、网络传输等
	// 在这个简化的示例中，我们只是简单地将输入缓冲区的数据发送给客户端

	// TODO: 在这里进行音频编码和发送给客户端的逻辑
}

// # 安装gstreamer
// sudo apt-get install libgstreamer1.0-0 gstreamer1.0-plugins-base gstreamer1.0-plugins-good gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly gstreamer1.0-libav
//
// # 安装gstreamer的Go绑定
// go get -u github.com/hajimehoshi/gstreamer
