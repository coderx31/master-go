package main

import "testing"

func BenchmarkKeepFirstTwoElementsOnly(b *testing.B) {
	foos := getFoos()
	for i := 0; i < b.N; i++ {
		keepFirstTwoElementsOnly(foos)
	}
}

func getFoos() []Foo {
	foos := make([]Foo, 1000)
	return foos
}
