package main

import (
	"fmt"
	"geometry/rectangle" //import custom package
	"log"
)

/*
 * 1. package variables
 */
var reacLen, reacWidth float64 = -6, 7

/*
 * 2. init function to check if length and width are greater than zero
 */
func init() {
	println("main package initialized")
	if reacLen < 0 {
		log.Fatal("length is less than zero")
	}
	if reacWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(reacLen, reacWidth))
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(reacLen, reacWidth))
}
