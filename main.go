package main

import "fmt"

// user defined function
// specifies two int parameters and the return type as int
// this is type restricting the return parameters as well as the type
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello Valerie, this is a test")

	fmt.Println("10 + 10 =", 10+10)

	str1 := "Welcome"

	str2 := "Valerie"

	result := str1 + str2

	fmt.Println(result)

	userResult := add(5, 3)

	fmt.Println("Sum: ", userResult)

}
