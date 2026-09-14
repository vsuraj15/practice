package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process2(ch chan int) {
	for i := 0; i < 10; i++ {
		val := rand.Intn(2000)
		time.Sleep(time.Duration(val) * time.Millisecond)
		ch <- val
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go process2(ch)
	fmt.Println("Waiting for response...")
	for value := range ch {
		fmt.Printf("Process took %dms for execution\n", value)
	}
}
