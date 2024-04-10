package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dir, err := os.OpenFile("/home/vincentyeung/tmp/1.txt", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dir)
	builder := strings.Builder{}
	for i := 0; i < 1024; i++ {
		builder.WriteString("a" + strconv.Itoa(i))
	}

	ticker := time.NewTicker(time.Second)
	count := 0
	for {
		if count == 600 {
			return
		}
		select {
		case <-ticker.C:
			count++
			n, err := dir.WriteString(builder.String())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(n)
			if err := dir.Sync(); err != nil {
				fmt.Println(err)
				continue
			}

		}
	}
}
