package main

import (
	"net/http"
	"time"

	"github.com/r3labs/sse/v2"
)

func main() {
	server := sse.New()
	_ = server.CreateStream("messages")
	// Create a new Mux and set the handler
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/events", func(w http.ResponseWriter, r *http.Request) {
			go func() {
				// Received Browser Disconnection
				<-r.Context().Done()
				println("The client is disconnected here")
				return
			}()

			server.ServeHTTP(w, r)
		},
	)

	go func() {
		for {
			server.Publish(
				"messages", &sse.Event{
					Data: []byte("ping"),
				},
			)
			time.Sleep(1 * time.Second)
		}
	}()

	http.ListenAndServe(":8080", mux)
}
