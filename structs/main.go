package main

import "fmt"

// define person staruct

type Person struct {
	name string
	age int
	gender string
}

func main() {
	
	// init person using struct
	personA := Person{name: "dazai", age: 35, gender: "m"}
	fmt.Println(personA)
	fmt.Println(personA.name)

	personB := Person{"mishima", 45, "m"}
	fmt.Println(personB)
	fmt.Println(personB.age)
	personB.age++
	fmt.Println(personB.age)
}