package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(div(7, 1))
}

func div(a, b int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	log.Println("Before panicking.")
	result = a / b
	log.Println("After panicking.")
	return result, err
}
