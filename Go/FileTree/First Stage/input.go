package main

import "fmt"

func GetFormat() string {
	var format string
	fmt.Println("Enter file format:")
	fmt.Scanln(&format)
	return format
}

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
