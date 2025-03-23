package main

import "fmt"

func main() {
	emails := make(map[string]string)

	// Assign key, value
	emails["bob"] = "bob@example.com"
	emails["dazai"] = "dazai@egample.com"

	fmt.Println((emails))
	fmt.Println((len(emails)))
	fmt.Println(emails["dazai"])

	// delete from map
	delete(emails, "bob")
	fmt.Println((emails))

	names := map[string]string{"A": "Alice", "B": "Bob"}
	fmt.Println((names))
	fmt.Println((len(names)))
	fmt.Println(names["A"])
}
