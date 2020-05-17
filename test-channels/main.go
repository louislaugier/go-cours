package main

import (
	"fmt"
	"time"
)

func sendValue(c chan string) {
	fmt.Println("executing goroutine")
	time.Sleep(1 * time.Second)
	c <- "hello w"
	fmt.Println("done executing goroutine")
}

func main() {
	values := make(chan string, 2)
	defer close(values)
	go sendValue(values)
	go sendValue(values)
	value := <-values
	fmt.Println(value)
	time.Sleep(1 * time.Second)
}
