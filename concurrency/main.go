package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process2(ch chan int) {
	val := rand.Intn(2000)
	time.Sleep(time.Duration(val) * time.Millisecond)
	ch <- val
}

func main() {
	ch := make(chan int)
	go process2(ch)
	fmt.Println("Waiting for response...")
	value := <-ch
	fmt.Printf("Process took %dms for execution\n", value)
}
