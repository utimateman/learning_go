// channel
// like a pipe
// if try to receive and there is nothing, you will get blocked

// pushing side 
// buffer - each buffer have capacity until cap is filled u wont block (like bond q)
// unbuffer - get block til someone receive from the other end
package main

import(
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// This will block
	/*
	<- ch
	fmt.Println("Here")
	*/

	go func() {
		// send number of the channel
		ch <- 353
	}()

	// receive from the channel
	val := <-ch
	fmt.Printf("got %d\n", val)
	
	fmt.Println("-----")

	// send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i 
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}
	
	fmt.Println("-----")

	// close to signal when we're done
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i 
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}