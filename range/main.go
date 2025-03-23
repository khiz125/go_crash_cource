package main

import "fmt"

func main() {
	ids := []int{1, 10, 11, 15, 23}

	// loop
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not useing index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// sum
	var sum = 0
	for _, id := range ids {
		sum += id
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// range with map
	names := map[string]string{"A": "Alice", "B": "Bob", "C": "Caty"}

	for k, v := range names {
		fmt.Printf("%s: %s\n", k, v)
	}

	for _, v := range names {
		fmt.Println("Name:", v)
	}
 
	num := 42
	str := "Hello"
	floatNum := 3.14
	boolean := true

	fmt.Printf("Integer: %d\n", num)
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Float: %f\n", floatNum)
	fmt.Printf("Boolean: %t\n", boolean)
	fmt.Printf("Value: %v\n", floatNum)
	fmt.Printf("Type: %T\n", str)

}
