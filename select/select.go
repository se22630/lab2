package main

import (
	"fmt"
	"time"
)

// slowSender sends a string every 2 seconds.
func slowSender(c chan<- string) {
	for {
		time.Sleep(2 * time.Second)
		c <- "I am the slowSender"
	}
}

// fastSender sends consecutive ints every 500 ms.
func fastSender(c chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		c <- i
	}
}
func fasterSender(c chan<- []int) {
	slic := []int{1, 2, 3}
	for {
		time.Sleep(200 * time.Millisecond)

		c <- slic
	}
}

// main starts the two senders and then goes into an infinite loop of receiving their messages.
func main() {
	ints := make(chan int)
	go fastSender(ints)
	strings := make(chan string)
	go slowSender(strings)
	slices := make(chan []int)
	go fasterSender(slices)

	for { // = while(true)
		select {
		case s := <-strings:
			fmt.Println("Recdeived a string", s)
		case i := <-ints:
			fmt.Println("Received an int", i)
		case j := <-slices:
			fmt.Println("Received a slice", j)
		}
	}
}
