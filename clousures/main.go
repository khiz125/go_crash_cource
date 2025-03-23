package main

import "fmt"

func counter() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := counter()
	
	for i := range 10 {
		fmt.Println(sum(i))
	}
}