package main

import (
	"fmt"
	"time"
)

func increment(x *int) {
	for i := 0; i < 10; i++ {
		*x++
		time.Sleep(10 * time.Millisecond)
	}
}

func print(x *int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("x=%v\n", *x)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	x := 0
	go increment(&x)
	go print(&x)
	time.Sleep(100 * time.Millisecond)
}
