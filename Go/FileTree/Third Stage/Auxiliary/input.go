package Auxiliary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getYesOrNoOption(question, errorMsg string) bool {
	var answer string
	var yes bool
	fmt.Println()
	fmt.Println(question)
	fmt.Scanln(&answer)

	for answer != "yes" && answer != "no" {
		fmt.Println()
		fmt.Println(errorMsg)
		fmt.Println()
		fmt.Println(question)
		fmt.Scanln(&answer)
	}
	if answer == "yes" {
		yes = true
	} else if answer == "no" {
		yes = false
	}
	return yes
}
func GetDeleteOption() bool {
	return getYesOrNoOption("Delete files?", "Please, enter yes or no.")
}

// auxiliary function for GetNumbersToDelete
// validate the input by checking if they are ints
// and if they are <= the numbers of repeated files
func checkInvalidAnswer(splitAnswer []string, maxRange int) bool {
	for _, numberString := range splitAnswer {
		intChoice, err := strconv.Atoi(numberString)
		if err != nil {
			return true
		}
		if intChoice > maxRange {
			return true
		}
	}
	return false
}

func GetNumbersToDelete(maxRange int) []int {
	var answer string
	var splitAnswer []string
	var invalidAnswer bool
	var indexOfFilesToDelete []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter file numbers to delete:")
	scanner.Scan()
	answer = scanner.Text()
	splitAnswer = strings.Split(answer, " ")
	invalidAnswer = checkInvalidAnswer(splitAnswer, maxRange)
	for invalidAnswer {
		fmt.Println("Wrong format")
		fmt.Println("Enter file numbers to delete:")
		scanner.Scan()
		answer = scanner.Text()
		splitAnswer = strings.Split(answer, " ")
		invalidAnswer = checkInvalidAnswer(splitAnswer, maxRange)
	}
	for _, indexString := range splitAnswer {
		indexInt, _ := strconv.Atoi(indexString)
		indexOfFilesToDelete = append(indexOfFilesToDelete, indexInt)
	}
	return indexOfFilesToDelete
}

func (c *FilesDataStructures) GetPathsToDelete(indexOfFilesToDelete []int) []string {
	var pathsToDelete []string
	for _, duplicatedNumber := range indexOfFilesToDelete {
		file, err := c.GetFileByDuplicatedNumber(duplicatedNumber)
		if err != nil {
			log.Fatal(err)
		}
		pathsToDelete = append(pathsToDelete, file.Path)
	}
	return pathsToDelete
}

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
	return getYesOrNoOption("Check for duplicates?", "Please, enter yes or no.")
}
