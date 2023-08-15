package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	/*	go func() {
			err := http.ListenAndServe("localhost:6060", nil)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	*/
	/*	a := []int{1, 2, 3, 5}
		// 结果就是这个刚开始p = []int{0,0,0} 变成底下的
		p := []int{1, 3, 6, 11}*/
	_ = make([]int, 0, 3)
	http.HandleFunc(
		"/test", func(writer http.ResponseWriter, request *http.Request) {
			bytes := getbuf()
			for i := range bytes {
				bytes[i] = 1
			}
			fmt.Fprint(writer, "你好世界")
		},
	)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func getbuf() []byte {
	return make([]byte, 1<<32)
}

// ./gcgo
// ab -n 500 -c 100 http://localhost:8080/test

// http://127.0.01:8080/debug/pprof/
// wget  http://127.0.0.1:8080/debug/pprof/trace\?seconds\=20 -O trace.out
// wget  http://127.0.0.1:8080/debug/pprof/profile\?seconds\=20 -O profile.out [trace]
// wget  http://127.0.0.1:8080/debug/pprof/heap\?seconds\=20 -O heap.out [trace]
// go tool trace trace.out
// go tool pprof http://localhost:8080/debug/pprof/profile ~/pprof/pprof.samples.cpu.001.pb.gz  [CPU Profiling]
// go tool pprof http://localhost:8080/debug/pprof/heap ~/pprof/pprof.samples.heap.001.pb.gz [Memory Profiling]
// go tool pprof http://localhost:8080/debug/pprof/trace  [trace Profiling]
