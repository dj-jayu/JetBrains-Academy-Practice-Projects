package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Directory is not specified")
	} else {
		folderName := os.Args[1]
		fileFormat := GetFormat()
		order := GetOrder()
		filepath.Walk(folderName, AddPath)
		Container.Order(order)
		Container.PrintInfo(fileFormat)
	}
}
