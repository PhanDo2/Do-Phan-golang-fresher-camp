package main

import (
	"fmt"
	"time"
)

func crawl(s []int, c chan int) {
	for i := 0; i < len(s); i++ {
		c <- s[i]
	}
}

func main() {
	a := make(chan int, 250)
	c := make(chan int, 250)
	b := make(chan int, 250)
	d := make(chan int, 250)
	s := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = i + 1
	}
	go crawl(s, a)
	go crawl(s[250:750], b)
	go crawl(s[500:750], c)
	go crawl(s[750:], d)

	fmt.Printf("\n channel a: \n")
	for i := 0; i < len(a); i++ {
		fmt.Print(<-a, "; ")
	}
	fmt.Printf("\n channel b: \n")
	for i := 0; i < len(b); i++ {
		fmt.Print(<-b, "; ")
	}
	fmt.Printf("\n channel c: \n ")
	for i := 0; i < 250; i++ {
		fmt.Print(<-c, "; ")
	}
	fmt.Printf("\n channel d: \n")
	for i := 0; i < 250; i++ {
		fmt.Print(<-d, "; ")
	}
	time.Sleep(time.Second)
}
