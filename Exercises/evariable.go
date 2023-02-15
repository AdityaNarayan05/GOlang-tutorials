package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor = "red"
	fmt.Println("My favorite color is", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, ageInYears := 2001, 21
	fmt.Println("Born in", birthYear, "aged", ageInYears, "years")

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "A"
		lastInitial  = "J"
	)
	fmt.Println("Initials = ", firstInitial, lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("I am", ageInDays, "days old")

}
