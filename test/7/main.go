package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func hError() error {
	return nil
}

const TimeYyyyMmDd = "2006-01-02"

func main() {

	d, _ := time.Parse(time.RFC3339, "2022-07-01T04:04:05Z")
	fmt.Println(d)

	d = d.AddDate(0, 0, -1)
	fmt.Println(d)

	d = d.AddDate(0, 0, -d.Day()+1)
	fmt.Println(d)

	timeStr := time.Now().Format(TimeYyyyMmDd)
	d, _ = time.ParseInLocation(TimeYyyyMmDd, timeStr, time.Local)
	fmt.Println(d)

	go hError()
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)

	fmt.Println(strings.TrimSuffix("0.3.3.3", "."+strconv.Itoa(3)))

	goFunc(
		context.Background(), func(ctx2 context.Context) {
			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func() {
					ch <- true
					count++
					time.Sleep(time.Millisecond)
					count--
					<-ch
					wg.Done()
				}()
			}

			fmt.Println(count)
			b := make([]byte, 0, 1024)
			buffer := bytes.NewBuffer(b)
			var mapMan map[string]string
			mapMan = make(map[string]string)
			mapMan["s"] = "v"
			decoder := json.NewEncoder(buffer)
			decoder.SetEscapeHTML(false)

			err := decoder.Encode(&mapMan)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(buffer.String())

			strings := make([]string, 0, 0)
			strings = append(strings, "")
			fmt.Println(buffer.String())
		},
	)
	time.Sleep(time.Second * 3)
	wg.Wait()
}

type Gf func(ctx2 context.Context)

func goFunc(ctx context.Context, gf Gf) {
	go func() {
		// defer api-log
		gf(ctx)
	}()
}
