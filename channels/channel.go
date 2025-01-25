package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("go routing definition") //Spinning up goroutine with named function
	fmt.Println("main")

	for i := 0; i < 5; i++ {
		i := i //Shadowing the i before using by goroutine
		go func () {			//anonymous functioin go routine
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second*2)
}