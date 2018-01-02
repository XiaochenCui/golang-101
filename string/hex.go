package main

import (
	"fmt"
)

func printBytes(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	fmt.Printf("\n")
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starys at byte %d\n", rune, index)
	}
}

func main() {
	byteSlice := []byte{0xe5, 0xb4, 0x94, 0xe6, 0x99, 0x93}
	str := string(byteSlice)
	fmt.Println(str)
	runeSlice := []rune{0xe5b494}
	str = string(runeSlice)
	fmt.Println(str)
}
