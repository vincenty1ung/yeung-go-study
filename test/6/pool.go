package main

import (
	"encoding/binary"
	"reflect"
	"sync"
	"unsafe"
)

var defaultPool *pool

const (
	defaultMax = 20
)

func init() {
	// 默认最大32k
	defaultPool = New(1, 15, 100)
}

type pool struct {
	pools     []sync.Pool
	bufChan   chan []byte
	closeChan chan struct{}
	pow       []uint8
	closeFlag uint8
	maxB      uint8
}

func New(min, max uint8, chanSize uint) *pool {
	// 最大为1 << 20
	if max > defaultMax {
		max = defaultMax
	}
	if min <= 0 {
		min = 1
	}

	var pool = &pool{
		pools:     make([]sync.Pool, 0, max-min),
		pow:       make([]uint8, 0, max-min),
		bufChan:   make(chan []byte, chanSize),
		closeChan: make(chan struct{}, 1),
	}
	pool.maxB = max
	for min <= max {
		p := sync.Pool{}
		// 这样写的思路是防止外部变量被引用,造成泄漏问题
		switch min {
		case 1:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<1)
			}
		case 2:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<2)
			}
		case 3:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<3)
			}
		case 4:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<4)
			}
		case 5:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<5)
			}
		case 6:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<6)
			}
		case 7:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<7)
			}
		case 8:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<8)
			}
		case 9:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<9)
			}
		case 10:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<10)
			}
		case 11:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<11)
			}
		case 12:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<12)
			}
		case 13:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<13)
			}
		case 14:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<14)
			}
		case 15:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<15)
			}
		case 16:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<16)
			}
		case 17:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<17)
			}
		case 18:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<18)
			}
		case 19:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<19)
			}
		case 20:
			p.New = func() interface{} {
				return make([]byte, 0, 1<<20)
			}
		}
		pool.pools = append(pool.pools, p)
		pool.pow = append(pool.pow, min)
		min++
	}

	// run
	go pool.run()

	return pool
}

func (p *pool) Put(arr []byte) {
	c := cap(arr)
	if c <= 1<<p.maxB && p.closeFlag == 0 {
		for i := 0; i < len(p.pow) && i < len(p.pools); i++ {
			if c == 1<<p.pow[i] {
				binary.BigEndian.PutUint16(arr[:2], uint16(i))
				p.bufChan <- arr
				break
			}
		}
	}
	return
}

func (p *pool) run() {
	for {
		select {
		case buf := <-p.bufChan:
			idx := int(binary.BigEndian.Uint16(buf[:2]))
			for i := 0; i < len(buf); i++ {
				buf[i] = 0
			}
			buf = buf[:0]
			p.pools[idx].Put(buf)
		case <-p.closeChan:
			p.close()
			return
		}
	}
}

func (p *pool) Close() {
	if p.closeFlag == 0 {
		p.closeFlag = 1
		p.closeChan <- struct{}{}
	}
}

func (p *pool) close() {
	close(p.closeChan)
	close(p.bufChan)
	p.pow = nil
	p.pools = nil
	p.maxB = 0
}

func (p *pool) Get(size uint64, init bool) []byte {
	if size <= 1<<p.maxB && p.closeFlag == 0 {
		for i := 0; i < len(p.pow) && i < len(p.pools); i++ {
			if size <= 1<<p.pow[i] {

				buf := p.pools[i].Get().([]byte)
				if init {
					v := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
					v.Len = int(size)
				}

				return buf
			}
		}
	}

	if init {
		return make([]byte, size, size)
	}
	return make([]byte, 0, size)
}

func Get(size uint64, init bool) []byte {
	return defaultPool.Get(size, init)
}

func Put(arr []byte) {
	defaultPool.Put(arr)
}
