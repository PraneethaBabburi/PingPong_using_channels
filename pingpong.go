package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)
	var wg sync.WaitGroup
	iterations := 5
	wg.Add(2)
	go func(iterations int) {
		for i := 0; i < iterations; i++ {
			pingmsg := <-ping
			fmt.Println(pingmsg)
			pong <- "pong"
		}
		wg.Done()
	}(iterations)

	go func(iterations int) {

		for i := 0; i < iterations; i++ {
			pongmsg := <-pong
			fmt.Println(pongmsg)
			if i < iterations-1 {
				ping <- "ping"
			}
		}

		wg.Done()
	}(iterations)


	time.Sleep(time.Second)
	ping <- "ping"
	
	wg.Wait()

}
