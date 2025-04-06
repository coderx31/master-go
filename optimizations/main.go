package main

/**
- make it correct, make it clear, make it concise, make it fast, in that order.
*/

/**
- as a Go developers, one avenue for improvement is making sure our applications use CPU caches
*/

/**
  - slice of structs vs struct of slice
  - slice of structs is a better solution
*/

/**
- data alignment
- byte, uint8, int8: 1 byte
- uint16, int16: 2 bytes
- uint32, int32, float32: 4 bytes
- uint64, int64, float46: 8 bytes
- complex128: 16 bytes
*/

/**
 - how we can reduce the amount of memory allocated ?
 - the rule of thumb is to reorganize a struct so that its fields are sorted by type size
	in descending order
*/

/**
- profiling

go test -bench=. -v -trace=trace.out
*/

type Foo struct {
	a int64
	b int64
}

func sumFoo(foos []Foo) int64 {
	var total int64
	for i := 0; i < len(foos); i++ {
		total += foos[i].a
	}
	return total
}

type Bar struct {
	a []int64
	b []int64
}

func sumBar(bar Bar) int64 {
	var total int64
	for i := 0; i < len(bar.a); i++ {
		total += bar.a[i]
	}
	return total
}

func main() {

}
