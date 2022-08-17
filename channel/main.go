package main

import (
	"context"
	"fmt"
	"time"
)

func f() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer fmt.Println("[f] defer")
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[gen] Done")
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	f()
	time.Sleep(3 * time.Second)
	fmt.Println("[main] done")
}
