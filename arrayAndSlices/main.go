package main

import "fmt"

func main() {
	// array
	var fruitArr [2]string

	// assign values
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"

	//declre and assign
	fruitArr2 := [2]string{"banana", "peach"}
    
	//slices
	fruitSlices := []string{"grape", "lemon", "mango", "cherry"}

	fmt.Println((fruitArr))
	fmt.Println((fruitArr2))
	fmt.Println(fruitSlices)

	fmt.Println(len(fruitArr))
	fmt.Println(fruitSlices[1:])
}