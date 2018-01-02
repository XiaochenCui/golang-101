package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Printf("Type of a is %T, and value is %v\n", a, a)
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T, and value is %v\n", a, a)
	}
}
