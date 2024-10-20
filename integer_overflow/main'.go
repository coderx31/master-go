package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int32 = math.MaxInt32
	fmt.Println(counter)
	counter++
	fmt.Printf("counter: %d\n", counter)
}
