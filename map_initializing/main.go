package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"runtime"
)

/**
The reason is that the number of buckets in a map cannot shrink, therefore removing elements from a map doesn't impact
the number of existing buckets.
*/

/**
adding n elements to a map and then deleting all the elements means keeping same number of buckets in memory
*/

const n = 1000000

func main() {
	m := make(map[int][128]byte)
	//defer func() {
	//	m = nil
	//}()
	printAlloc()

	for i := 0; i < n; i++ {
		byteArr, err := generateRandomBytesArray()
		if err != nil {
			log.Fatal(err)
		}
		m[i] = byteArr
	}

	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()

	runtime.KeepAlive(m)

}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024*1024)
}

func generateRandomBytesArray() ([128]byte, error) {
	var arr [128]byte
	_, err := rand.Read(arr[:])
	if err != nil {
		return arr, err
	}
	// Ensure each byte is within the range 1-127
	for i := 0; i < len(arr); i++ {
		arr[i] = 1 + (arr[i] % 127)
	}
	return arr, nil
}

func CreateMap() map[string]int {
	m := make(map[string]int, 1000000)

	for i := 1; i <= 1000000; i++ {
		k := fmt.Sprintf("%d", i)
		m[k] = i
	}

	return m
}
