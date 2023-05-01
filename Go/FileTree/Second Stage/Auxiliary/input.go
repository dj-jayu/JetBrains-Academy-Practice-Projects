package Auxiliary

import (
	"fmt"
)

// GetFormat get the users choice for file format
func GetFormat() string {
	var format string
	fmt.Println("Enter file format:")
	fmt.Scanln(&format)
	return format
}

// GetOrder get the users choice for the order to show the files (asc/desc)
func GetOrder() int {
	var choice int
	for choice != 1 && choice != 2 {
		fmt.Println("Size sorting options:")
		fmt.Println("1. Descending")
		fmt.Println("2. Ascending")
		fmt.Scanln(&choice)

		if choice != 1 && choice != 2 {
			fmt.Println("Wrong Option")
			fmt.Println()
		}
	}

	return choice
}

// GetDuplicatedOption gets the user choice for duplicated files checking
func GetDuplicatedOption() bool {
	var answer string
	var checkDuplicate bool
	fmt.Println()
	fmt.Println("Check for duplicates?")
	fmt.Scanln(&answer)

	for answer != "yes" && answer != "no" {
		fmt.Println()
		fmt.Println("Please, enter yes or no.")
		fmt.Println()
		fmt.Println("Check for duplicates?")
		fmt.Scanln(&answer)
	}
	if answer == "yes" {
		checkDuplicate = true
	} else if answer == "no" {
		checkDuplicate = false
	}
	return checkDuplicate
}
