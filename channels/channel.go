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
			//time.Sleep(time.Second * 2)
		}
		close(strChan)
	}()

	for v := range strChan {
		fmt.Println(v)
	}
	fmt.Println("------------------------------------------------------------")
	values := []int{345, 243, 454, 76, 832, 233}
	fmt.Println(sleepSort(values))
}

func sleepSort(values []int) []int {
	vChannel := make(chan int)
	for _, v := range values {
		v := v
		go func() {
			time.Sleep(time.Millisecond * time.Duration(v))
			vChannel <- v
		}()
	}

	s := make([]int, 0, len(values))
	for range values {
		n := <-vChannel
		s = append(s, n)
	}

	return s
}
