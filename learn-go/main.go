package main

import (
	"fmt"
	"strings"
)

// Both a and b are int
func multiply(a, b int) int {
	return a * b
}

// Return multiple result
func lenAndUpper(name string) (int, string) {
	defer fmt.Println("lenAndUpper is done..!")
	return len(name), strings.ToUpper(name)
}

func lenAndUpperNaked(name string) (length int, upper string) {
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	// for i:= 0; i < len(numbers); i++ { ... } is also possible
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge >= 20 {
		return true
	}
	return false
}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 20:
		return true
	}
	return false
}

func arrayAndSlice() {
	// Array
	names := [4]string{"youngkwon", "kim"}
	names[2] = "Jake"
	names[3] = "Alden"
	fmt.Println("names", names)

	// Slice
	nameSlice := []string{"youngkwon", "kim"}
	nameSlice = append(nameSlice, "Jake")
	fmt.Println("nameSlice", nameSlice)
}

func mapLikeObject() {
	person := map[string]string{"name": "youngkwon", "address": "Seoul, Korea"}
	fmt.Println(person)
}

func main() {

	fmt.Println("Hello World!")

	const name string = "Youngkwon"
	age := 25 // Only possible inside of a func
	// var age int = 25

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(multiply(2, 3))

	totalLength, upperName := lenAndUpper("Youngkwon")
	fmt.Println(totalLength, upperName)

	totalLength, upperName = lenAndUpperNaked("Jake Alden")
	fmt.Println(totalLength, upperName)

	repeatMe("youngkwon", "kim", "jake", "alden")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkSwitch(18))

	arrayAndSlice()
	mapLikeObject()
}
