package main

// Concurrency in Go tutorial from: https://www.youtube.com/watch?v=LvgVSSpwND8

import (
	"fmt"
	"sync"
	"time"
)

func test() {
	fmt.Println("Concurrency in Go")
	// basic()
	// waitGroup()
	// channels()
	// buffered(3)
	// selecter()
	queue(6) // Can specify number of workers in queue with argument
}

// Part 6: Worker Queue
func queue(num int) {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < num; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// Part 5: Select
func selecter() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Two Go Routines that send a message every 500 ms and 2 seconds
	go func() {
		for {
			c1 <- "Every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c1 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// Part 4: Buffered Channel
func buffered(buffer int) {
	c := make(chan string, buffer)
	c <- "hello"
	for i := 0; i < buffer; i++ {
		c <- "hello"
		msg := <-c
		fmt.Println(msg)
	}

}

// Part 3: Channels
func channels() {
	c := make(chan string)
	go channelCount("sheep", c)

	// Simpler way to accept messages from channel
	for msg := range c {
		fmt.Println(msg)
	}

	// Longer way to accept message from channel
	// // Use for loop to keep accepting messages from channel
	// for {
	// 	// Send blocking call to let process know to wait for msg on the channel
	// 	msg, open := <-c

	// 	// Check if channel is still open to prevent deadlock
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
}

func channelCount(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	// Must close channel to prevent deadlock
	close(c)
}

// Part 2: Wait Groups
func waitGroup() {
	// Must use wait groups to sync
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("gish")
		wg.Done()
	}()
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

// Part 1: Basic
func basic() {
	// Using go keyword to run function in 'go' routine on another core.
	go infinteCount("sheep")
	infinteCount("gish")
	fmt.Scanln()
}

func infinteCount(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
