package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
    printf("%s\n", s);
}
*/

import (
	"fmt"
	"strconv"

	"C"
)

type name struct {
}

func main() {

	/*	v := 1
		i := If(v == 2, 2, If(v == 3, 3, 1))
		_ = i*/
	int64s := make([]int64, 0, 100)
	for i := 0; i < 100; i++ {
		int64s = append(int64s, int64(i))
	}
	// test1(int64s)
	is := int64s[100:]

	fmt.Println(is)

	fmt.Println(1002 / 500)
	fmt.Println(1002 % 500)
	fmt.Println(499 % 500)
	fmt.Println(499 / 500)

	C.myprint("s")

}

func If(cond bool, v1, v2 interface{}) interface{} {
	if cond {
		return v1
	}
	return v2
}

var batchSize = 100

func test1(ids []int64) {
	batch := len(ids)/batchSize + 1
	// 100 /100 =1 batch 2
	for i := 0; i < batch; i++ {
		var batchIDs []int64
		if i < batch-1 {
			batchIDs = ids[i*batchSize : (i+1)*batchSize]
		} else {
			batchIDs = ids[i*batchSize:]
		}
		toAccount := make([]string, len(batchIDs))
		for i, to := range batchIDs {
			toAccount[i] = strconv.FormatInt(to, 10)
		}
		send(toAccount)
	}
}

func send(account []string) {
	fmt.Println(fmt.Sprintf("发送数据:%v", account))
}
