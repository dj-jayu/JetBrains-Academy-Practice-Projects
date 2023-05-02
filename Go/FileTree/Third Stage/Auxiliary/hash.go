package Auxiliary

import (
	"crypto/md5"
	"fmt"
	"log"
)

// CreateSizeHashesFilesMapMap create a map[int64]map[[]byte][]FileStruct:
//
//	map{
//	    size1{
//		      hash1:{file1, file2...}
//		      hash2:{file3, file4...}
//	    }
//	    size2{
//		      hash1:{file1, file2...}
//		      hash2:{file3, file4...}
//	    }
//	}
func (c *FilesDataStructures) CreateSizeHashesFilesMapMap() {
	for _, files := range c.SizeFilesMap {
		for _, file := range files {
			if len(c.SizeHashesFilesMapMap[file.Size]) == 0 {
				c.SizeHashesFilesMapMap[file.Size] = make(map[[16]byte][]FileStruct)
			}
			c.SizeHashesFilesMapMap[file.Size][file.Hash] = append(c.SizeHashesFilesMapMap[file.Size][file.Hash], file)
		}
	}
	c.clearNonDuplicates()
}

// Helper method to remove map entries (size) that corresponds to non-duplicated files in SizeHashesFilesMapMap
func (c *FilesDataStructures) clearNonDuplicates() {
	for _, size := range c.OrderedSizeSlice {
		var hashesToDelete [][md5.Size]byte
		// delete unique files hash entry
		for hash, files := range c.SizeHashesFilesMapMap[size] {
			if len(files) == 1 {
				hashesToDelete = append(hashesToDelete, hash)
			}
		}
		for _, hash := range hashesToDelete {
			delete(c.SizeHashesFilesMapMap[size], hash)
		}
		// delete unique files size entry
		if len(c.SizeHashesFilesMapMap[size]) == 0 {
			delete(c.SizeHashesFilesMapMap, size)
		}
	}
}

// PrintDuplicatesInfo FilesDataStructures iterate through the SizeHashesFilesMapMap based on the slice of ordered file sizes
// printing the duplicated files
func (c *FilesDataStructures) PrintDuplicatesInfo() {
	var duplicatedNumber = 1
	for _, size := range c.OrderedSizeSlice {
		if _, ok := c.SizeHashesFilesMapMap[size]; !ok {
			continue
		}
		fmt.Printf("\n%d bytes\n", size)
		hashFilesMap := c.SizeHashesFilesMapMap[size]
		for hash, files := range hashFilesMap {
			fmt.Printf("Hash: %x\n", hash)
			for _, file := range files {
				c.UpdateFileDuplicatedNumberByPath(file.Path, duplicatedNumber)
				fmt.Printf("%d. %s\n", duplicatedNumber, file.Path)
				duplicatedNumber++
			}
		}
	}
}

func (c *FilesDataStructures) UpdateFileDuplicatedNumberByPath(path string, duplicatedNumber int) {
	file, err := c.GetFileByPath(path)
	if err != nil {
		log.Fatal(err)
	}
	file.DuplicatedNumber = duplicatedNumber
}
