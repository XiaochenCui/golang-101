package main

import (
	"fmt"
)

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a befor deferred function call", a)
}
