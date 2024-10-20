package main

import "fmt"

/**
In summary, when we call defer on a function or method, the call’s
arguments are evaluated immediately. If we want to mutate the arguments
provided to defer afterward, we can use pointers or closures. For a
method, the receiver is also evaluated immediately; hence, the behavior
depends on whether the receiver is a value or a pointer.
*/

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

func f() error {
	var status string
	defer notify(&status)
	defer incrementCounter(&status)

	status = StatusSuccess

	return nil
}

func notify(status *string) {
	if *status == "" {
		fmt.Println("empty string")
	}
	fmt.Println(*status)
}

func incrementCounter(status *string) {
	if *status == "" {
		fmt.Println("empty string")
	}
	fmt.Println(*status)
}

func errOccurred() error {
	var errStatus error
	defer errorPrint(&errStatus)

	errStatus = fmt.Errorf("test error")
	return nil
}

func errorPrint(err *error) {
	if err != nil {
		fmt.Println(*err)
	} else {
		fmt.Println("error is nil")
	}
}

func main() {
	//err := f()
	//if err != nil {
	//	fmt.Println("error")
	//}

	//err := errOccurred()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//i := 0
	//j := 0
	//
	//defer func(i int) {
	//	fmt.Println(i, j)
	//}(i)
	//i++
	//j++

	var err1 error
	var err2 error

	defer func() {
		fmt.Println(err1)
		fmt.Println(err2)
	}()

	err1 = fmt.Errorf("error 1")
	err2 = fmt.Errorf("error 2")
}
