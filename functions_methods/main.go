package main

import "fmt"

/**
- when to use value or pointer receivers
- when to use named result parameters and their potential side effects
- avoiding a common mistake while returning a nil receiver
- why using functions that accept a filename isn't best practise
- handling defer arguments
*/

/**
 - with a value receiver, Go makes a copy of the value and passes it to the method (any changes make to object remain local to method)
	(original object remain unchanged)
 - with a pointer receiver, Go passes the address of an object to the method, it remains a copy, but only a copy of pointer
	not the object itself. (passing by reference doesn't exist in Go)


 - a receiver must be pointer if the method need to mutate the receiver
 - a receiver should bea pointer if the receiver is large object

 - a receiver must be a value,
	- if we have to enforce a receiver's immutability
	- if the receiver is a map, function or a channel (otherwise a compilation error)

 - a receiver should be a value
	- if the receiver is a slice that doesn't have to a mutated
	- if the receiver is a small array or struct that is naturally a value type without mutable fields (such as time)
	- if the receiver is a basic type such as int, float64, or string
*/

/**
 - when to use named result parameters depends on the context. In most cases, if it's not clear whether using
	them makes our code more readable, we shouldn't use named result parameters.
*/

type customer struct {
	balance float64
}

func (c customer) add(v float64) {
	c.balance += v
}

func main() {
	c := customer{balance: 100}
	c.add(50)
	fmt.Printf("balance: %.2f\n", c.balance)
}
