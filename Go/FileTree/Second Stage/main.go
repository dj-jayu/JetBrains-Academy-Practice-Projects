package main

import (
	"fileProject/Auxiliary"
	"fmt"
	"os"
)

// Container create new struct to hold all data structures relating to the files
var Container Auxiliary.FilesDataStructures = Auxiliary.NewFilesDataStructures()

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Directory is not specified")
	} else {
		folderName := os.Args[1]
		fileFormat := Auxiliary.GetFormat()
		order := Auxiliary.GetOrder()
		Container.CreateFilesList(folderName, fileFormat)
		Container.CreateSizeFileMap()
		Container.CreateSizeSlice()
		Container.OrderSliceOfSizes(order)
		Container.PrintSizeInfo()
		checkDuplicates := Auxiliary.GetDuplicatedOption()
		if checkDuplicates {
			Container.CreateSizeHashesFilesMapMap()
			Container.PrintDuplicatesInfo()
		}
	}
}
