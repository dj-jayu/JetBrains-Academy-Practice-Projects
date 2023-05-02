package Auxiliary

import (
	"fmt"
	"log"
	"math"
	"os"
)

func (c *FilesDataStructures) GetNumberOfRepeatedFiles() int {
	var max float64
	for _, file := range c.FileStructList {
		max = math.Max(float64(file.DuplicatedNumber), max)
	}
	return int(max)
}

func (c *FilesDataStructures) GetFileByDuplicatedNumber(duplicatedNumber int) (*FileStruct, error) {
	for _, file := range c.FileStructList {
		if file.DuplicatedNumber == duplicatedNumber {
			return &file, nil
		}
	}
	return &FileStruct{}, fmt.Errorf("no file found with duplicated number %d", duplicatedNumber)
}

func (c *FilesDataStructures) GetFileByPath(path string) (*FileStruct, error) {
	for index, file := range c.FileStructList {
		if file.Path == path {
			return &c.FileStructList[index], nil
		}
	}
	return &FileStruct{}, fmt.Errorf("no file found for path: %s", path)
}

func Delete(pathsToDelete []string, c *FilesDataStructures) int64 {
	var bytesDeleted int64
	for _, path := range pathsToDelete {
		file, err := c.GetFileByPath(path)
		if err != nil {
			log.Fatal(err)
		}
		bytesDeleted += file.Size
		err = os.Remove(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	return bytesDeleted
}
