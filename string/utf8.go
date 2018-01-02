package main

import (
	"fmt"
	"unicode/utf8"
)

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func main() {
	word1 := "崔晓晨"
	length(word1)
}
