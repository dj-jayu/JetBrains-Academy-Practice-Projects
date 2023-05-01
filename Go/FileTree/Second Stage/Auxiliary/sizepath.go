package Auxiliary

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

// CreateFilesList read system folder and create list of file structs
func (c *FilesDataStructures) CreateFilesList(folderName, fileFormat string) {
	addPath := func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}
		if fileFormat != "" && !strings.HasSuffix(path, fileFormat) {
			return err
		}
		c.FileStructList = append(c.FileStructList, CreateFileStruct(path, info.Size()))
		return err
	}
	err := filepath.Walk(folderName, addPath)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateSizeFileMap creates a map with the size of the file as key, and the FileStruct as value
func (c *FilesDataStructures) CreateSizeFileMap() {
	for _, file := range c.FileStructList {
		c.SizeFilesMap[file.Size] = append(c.SizeFilesMap[file.Size], file)
	}
}

// CreateSizeSlice create unordered list of file sizes from list of file structs (only unique sizes)
// after, it will be ordered by the method OrderSizeSlice() in main
// Then, it will be used as a guide to access the file sizes map, and file hashes map in order
func (c *FilesDataStructures) CreateSizeSlice() {
	done := make(map[int64]bool)
	for _, file := range c.FileStructList {
		size := file.Size
		if _, ok := done[size]; ok {
			continue
		}
		done[size] = true
		c.OrderedSizeSlice = append(c.OrderedSizeSlice, size)
	}
}

// OrderSliceOfSizes order the list of file sizes according to the user choice asc / desc
func (c *FilesDataStructures) OrderSliceOfSizes(option int) {
	if option == 1 {
		sort.Slice(c.OrderedSizeSlice, func(i, j int) bool {
			return c.OrderedSizeSlice[i] > c.OrderedSizeSlice[j]
		})
	} else if option == 2 {
		sort.Slice(c.OrderedSizeSlice, func(i, j int) bool {
			return c.OrderedSizeSlice[i] < c.OrderedSizeSlice[j]
		})
	}
}

// PrintSizeInfo display information about file sizes
func (c *FilesDataStructures) PrintSizeInfo() {
	var path string
	for _, size := range c.OrderedSizeSlice {
		for i, file := range c.SizeFilesMap[size] {
			path = file.Path
			if i == 0 {
				fmt.Println()
				fmt.Printf("%d bytes\n", size)
			}
			fmt.Println(path)
		}
	}
}
