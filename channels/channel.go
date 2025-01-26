package main

import (
	"fmt"
	"time"
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
		i := 1
		for i < 5 {
			strChan <- fmt.Sprintf("val - %d", i)
			i++
			time.Sleep(time.Second * 2)
		}
		close(strChan)
	}()

	for v := range strChan {
		fmt.Println(v)
	}
}
