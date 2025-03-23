package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("Type: %T\n", b)

	fmt.Println(*b)

	*b = 10
	fmt.Println(a, b)
}