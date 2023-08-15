package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	chan1 = make(chan struct{}, 1)
	// chan2 = make(chan struct{} ,1)
)

func main() {
	runtime.GOMAXPROCS(runtime.GOMAXPROCS(runtime.NumCPU()))
	i := int64(12132313)
	go test(i)
	fmt.Println(runtime.NumGoroutine())

	ticker := time.NewTicker(time.Second * 6)
	select {
	case <-chan1:
		fmt.Println("test结束")
	case <-ticker.C:
		fmt.Println("定时器")
	}

	fmt.Println("会阻塞吗?")

	var smap sync.Map
	smap.Store("ni", "hao")
	smap.Store("ni1", "s")
	s, _ := smap.Load("ni")
	s, _ = smap.Load("ni")
	smap.Delete("ni")
	smap.Store("ni2", "s")
	s, _ = smap.Load("ni1")

	s2 := s.(string)
	fmt.Println(s2)
	// e := getNil()
	var value atomic.Value
	e, _ := value.Load().(ex)
	fmt.Println(e)
	var mymap MyMap
	mymap.Store("sd", "s")

}

type ex struct {
	name string
}

func getNil() ex {
	value := new(atomic.Value)
	e := value.Load().(ex)
	return e
}

// /can only use //go:noescape with external func implementations
// /go:noescape
func test(arg int64) {
	go func() {
		fmt.Println("+")
	}()
	abc(arg)
}

func abc(int642 int64) {
	time.Sleep(time.Second * 5)
	sprintf := fmt.Sprintf("int:%d", int642)
	fmt.Println(sprintf)
	fmt.Println(unsafe.Sizeof(int642))
	chan1 <- struct{}{}

}

type MyMap struct {
	mu sync.Mutex

	// read contains the portion of the map's contents that are safe for
	// concurrent access (with or without mu held).
	//
	// The read field itself is always safe to load, but must only be stored with
	// mu held.
	//
	// Entries stored in read may be updated concurrently without mu, but updating
	// a previously-expunged entry requires that the entry be copied to the dirty
	// map and unexpunged with mu held.
	read atomic.Value // readOnly

	// dirty contains the portion of the map's contents that require mu to be
	// held. To ensure that the dirty map can be promoted to the read map quickly,
	// it also includes all of the non-expunged entries in the read map.
	//
	// Expunged entries are not stored in the dirty map. An expunged entry in the
	// clean map must be unexpunged and added to the dirty map before a new value
	// can be stored to it.
	//
	// If the dirty map is nil, the next write to the map will initialize it by
	// making a shallow copy of the clean map, omitting stale entries.
	dirty map[any]*Myentry

	// misses counts the number of loads since the read map was last updated that
	// needed to lock mu to determine whether the key was present.
	//
	// Once enough misses have occurred to cover the cost of copying the dirty
	// map, the dirty map will be promoted to the read map (in the unamended
	// state) and the next store to the map will make a new dirty copy.
	misses int
}

type MyreadOnly struct {
	m       map[any]*Myentry
	amended bool // true if the dirty map contains some key not in m.
}

var Myexpunged = unsafe.Pointer(new(any))

type Myentry struct {
	// p points to the interface{} value stored for the entry.
	//
	// If p == nil, the entry has been deleted, and either m.dirty == nil or
	// m.dirty[key] is e.
	//
	// If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
	// is missing from m.dirty.
	//
	// Otherwise, the entry is valid and recorded in m.read.m[key] and, if m.dirty
	// != nil, in m.dirty[key].
	//
	// An entry can be deleted by atomic replacement with nil: when m.dirty is
	// next created, it will atomically replace nil with expunged and leave
	// m.dirty[key] unset.
	//
	// An entry's associated value can be updated by atomic replacement, provided
	// p != expunged. If p == expunged, an entry's associated value can be updated
	// only after first setting m.dirty[key] = e so that lookups using the dirty
	// map find the entry.
	p unsafe.Pointer // *interface{}
}

func (m *MyMap) Store(key, value any) {
	// readtmp := m.read.Load()
	// read := readtmp.(MyreadOnly) 这样的写法会报以下的错
	read, _ := m.read.Load().(MyreadOnly) // 这样就没事
	if e, ok := read.m[key]; ok && e.tryStore(&value) {
		return
	}

	m.mu.Lock()
	read, _ = m.read.Load().(MyreadOnly)
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			// The entry was previously expunged, which implies that there is a
			// non-nil dirty map and this entry is not in it.
			m.dirty[key] = e
		}
		e.storeLocked(&value)
	} else if e, ok := m.dirty[key]; ok {
		e.storeLocked(&value)
	} else {
		if !read.amended {
			// We're adding the first new key to the dirty map.
			// Make sure it is allocated and mark the read-only map as incomplete.
			m.dirtyLocked()
			m.read.Store(MyreadOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
	}
	m.mu.Unlock()
}

// storeLocked unconditionally stores a value to the entry.
//
// The entry must be known not to be expunged.
func (e *Myentry) storeLocked(i *any) {
	atomic.StorePointer(&e.p, unsafe.Pointer(i))
}

func (e *Myentry) tryStore(i *any) bool {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == Myexpunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
	}
}

func (e *Myentry) unexpungeLocked() (wasExpunged bool) {
	return atomic.CompareAndSwapPointer(&e.p, Myexpunged, nil)
}

func newEntry(i any) *Myentry {
	return &Myentry{p: unsafe.Pointer(&i)}
}

func (m *MyMap) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read, _ := m.read.Load().(MyreadOnly)
	m.dirty = make(map[any]*Myentry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}

func (e *Myentry) tryExpungeLocked() (isExpunged bool) {
	p := atomic.LoadPointer(&e.p)
	for p == nil {
		if atomic.CompareAndSwapPointer(&e.p, nil, Myexpunged) {
			return true
		}
		p = atomic.LoadPointer(&e.p)
	}
	return p == Myexpunged
}
