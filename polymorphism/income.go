package main

import (
	"fmt"
)

type project interface {
	getIncome() int
	getName() string
}

type fixedBidding struct {
	projectName string
	bidding     int
}

type timeCharging struct {
	projectName string
	costOfHours int
	noOfHours   int
}

type advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (fb fixedBidding) getName() string {
	return fb.projectName
}

func (fb fixedBidding) getIncome() int {
	return fb.bidding
}

func (tc timeCharging) getName() string {
	return tc.projectName
}

func (tc timeCharging) getIncome() int {
	return tc.costOfHours * tc.noOfHours
}

func (a advertisement) getName() string {
	return a.adName
}

func (a advertisement) getIncome() int {
	return a.CPC * a.noOfClicks
}

func getTotalIncome(projects []project) int {
	totalIncome := 0
	for _, p := range projects {
		fmt.Printf("%s has %d income\n", p.getName(), p.getIncome())
		totalIncome += p.getIncome()
	}
	return totalIncome
}

func main() {
	project1 := fixedBidding{
		projectName: "project1",
		bidding:     10000,
	}
	project2 := fixedBidding{
		projectName: "project2",
		bidding:     20000,
	}
	project3 := timeCharging{
		projectName: "project3",
		costOfHours: 130,
		noOfHours:   3,
	}
	project4 := advertisement{
		adName:     "ad1",
		CPC:        45,
		noOfClicks: 123,
	}
	project5 := advertisement{
		adName:     "ad2",
		CPC:        1,
		noOfClicks: 56,
	}
	projects := []project{project1, project2, project3,
		project4, project5}
	totalIncome := getTotalIncome(projects)
	fmt.Printf("Total income is: %d\n", totalIncome)
}
