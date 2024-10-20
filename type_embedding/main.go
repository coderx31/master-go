package main

import "sync"

// we use embedded to promote the fields and methods of an embedded type

type Foo struct {
	Bar // embedded field (due to no associated name)
}

type Bar struct {
	Baz int
}

type InMem struct {
	mu sync.Mutex
	m  map[string]int
}

func New() *InMem {
	return &InMem{m: make(map[string]int)}
}

// Deprecated: this is no longer user
func (i *InMem) Get(key string) (int, bool) {
	i.mu.Lock()
	v, contains := i.m[key]
	i.mu.Unlock()
	return v, contains
}

func main() {

	foo := Foo{}
	foo.Baz = 42

	//m := New()
	//m.Lock()

}
