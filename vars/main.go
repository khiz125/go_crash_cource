package main
import "fmt"

func main() {
	// using var
	var name string = "khiz"
	var age int32 = 40
	
	// using const
	const isCool = true

	// shorthand
	hello := "hello"
	size := 1.3

	fmt.Println(name, age, isCool, hello, size)
	fmt.Printf("%T\n", isCool)


}