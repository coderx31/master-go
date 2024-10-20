package main

import (
	"log"
	"os"
)

func main() {

}

// it is not best practise to call defer inside a for loop
//func readFiles(ch <-chan string) error {
//	for path := range ch {
//		file, err := os.Open(path)
//		if err != nil {
//			return err
//		}
//		defer func() {
//			if err = file.Close(); err != nil {
//				log.Println("file close err", err)
//			}
//		}()
//
//		// doing something
//	}
//	return nil
//}

func readFiles(ch <-chan string) error {
	for path := range ch {
		err := readFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// doing something
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("file close err:", err)
		}
	}()

	return nil
}
