package main

import (
	"encoding/json"
	"fmt"
)

/*type s struct {
	A string `json:"a"`
	B string `json:"b"`
}*/
type s struct {
	A string
	B string
}

func getHelloByte() {
	h := "h"
	fmt.Println(fmt.Sprintf("byte h:%08b", []byte(h)))
	e := "e"
	fmt.Println(fmt.Sprintf("byte e:%08b", []byte(e)))

	fmt.Println(fmt.Sprintf("byte he:%08b", []byte("he")))

	s2 := new(s)
	s2.B = "23"
	s2.A = "2"

	marshal, _ := json.Marshal(s2)

	fmt.Println(marshal)
	fmt.Println(string(marshal))
	bytes, _ := json.Marshal(marshal)
	fmt.Println(bytes)
	fmt.Println(string(bytes))

}
