package test

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	redisn "github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid/v5"

	"go.uber.org/zap"

	"github.com/vincenty1ung/yeung-go-study/utils/ylog"
	"github.com/vincenty1ung/yeung-go-study/utils/yredis"
)

func TestLock(t *testing.T) {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	// uuid.Gen{}
	background := context.Background()

	go func() {
		v1, _ := uuid.NewV1()
		value1 := context.WithValue(background, yredis.ContextKey, "go1 id:"+v1.String())
		redisLock1 := yredis.NewRedisLock(value1, redisX, "test:lock")

		defer func() {
			err := redisLock1.Unlock()
			if err != nil {
				ylog.Error("", zap.Error(err))
			}
			wg.Done()
		}()

		err := redisLock1.Lock()
		if err != nil {
			ylog.Error("", zap.Error(err))
			return
		}
		fmt.Println("go1 dosomething")
		<-time.After(time.Millisecond * 2900)
		fmt.Println("go1 dosomething sussce")

	}()
	go func() {
		v2, _ := uuid.NewV1()
		value2 := context.WithValue(background, yredis.ContextKey, "go2 id:"+v2.String())
		redisLock2 := yredis.NewRedisLock(value2, redisX, "test:lock")

		defer wg.Done()
		err := redisLock2.Lock()
		if err != nil {
			ylog.Error("", zap.Error(err))
			return
		}
		fmt.Println("go2 dosomething")
		<-time.After(time.Millisecond * 100)
		fmt.Println("go2 dosomething sussce")
		err = redisLock2.Unlock()
		if err != nil {
			ylog.Error("", zap.Error(err))
		}
	}()

	go func() {
		v3, _ := uuid.NewV1()
		value3 := context.WithValue(background, yredis.ContextKey, "go3 id:"+v3.String())
		redisLock3 := yredis.NewRedisLock(value3, redisX, "test:lock")

		defer wg.Done()
		err := redisLock3.Lock()
		if err != nil {
			ylog.Error("", zap.Error(err))
			return
		}
		fmt.Println("go3 dosomething")
		<-time.After(time.Millisecond * 30)
		fmt.Println("go3 dosomething sussce")
		err = redisLock3.Unlock()
		if err != nil {
			ylog.Error("", zap.Error(err))
		}
	}()

	wg.Wait()
	fmt.Print("总计执行:")
	fmt.Println(time.Since(now))
}
func TestLock1(t *testing.T) {
	now := time.Now()
	// uuid.Gen{}
	background := context.Background()
	v1, _ := uuid.NewV1()
	value1 := context.WithValue(background, yredis.ContextKey, "go1 id:"+v1.String())
	redisLock1 := yredis.NewRedisLock(value1, redisX, "test:lock")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			err := redisLock1.Unlock()
			if err != nil {
				ylog.Error("", zap.Error(err))
			}
			wg.Done()
		}()

		err := redisLock1.Lock()
		if err != nil {
			ylog.Error("", zap.Error(err))
			return
		}
		fmt.Println("go1 dosomething")
		<-time.After(time.Millisecond * 2900)
		fmt.Println("go1 dosomething sussce")

	}()
	wg.Wait()
	fmt.Print("总计执行:")
	fmt.Println(time.Since(now))
}

func TestUUID(t *testing.T) {
	// u1 := uuid.Must(uuid.NewV4())

	// Create a Version 4 UUID.
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u2)

	// Parse a UUID from a string.
	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	u3, err := uuid.FromString(s)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", s, err)
	}
	log.Printf("successfully parsed UUID %v", u3)
}

func TestGet(tt *testing.T) {
	// C.myprint("s")
	g := "g" // 6750208 [0110,0111,0000,0000,0000,0000]
	e := 'e' // 25856               [0110,0101,0000,0000]
	t := "t" // 116                             [0111,0100]
	// 6776180 ==0x676574
	ten := 6776180
	ten = 0x676574
	ij := ten & (int(math.Pow(2, 8)) - 1)
	fmt.Println(ij)

	tmp := 1<<8 - 1
	i1 := ten & tmp
	i2 := (ten & (tmp << 8)) >> 8
	i3 := (ten & (tmp << 16)) >> 16

	fmt.Println(string(rune(i3)))
	fmt.Println(string(rune(i2)))
	fmt.Println(string(rune(i1)))

	fmt.Println([]byte(g))
	fmt.Println(rune(e))
	fmt.Println([]byte(t))
	fmt.Println([]byte("get"))

	var builder strings.Builder
	builder.WriteByte(byte(116))
	builder.WriteByte(byte(101))
	builder.WriteByte(byte(103))
	builder.WriteString("你好=")

	fmt.Println(builder.String())
	builder.Reset()
	fmt.Println(builder.String() + "end")

	fmt.Println(logs("你好", "世界", "哈哈"))
	var sbs strings.Builder
	sbs.WriteString("sd ")

	/*sb := (*strings.Builder)(nil)
	sb.WriteString("s")*/

	var v url.Values
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew/")
	v.Add("cat sounds", "mau$")
	fmt.Println(v.Encode())

}

func TestName(t *testing.T) {
	var v url.Values
	var builder strings.Builder
	builder1 := strings.Builder{}
	s := (*strings.Builder)(nil)
	s2 := struct {
	}{}
	s3 := struct {
	}{}

	fmt.Println(fmt.Sprintf("%p", &v))
	fmt.Println(fmt.Sprintf("%p", v))
	fmt.Printf("builder:p: %p\n", &builder)
	fmt.Printf("builder: %#v\n", builder)
	builder.Grow(1)
	fmt.Printf("builder:p: %p\n", &builder)
	fmt.Printf("builder: %#v\n", builder)
	fmt.Printf("builder1:p: %#p\n", &builder1)
	fmt.Printf("builder1: %#v\n", builder1)
	// fmt.Println(&builder)
	fmt.Println(fmt.Sprintf("%p", &s))
	fmt.Printf("s: %#v\n", s)
	fmt.Println(fmt.Sprintf("%p", &s2))
	fmt.Printf("s2: %#v\n", s2)
	fmt.Println(fmt.Sprintf("%p", &s3))
	fmt.Printf("s2: %#v\n", s3)
}
func logs(strs ...string) string {
	builder := strings.Builder{}
	defer builder.Reset()
	for i := range strs {
		builder.WriteString(strs[i])
	}
	target := builder.String()
	return target
}
func logs2(strs ...string) string {
	var s string
	for i := range strs {
		s += strs[i]
	}
	return s
}

func BenchmarkStrBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = logs(
			"你好", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈",
		)
	}
}
func BenchmarkStrBuilder2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = logs2(
			"你好", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈",
			"世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界", "哈哈", "世界",
			"哈哈",
		)
	}
}
func BenchmarkStr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = "你好" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈" + "世界" + "哈哈"
	}
}

func TestNamehha(t *testing.T) {
	userid := "lock:key12938923893737763"
	background := context.Background()
	boolCmd, _ := redisX.SetNX(background, userid, 0, time.Second*60).Result()
	if !boolCmd {
		incr, _ := redisX.Incr(background, userid).Result()
		if incr > 10 {
			return
		}
	}
}

type ZV struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func TestZset(t *testing.T) {
	key := "ztest2"
	background := context.Background()
	zvs := make([]*redisn.Z, 0)
	zvs = append(
		zvs, &redisn.Z{
			Score: 12, Member: "chenfang",
		},
	)
	zvs = append(
		zvs, &redisn.Z{
			Score: 11, Member: "yanbo",
		},
	)
	result, err := redisX.ZAdd(background, key, zvs...).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
