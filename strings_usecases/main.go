package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello"
	fmt.Println(utf8.RuneCountInString(s))
}

//func concat(values []string) string {
//	s := ""
//	for _, value := range values {
//		s += value
//	}
//	return s
//}

//func concat(values []string) string {
//	sb := strings.Builder{}
//	for _, value := range values {
//		_, _ = sb.WriteString(value)
//	}
//	return sb.String()
//}

//func concat(values []string) string {
//	total := 0
//	for i := 0; i < len(values); i++ {
//		total += len(values[i])
//	}
//
//	sb := strings.Builder{}
//	sb.Grow(total)
//
//	for _, value := range values {
//		_, _ = sb.WriteString(value)
//	}
//	return sb.String()
//}

func sanitize(b []byte) []byte {
	return bytes.TrimSpace(b)
}
