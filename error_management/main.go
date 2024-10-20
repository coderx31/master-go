package main

import "fmt"

/**
- expected errors should be designed as error values(sentinel errors) var errFoo = errors.New("foo")
- unexpected errors should be designed as error types BarErr struct {} with implementing error interface
*/

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r) // it doesn't stop the go routine
		}
	}()
	fmt.Println("a")
	panic("foo")
	fmt.Println("b")

	// errors.As
	// errors.Is
}
