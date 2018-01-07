package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, error := filepath.Glob("*.go")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}
