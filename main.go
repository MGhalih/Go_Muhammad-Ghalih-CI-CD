package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("simple math")

	var a, b int

	fmt.Print("enter a first number: ")

	fmt.Scan(&a)

	fmt.Print("enter a second number: ")

	fmt.Scan(&b)

	fmt.Println("result: ", Add(a, b))
}
