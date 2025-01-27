package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func ()  {
		time.Sleep(1*time.Second)
		ch1 <- 1	
	}()

	go func ()  {
		time.Sleep(1*time.Second)
		ch2 <- 2	
	}()

	go func ()  {
		time.Sleep(1*time.Second)
		ch3 <- 3	
	}()

	select {
	case <- ch1:
		fmt.Println("Got from channel 1")
	case <- ch2:
		fmt.Println("Got from channel 2")
	case <- ch3:
		fmt.Println("Got from channel 3")
	case <- time.After((time.Second*2)):
		fmt.Println("Cancel the remaining go routines.")
	}
}