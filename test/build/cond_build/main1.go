//go:build ni && n

// 1.17之前都是用这个`+build ni`
package main

import (
	"fmt"
)

func main() {
	fmt.Println("ni")
}
