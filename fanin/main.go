package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int, name string) {
	for {
		sleep()
		//
		n := rand.Intn(100)
		fmt.Printf("ChannelName: %s and Value: %d\n", name, n)
		ch <- n
	}
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("Value consumed: %d\n", n)
	}
}

func sleep() {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func fanIn(chA, chB <-chan int, chC chan<- int) {
	var n int
	for {
		select {
		case n = <-chA:
			chC <- n
		case n = <-chB:
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}
