package main

import (
	"fmt"

	"github.com/r3labs/sse/v2"
)

func main() {
	client := sse.NewClient("http://localhost:8080/events")
	err := client.SubscribeRaw(
		func(msg *sse.Event) {
			// Got some data!
			fmt.Println(string(msg.Data))
			// fmt.Println(string(msg.ID))
		},
	)
	if err != nil {
		panic(err)
	}
}
