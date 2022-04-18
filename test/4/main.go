package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type (
	Base struct {
		Id  int
		Age int
	}
	Man struct {
		Base
		Jj uint8
	}
	Woman struct {
		Base
		Bb uint8
	}
	TypeV interface {
		GetType() VTYPE
	}
	VTYPE int32
)

const (
	VTYPE_MAN  = iota // 男struct
	VTYPE_WMAN = 1    // 女struct
)

var (
	name = map[VTYPE]interface{}{}
)

func (w Woman) GetType() VTYPE {
	return VTYPE_WMAN
}
func (w Man) GetType() VTYPE {
	return VTYPE_MAN
}

func (w Woman) String() string {
	panic("implement me")
}

func hand(v TypeV) {
	switch v.GetType() {
	case VTYPE_MAN:
		if v, ok := v.(Man); ok {
			fmt.Println(v)
		}
	case VTYPE_WMAN:
		if v, ok := v.(Woman); ok {
			fmt.Println(v)
		}
	}
}

func main() {
	var mt TypeV
	var wmt TypeV
	man := Man{
		Jj: 1,
	}
	man.Id = 1
	man.Age = 1
	mt = man

	woman := Woman{
		Bb: 11,
	}
	woman.Id = 2
	woman.Age = 32
	wmt = woman

	//mt.GetType()
	//wmt.GetType()

	hand(mt)
	hand(wmt)

	var state int32
	fmt.Println(state)
	state = mutexLocked
	new1 := state
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)
	new1 += 1 << mutexWaiterShift
	fmt.Println(new1)

	fmt.Println("===========")
	new1 |= mutexWoken // 加状态
	fmt.Println(new1)

	new1 &^= mutexWoken // 去状态
	fmt.Println(new1)

	addInt32 := atomic.AddInt32(&new1, -mutexLocked)
	fmt.Println(state)

	delta := int32(mutexLocked - 1<<mutexWaiterShift)

	fmt.Println(delta)
	fmt.Println("+=========+")
	fmt.Println(addInt32)
	var w sync.WaitGroup
	// var lock1 sync.Mutex
	w.Add(96)
	for i := 0; i < 96; i++ {
		go func() {
			// lock1.Lock()
			_ = atomic.AddInt32(&addInt32, -1)
			// fmt.Println(swapInt32)
			// addInt32--
			// lock1.Unlock()
			w.Done()
		}()
	}

	w.Wait()
	// time.Sleep(time.Second*10)
	fmt.Println(addInt32)

	background := context.Background()
	background = context.WithValue(background, key1, "key1")
	background = context.WithValue(background, key2, "key2")
	background = context.WithValue(background, key3, "key3")

	fmt.Println(background.Value(key1))
	fmt.Println(background.Value(key2))
	fmt.Println(background.Value(key3))

	cancel, cancelFunc := context.WithCancel(background)

	go func() {
		time.Sleep(time.Second * 5)
		cancelFunc()
	}()

	err := conn(cancel)
	fmt.Println(err)

	todo := context.TODO()
	timeout, c := context.WithTimeout(todo, time.Second*3)
	defer c()
	err = conn(timeout)
	fmt.Println(err)

	sync.Cond{}.L.Lock()
}

// conn returns a newly-opened or cached *driverConn.
func conn(ctx context.Context) error {
	fmt.Println("处理业务4s")
	time.Sleep(time.Second * 4)
	// Check if the context is expired.
	select {
	default:
		fmt.Println("default")
	case <-ctx.Done():
		fmt.Println("ctx.Done")
		return ctx.Err()
	}
	fmt.Println("conn Done")
	return nil
}

var (
	key1 = struct {
	}{}
	key2 = struct {
	}{}
	key3 = struct {
	}{}
)
