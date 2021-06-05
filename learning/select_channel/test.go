package main

import (
	"fmt"
	"time"
)

func main() {

	var dataCh = make(chan int)
	var doneCh = make(chan struct{})
	go sendData(dataCh, doneCh)

forLoop:
	for {
		select {
		case v := <-dataCh:
			fmt.Println(v)

		case <-time.After(1 * time.Second):
			fmt.Println("timeout:")

		case <-doneCh:
			fmt.Println("Done select")
			break forLoop

		}
	}

	var myMap = make(map[string]struct{})
	myMap["anand"] = struct{}{}

	if _, ok := myMap["anand"]; ok {
		fmt.Println("exist")
	}

	if _, ok := myMap["ananda"]; !ok {
		fmt.Println(" noot exist")
	}
}

func sendData(dataCh chan int, doneCh chan struct{}) {

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {

		dataCh <- i
	}
	fmt.Println("Done")
	doneCh <- struct{}{}
}
