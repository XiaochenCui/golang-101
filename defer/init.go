package main

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Printf("Largest number in %#v is %d\n", nums, max)
}

func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}
