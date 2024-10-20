package main

import "fmt"

/**
- slice can cause a leak because slice capacity
*/

func main() {
	//s := make([]int, 3, 6) // 3 is the length, 6 is the capacity
	//
	//s = append(s, 1)
	//s = append(s, 1)
	//s = append(s, 1)
	//
	//fmt.Println(s)
	//fmt.Println("-----------------------")
	//
	//s1 := s[:3]
	//fmt.Println(s1)
	//
	//s2 := s[3:]
	//fmt.Println(s2)
	//
	//fmt.Println("-----------------------")
	//
	//s2[1] = 2
	//fmt.Println(s2)
	//fmt.Println(s1)
	//fmt.Println(s)
	//
	//fmt.Println("-----------------------")
	//
	//s2 = append(s2, 3)
	//s2 = append(s2, 3)
	//s2 = append(s2, 3)
	//
	//fmt.Println(s2)
	//fmt.Println(s1)
	//fmt.Println(s)
	//
	//fmt.Println("-----------------------")
	//
	//src := []int{0, 1, 2}
	//dst := make([]int, len(src))
	//copy(dst, src)
	//
	//fmt.Println(dst)
	//fmt.Println(src)

	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func CreateSlice(numbers []int) []int {
	value := make([]int, len(numbers))

	for i, num := range numbers {
		value[i] = num
	}

	return value
}
