package main

import (
	"fmt"
	"runtime"
)

/**
if the element is a pointer or a struct with pointer fields, the elements won't be reclaimed by the GC.
in this example, slice is a pointer to backing array.
*/

type Foo struct {
	v []byte
}

func main() {
	foos := make([]Foo, 1000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

//func keepFirstTwoElementsOnly(foos []Foo) []Foo {
//	return foos[:2]
//}

/*
*
Because we copy the first two elements of the slice, the GC knows that the 998 elements won't be referenced anymore
and can now be collected
*/
func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	res := make([]Foo, 2)
	copy(res, foos)
	return res
}

//func keepFirstTwoElementsOnly(foos []Foo) []Foo {
//	for i := 2; i < len(foos); i++ {
//		foos[i].v = nil
//	}
//	return foos[:2]
//}
