package main

import (
	"fmt"
	"time"
)

func main() {
	//Wait Groups + Go Routines
	//=========================
	//var waitGroup sync.WaitGroup
	//waitGroup.Add(1)

	//Anonymous function
	//go func() {
		//countThings("Sheep")
		//Decrements counter by 1
		//waitGroup.Done()
	//} ()

	//Waits until go routine terminates
	//waitGroup.Wait()

	//Channels
	//========
	//channel := make(chan string)
	//go countThings("Sheep", channel)

	//for {
		//Receive message from channel
		//message, open := <- channel
		//if !open {
		//	break
		//}
	//}

	//Syntactic sugar ^
	//for message := range channel{
		//fmt.Println(message)
	//}

	//Buffered Channels
	//=================
	//Send blocks until something is ready to receive
	//Need to receive in separate go routine or make buffered channel
	// - provide a capacity when creating the channel
	//Can send 2 strings to channel before receiving
	//channel := make(chan string, 2)
	//channel <- "hello"
	//channel <- "world"
	//message := <- channel
	//fmt.Println(message)
	//message = <- channel
	//fmt.Println(message)

	//Select Statements
	//=================
	//channel1 := make(chan string)
	//channel2 := make(chan string)
	//
	//go func() {
	//	for {
	//		channel1 <- "Every half a second"
	//		time.Sleep(time.Millisecond * 500)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		channel2 <- "Every 2 seconds"
	//		time.Sleep(time.Second * 2)
	//	}
	//}()
	//
	//for {
	//	//Can receive from a ready channel
	//	select {
	//	case message1 := <- channel1:
	//		fmt.Println(message1)
	//	case message2 := <- channel2:
	//		fmt.Println(message2)
	//	}
	//}

	//Worker Pools
	//============
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<- results)
	}
}

func countThings(thing string, channel chan string) {
	for i := 1; i <= 5; i++ {
		//Send message through channel
		channel <- thing
		time.Sleep(time.Millisecond * 500)
	}

	//Avoids deadlock (main function still waiting to receive on channel)
	//Only close channel if acting as sender
	// - as receiver can prematurely close channel while data still being sent
	close(channel)
}

//Only receive from jobs channel and send on results channel
func worker(jobs <- chan int, results chan <- int) {
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n -1) + fibonacci(n - 2)
}
