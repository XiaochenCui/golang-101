package main

import "fmt"

func sendData(search chan<- int) {
	search <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}
