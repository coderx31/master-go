package main

import "testing"

func BenchmarkCreateMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateMap()
	}
}
