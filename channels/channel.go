package main

import (
	"fmt"
)

func main() {
	go fmt.Println("go routing definition") //Spinning up goroutine with named function

	for i := 0; i < 5; i++ {
		i := i      //Shadowing the i before using by goroutine
		go func() { //anonymous functioin go routine
			fmt.Println(i)
		}()
	}

	strChan := make(chan string) //create a channel with a type
	go func() {
		for i := 0; i < 5; i++ {
			strChan <- fmt.Sprintf("val%d", i+1)
		}
		close(strChan)
	}()
 
	for v := range strChan {
		fmt.Println(v)
	}
}
