package main

import "fmt"

// Both a and b are int
func multiply(a, b int) int {
	return a * b
}

func main() {

	fmt.Println("Hello World!")

	const name string = "Youngkwon"
	age := 25 // Only possible inside of a func
	// var age int = 25

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(multiply(2, 3))

}