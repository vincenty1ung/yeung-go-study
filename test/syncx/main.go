package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*fmt.Println("Hello, playground")
	m := new(MuOnce)
	fmt.Println(m.strings())
	fmt.Println(m.strings())
	get, _ := http.Get("")
	body, err := ioutil.ReadAll(get.Body)
	fmt.Printf("body len:%v, read err:%v", len(body), err)
	get.Body.Close()*/
	b := make([]byte, 0, 6)
	bytes := b[len(b):cap(b)]
	fmt.Println(bytes)
	i := []byte{1, 1, 23}
	n := copy(bytes, i[0:])
	fmt.Println(n)
	fmt.Println(i)
	fmt.Println(bytes)
	b = b[:len(b)+n]
	fmt.Println(b)
	fmt.Println("=========2=========")
	i = []byte{2, 2, 38}
	fmt.Println(i)
	bytes = b[len(b):cap(b)]
	fmt.Println(bytes)
	n = copy(bytes, i[0:])
	b = b[:len(b)+n]
	fmt.Println(b)

	b = append(b, 0)[:len(b)]
	fmt.Println(b)
	bytes = b[len(b):cap(b)]
	fmt.Println(bytes)
	fmt.Println(fmt.Sprintf("byte h:%08b", []byte(`h`)))
	// fmt.Println(fmt.Sprintf("byte h:%08b", byte(1)))
	i2 := 2*2*2*2*2*2*2*2 - 1
	i2 = 2 ^ 8 - 1
	fmt.Println(i2)
	fmt.Println(fmt.Sprintf("byte h:%016b", byte(i2)))
	fmt.Println(fmt.Sprintf("byte h:%016b", int(-i2)))
	fmt.Println(fmt.Sprintf("byte h:%08b", int8(127)))
	fmt.Println(fmt.Sprintf("byte h:%08b", int8(-127)))
	fmt.Println(fmt.Sprintf("byte h:%08b", int8(-128)))
	fmt.Println(fmt.Sprintf("byte h:%08b", uint8(128)))
	fmt.Println(fmt.Sprintf("byte h:%08b", uint8(255)))

}

type MuOnce struct {
	sync.RWMutex
	sync.Once
	mtime time.Time
	vals  []string
}

// 相当于reset方法，会将m.Once重新复制一个Once
func (m *MuOnce) refresh() {
	m.Lock()
	defer m.Unlock()
	m.Once = sync.Once{}
	m.mtime = time.Now()
	m.vals = []string{m.mtime.String()}
}

// 获取某个初始化的值，如果超过某个时间，会reset Once
func (m *MuOnce) strings() []string {
	now := time.Now()
	m.RLock()
	if now.After(m.mtime) {
		defer m.Do(m.refresh) // 使用refresh函数重新初始化
	}
	vals := m.vals
	m.RUnlock()
	return vals
}
