package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c) // takes turns printing ping/pong

	fmt.Scanln(new(string))

	// c := make(chan int, 1) //buffered channel capacity 1
}

func pinger(c chan<- string) { // chan<- means c can only be sent to
	for i := 0; ; i++ { // 1 type of infinite loop
		c <- "ping" // "blocks" until printer is ready to receive a message
	}
}

func printer(c <-chan string) { // <-chan means c can only receive
	for { // another type of infinite loop
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}
