package main

import "fmt"

func square(x int) int {
	return x * x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from Go!")
}

func main() {
	greet()

	dozen := square(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(square(6), 1)
	fmt.Println("Have another", anotherBakersDozen)
}
