package main

import (
	"fmt"
)

func assert(i interface{}) {
	v, ok := i.(string)
	fmt.Println(v, ok)
}

func main() {
	var s interface{} = "Xiaochen Cui"
	assert(s)
}
