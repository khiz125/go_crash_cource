package main

import (
	"fmt"
	"strconv"
)

// define person staruct

type Person struct {
	name   string
	age    int
	gender string
}

// greeting method
func (p Person) greet() string {
	return "Hello, my name is " + p.name + ", and I am " + 
	strconv.Itoa(p.age) + " years old."
}

// hasBirthday method
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method
func (p *Person) getMarried(spouseName string) {
	if p.gender == "m" {
		return
	} else {
		p.name = spouseName
	}
}



func main() {

	// init person using struct
	personA := Person{name: "dazai", age: 35, gender: "m"}
	fmt.Println(personA)
	fmt.Println(personA.name)

	personB := Person{"mishima", 45, "m"}
	personC := Person{"masaoka", 25, "f"}
	fmt.Println(personB)
	fmt.Println(personB.age)
	personB.age++
	fmt.Println(personB.age)

	fmt.Println(personA.greet())
	personA.hasBirthday()
	fmt.Println(personA.greet())

	fmt.Println(personC.greet())
	personC.getMarried("yamada")
	fmt.Println(personC.greet())
}
