package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
假设有一台本地机器，无法做加减乘除运算（包括位运算），因此无法执行 a + b、a+ = 1 这样的运算代码
然后我们提供一个服务器端的 HTTP API，可以传两个数字类型的参数，响应结果是这两个参数的和。
HTTP API提供的远程调用方法模拟如下：
func addRemote(a, b int) int{
    // 模拟耗时的操作
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(1000)
    sec := time.Duration(randomNumber)
    time.Sleep(sec* time.Millisecond)
    return a + b
}
现在要求在本地机器上实现一个 sum 函数，支持以下用法：
func main() {
    result1 := sum(11, 23, 76, 25, 5)
    result2 := sum(1, 20, 21, 33, 56, 90)
    result3 := sum(4, 6, 0, 5)
    fmt.Println("result1: ", result1) // 140
    fmt.Println("result2: ", result2) // 221
    fmt.Println("result3: ", result3) // 15
}
要求： sum 能在最短的时间里返回以上结果。
说明：本地代码中的流程控制运算时可以使用的，例如：
for i:=0; i < 10; i ++{
}

*/

func TestSum(t *testing.T) {
	result1 := sum2(11, 23, 76, 25)
	result2 := sum(1, 20, 21, 33, 56, 90)
	result3 := sum(4, 6, 0, 5, 8)
	fmt.Println("result1: ", result1) // 135
	fmt.Println("result2: ", result2) // 221
	fmt.Println("result3: ", result3) // 23
}
func Benchmark_sum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sum(11, 23, 76, 25)
		_ = sum(1, 20, 21, 33, 56, 90)
		_ = sum(4, 6, 0, 5, 8)
	}
}
func Benchmark_sum1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sum1(11, 23, 76, 25)
		_ = sum1(1, 20, 21, 33, 56, 90)
		_ = sum1(4, 6, 0, 5, 8)
		_ = sum1()
	}
}

func sum(ints ...int) int {
	if len(ints) < 2 {
		return ints[0]
	}
	// 0,1->2,3->4,5
	u := len(ints) & 1 // 奇数 第一位肯定是1 偶数第一位肯定是0
	tmp := 0
	if u > 0 {
		tmp += ints[len(ints)-1]
	}
	count := len(ints) >> 1
	j := 0
	k := 1
	for i := 0; i < count; i++ {
		tmp += addRemote(ints[j], ints[k])
		j += 2
		k += 2
	}
	return tmp
}

// [1,1,2,3,4,6,7]
// [1,1,2,3][4,6,7]
func sum1(nums ...int) int {
	group := len(nums) / 2
	loopSumFunc := func(num []int) int {
		loopTmp := 0
		for i := 0; i < len(num); i++ {
			loopTmp = addRemote(loopTmp, num[i])
		}
		return loopTmp
	}
	return addRemote(loopSumFunc(nums[0:group]), loopSumFunc(nums[group:]))
}
func sum2(ints ...int) int {
	tmp := 0
	for i := range ints {
		tmp = addRemote(tmp, ints[i])
	}
	return tmp
}

func sum3(ints ...int) int {
	// 11, 23, 76
	_ = addRemote(addRemote(addRemote(addRemote(0, ints[0]), ints[1]), ints[2]), ints[3]) // 通过for循环把数组元素的相加变成这种
	var tmp int
	for i := 0; i < len(ints); i++ {

	}
	return tmp
}

func addRemote(a, b int) int {
	// 模拟耗时的操作
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000)
	sec := time.Duration(randomNumber)
	time.Sleep(sec * time.Millisecond)
	return a + b
}

/*
思路和算法

虽然题目只要求了不能使用运算符＋、一、* 和/，但是原则上来说也不宜使用类似的运算符+=、 -=、*=和/=，以及sum等方法。于是，我们使用位运算来处理这个问题。

首先，考虑两个二进制位相加的四种情况如下：

0 + 0 = 0
0 + 1 = 1
1 + 0 = 1
1 + 1 = 0 (进位)

可以发现，对于整数a和6：

	· 在不考虑进位的情况下，其无进位加法结果为a田b。
	· 而所有需要进位的位为a&b，进位后的进位结果为(a&b）<<1。

于是，我们可以将整数a和b的和，拆分为。a和b的无进位加法结果与进位结果的和。因为每一次拆分都可以让需要进位的最低位至少左移一位，又因为a和6可以取到负数，所以我们最多需要log(marrint） 次拆分即可完成运算。
*/
func add(x, y int) int {
	for y != 0 {
		// 异或操作得到无进位加法结果
		sum := x ^ y

		// 与操作并左移一位得到进位值
		carry := (x & y) << 1

		x = sum
		y = carry
	}

	return x
}

func subtract(x, y int) int {
	for y != 0 {
		// 取反加一得到 y 的补码
		complement := (^y) + 1

		// 使用加法计算差值
		x = add(x, complement)

		y = 1 // 用于退出循环
	}
	return x
}
