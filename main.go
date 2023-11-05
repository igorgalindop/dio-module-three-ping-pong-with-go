package main

import (
	"fmt"
	"time"
)

func ping(c1 chan string) {
	for i := 0; ; i++ {
		c1 <- "ping"
	}
}

func pong(c2 chan string) {
	for i := 0; ; i++ {
		c2 <- "pong"
	}
}

func printPingPong(c1 chan string, c2 chan string) {
	for {
		msgChan1 := <-c1
		msgChan2 := <-c2
		fmt.Println(msgChan1)
		time.Sleep(time.Second * 1)
		fmt.Println(msgChan2)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go ping(c1)
	go pong(c2)
	go printPingPong(c1, c2)

	var stop string
	fmt.Scanln(&stop)

}
