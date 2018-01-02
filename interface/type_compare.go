package main

import "fmt"

type Describer interface {
	Describer()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describer() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describer()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("xiaochen")
	p := Person{
		name: "xiaochen cui",
		age:  25,
	}
	findType(p)
}
