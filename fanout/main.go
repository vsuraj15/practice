package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int) {
	for {
		sleep()
		n := rand.Intn(100)
		fmt.Printf("Producing: %d\n", n)
		ch <- n
	}
}

func consumer(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("ConsumerName: %s and value: %d\n", name, n)
	}
}

func fanOut(chA <-chan int, chB, chC chan<- int) {
	for value := range chA {
		if value < 50 {
			chB <- value
		} else {
			chC <- value
		}
	}
}

func sleep() {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA)
	go consumer(chB, "B")
	go consumer(chC, "C")

	fanOut(chA, chB, chC)
}
