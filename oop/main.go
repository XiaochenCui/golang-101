package main

import "oop/employee"

func main() {
	e := employee.New("Xiaochen", "Cui", 30, 10)
	e.LeavesRemaining()
}
