package main

import (
	"fmt"
	"time"
)

func crawl(ch chan int, x int) {
	for i := range ch {
		fmt.Printf(" url %d crawl %d \n", x, i)
	}
}
func main() {
	url := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go crawl(url, i)
	}
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 100)
		url <- i
	}
	fmt.Println("ok", cap(url))
}
