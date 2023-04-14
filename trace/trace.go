package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	Demo()
}

func Demo() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
