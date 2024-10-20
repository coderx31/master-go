package main

import (
	"fmt"
)

type customer struct {
	id         string
	operations []float64
}

func (a customer) equal(b customer) bool {
	if a.id != b.id {
		return false
	}

	if len(a.operations) != len(b.operations) {
		return false
	}

	for i := 0; i < len(a.operations); i++ {
		if a.operations[i] != b.operations[i] {
			return false
		}
	}
	return true
}

func main() {
	cust1 := customer{id: "x", operations: []float64{1.}}
	cust2 := customer{id: "x", operations: []float64{1.}}
	//fmt.Println(reflect.DeepEqual(cust1, cust2))
	fmt.Println(cust1.equal(cust2))
}
